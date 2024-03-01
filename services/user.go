package services

//go:generate go run github.com/99designs/gqlgen generate
import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hossam1231/logger-go-pkg"
	"mosque.icu/go_server/db"
	"mosque.icu/go_server/graph/model"
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

func GetUser(id *string) (*model.User, error) {
	// Construct the SELECT query
	query := "SELECT id, firstName, lastName FROM User WHERE id = ?"

	// Call the Connect function to establish a database connection
	database, err := db.Connect()
	if err != nil {
		logger.Error("Error connecting to db:", err)
		return nil, err
	}

	defer database.Close() // Close the database connection when done

	// Execute the query and get the result
	row := database.QueryRow(query, *id)

	// Create an empty model.User to populate with the result
	user := &model.User{}

	// Scan the result into the user model
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Handle the case when the record doesn't exist
			logger.Info("User with ID %s not found", *id)
		}
		// Handle other errors
		logger.Error("Error scanning rows:", err)
	}

	logger.Info("User:", user)
	logger.Success("Success retrieving user with ID %s", *id)

	// Return the retrieved user
	return user, nil

}
