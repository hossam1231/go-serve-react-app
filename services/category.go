package services

//go:generate go run github.com/99designs/gqlgen generate
import (
	"database/sql"
	"errors"
	"log"

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

func GetCategory(id *string) (*model.Category, error) {

	// Construct the SELECT query
	query := `SELECT id, title, userId, userId, establishmentId FROM Category WHERE id = ?`

	database, err := db.Connect() // Call the Connect function
	if err != nil {
		log.Fatal(err)
	}

	// Execute the query and get the result
	row := database.QueryRow(query, id)

	// Check for errors in executing the query
	if err := row.Err(); err != nil {
		// Handle the error appropriately
		// For example, you can log the error, return an error response, or take other actions
		logger.Error("Error executing query:", err)
		// return an error or handle it in some way
	}

	// Create an empty model.Category to populate with the result
	category := &model.Category{}

	// Scan the result into the category model
	err = row.Scan(&category.ID, &category.Title, &category.UserID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Handle the case when the record doesn't exist
			logger.Info("Category with ID %s not found", id)
			// category = nil // or set category to some default values if needed
			// return category, nil // or return a custom error if you prefer
		}
		// Handle other errors
		return nil, err
	}

	logger.Success("Success retrieving category with ID %s", id)

	// Return the retrieved category
	return category, nil

}
