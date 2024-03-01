package services

//go:generate go run github.com/99designs/gqlgen generate
import (
	"database/sql"
	"errors"
	"fmt"
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
func GetArticleComments(id string) ([]*model.Comment, error) {

	tableName := fmt.Sprint("A_ArticleComment")

	// Construct the SELECT query
	query := "SELECT id, title, description, createdAt, likes  FROM " + tableName + " WHERE articleId = ?"

	database, err := db.Connect() // Call the Connect function
	if err != nil {
		logger.Error("Error conntecting to db:", err)

	}

	// Execute the query and get the result
	rows, err := database.Query(query, id)

	if err != nil {
		logger.Error("Error executing query:", err)
		return nil, err
	} else {
		logger.Success("Success getting rows")
	}
	defer rows.Close()

	// Create a slice to store the retrieved comments
	comments := []*model.Comment{}

	// Iterate through the rows and scan each into an Comment struct
	for rows.Next() {
		comment := &model.Comment{}
		err := rows.Scan(&comment.ID, &comment.Title, &comment.Likes, &comment.CreatedAt, comment.UserID)
		if err != nil {
			logger.Error("Error scanning row:", err)
			return nil, err
		}
		comments = append(comments, comment)
	}

	// Check for errors in iterating over rows
	if err := rows.Err(); err != nil {
		logger.Error("Error iterating over rows:", err)
		return nil, err
	}

	// Log success message
	if comments != nil && len(comments) > 0 {
		logger.Success("Success retrieving comments")
	} else {
		logger.Warning("No comments satisfying the conditions")
		return nil, errors.New("No comments satisfying the conditions")
	}
	// Return the retrieved comments

	return comments, nil
}

func GetArticleComment(id *int) (*model.Comment, error) {

	tableName := fmt.Sprint("A_ArticleComment")

	// Construct the SELECT query
	query := "SELECT id, title, comment, createdAt, likes, userId  FROM " + tableName + " WHERE articleId = ?"

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

	// Create an empty model.Comment to populate with the result
	comment := &model.Comment{}

	// Scan the result into the comment model
	err = row.Scan(&comment.ID, &comment.Title, &comment.UserID, &comment.CreatedAt, &comment.Likes)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Handle the case when the record doesn't exist
			logger.Info("Comment with ID %s not found", id)
			// comment = nil // or set comment to some default values if needed
			// return comment, nil // or return a custom error if you prefer
		}
		// Handle other errors
		return nil, err
	}

	logger.Success("Success retrieving comment with ID %s", id)

	// Return the retrieved comment
	return comment, nil

}
