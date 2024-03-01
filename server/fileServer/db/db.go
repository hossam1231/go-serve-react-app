package db

//go:generate go run github.com/99designs/gqlgen generate
import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hossam1231/logger-go-pkg"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is a struct that holds the state of the application.
// In this case, it holds a slice of pointers to Todo models.
// This slice is used to store all the Todo items created in the application.
// The Resolver struct is used in the resolvers to access and manipulate the application state.

// db is a function that connects to a MySQL database.
// It returns a sql.DB object that allows the user to execute SQL statements and get the result.
// dbConnect connects to a MySQL database using the provided dataSourceName (DSN).
// It returns a *sql.DB object to interact with the database and any error encountered.
func Connect() (*sql.DB, error) {
    db, err := sql.Open("mysql", "3n56ba5zfck0hxw8d5tf:pscale_pw_jtsOJoy1TTf4kNkLIfenpgmtrQ7UWcTRCxlkxgivT9f@tcp(aws.connect.psdb.cloud)/backend?tls=true&interpolateParams=true" )
    if err != nil {
        logger.Error("failed to connect: %v", err)
        return nil, err
    }

    if err := db.Ping(); err != nil {
        logger.Error("failed to ping: %v", err)
        return nil, err
    }

    logger.Success("Successfully connected to PlanetScale!")
    return db, nil
}


