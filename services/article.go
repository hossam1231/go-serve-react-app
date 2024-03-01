package services

//go:generate go run github.com/99designs/gqlgen generate
import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hossam1231/logger-go-pkg"
	"mosque.icu/go_server/db"
	"mosque.icu/go_server/graph/model"
)

// GetArticle retrieves an article from the database by its ID.
func GetArticle(id string, dbID *string) (*model.Article, error) {

	tableName := fmt.Sprintf("Article%s", *dbID)

	// Construct the SELECT query
	query := `SELECT id, description, title, userId, establishmentId, createdAt FROM ` + tableName + ` WHERE id = ?`

	// Connect to the database
	database, err := db.Connect()
	if err != nil {
		logger.Error("Error connecting to db query:", err)
		return nil, err
	}
	defer database.Close()

	// Execute the query and get the result
	row := database.QueryRow(query, id)

	// Create an empty model.Article to populate with the result
	article := &model.Article{}

	// Scan the result into the article model
	err = row.Scan(&article.ID, &article.Description, &article.Title, &article.UserID, &article.EstablishmentID, &article.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Handle the case when the record doesn't exist
			logger.Info("Article with ID %s not found", id)
		}
		// Handle other errors
		logger.Error("Error scanning row:", err)
	}

	logger.Success("Success retrieving article with ID %s", id)

	// Return the retrieved article
	return article, nil
}
