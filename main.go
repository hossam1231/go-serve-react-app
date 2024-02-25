package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	router.Get("/api/increment/{by}", incrementHandler())
	router.Handle("/*", http.FileServerFS(FallbackFS{os.DirFS("web/dist")}))

	log.Fatal(http.ListenAndServe(":3000", router))
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
