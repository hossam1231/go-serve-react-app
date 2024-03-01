package main

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"time" // Added for time.Minute in httprate.LimitByIP

	"github.com/99designs/gqlgen/graphql/handler"    // Added for handler.NewDefaultServer
	"github.com/99designs/gqlgen/graphql/playground" // Added for playground.Handler
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"        // Added for middleware like Logger, Recoverer, RequestID, RealIP, Timeout
	"github.com/go-chi/httprate"                 // Added for httprate.LimitByIP
	httpSwagger "github.com/swaggo/http-swagger" // Added for httpSwagger.WrapHandler
	"mosque.icu/go_server/graph"                 // Placeholder for actual import path of your graph package
)

//go:embed web/dist
var assetsFS embed.FS

func main() {

	subFS, err := fs.Sub(assetsFS, "web/dist")
	if err != nil {
		log.Fatalf("failed to create sub file system: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Setup GraphQL server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// Setup routes
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	r.Mount("/swagger", httpSwagger.WrapHandler)

	ingress := chi.NewRouter()
	ingress.Use(httprate.LimitByIP(100, 1*time.Minute))

	r.Mount("/ingress", ingress)

	web := chi.NewRouter()
	web.Use(httprate.LimitByIP(100, 1*time.Minute))

	web.Handle("/", http.FileServerFS(FallbackFS{subFS}))

	ingress.Mount("/web", web)

	esgress := chi.NewRouter()
	esgress.Use(middleware.RequestID)
	esgress.Use(middleware.RealIP)
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	esgress.Use(middleware.Timeout(60 * time.Second))
	esgress.Use(httprate.LimitByIP(100, 1*time.Minute))

	r.Mount("/esgress", esgress)

	api := chi.NewRouter()
	api.Use(httprate.LimitByIP(100, 1*time.Minute))
	api.Get("/increment/{by}", incrementHandler())

	esgress.Mount("/api", server)

	server := chi.NewRouter()
	server.Use(httprate.LimitByIP(100, 1*time.Minute))

	r.Mount("/server", server)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// FallbackFS is a file system that falls back to index.html if the requested file does not exist.
type FallbackFS struct {
	fs fs.FS
}

func (f FallbackFS) Open(name string) (fs.File, error) {
	file, err := f.fs.Open(name)
	if errors.Is(err, fs.ErrNotExist) {
		return f.fs.Open("index.html")
	}
	return file, err
}

func incrementHandler() http.HandlerFunc {
	var count int
	return func(w http.ResponseWriter, r *http.Request) {
		by, err := strconv.Atoi(chi.URLParam(r, "by"))
		if err != nil {
			http.Error(w, "invalid number", http.StatusBadRequest)
			return
		}

		count += by

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"count": %d}`, count)
	}
}
