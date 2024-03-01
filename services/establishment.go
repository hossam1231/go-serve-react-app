package services

//go:generate go run github.com/99designs/gqlgen generate
import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

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
func GetEstablishmentArticles(establishmentID string, dbID string) ([]*model.Article, error) {

	tableName := fmt.Sprint("Article" + dbID)

	// Construct the SELECT query
	query := "SELECT id, title, description, createdAt, establishmentId  FROM " + tableName + " WHERE establishmentId = ? ORDER BY createdAt DESC LIMIT 10"

	database, err := db.Connect() // Call the Connect function
	if err != nil {
		logger.Error("Error conntecting to db:", err)

	}

	// Execute the query and get the result
	rows, err := database.Query(query, establishmentID)

	if err != nil {
		logger.Error("Error executing query:", err)
		return nil, err
	} else {
		logger.Success("Success getting rows")
	}
	defer rows.Close()

	// Create a slice to store the retrieved articles
	articles := []*model.Article{}

	// Iterate through the rows and scan each into an Article struct
	for rows.Next() {
		article := &model.Article{}
		err := rows.Scan(&article.ID, &article.Title, &article.Description, &article.CreatedAt, &article.EstablishmentID)
		if err != nil {
			logger.Error("Error scanning row:", err)
			return nil, err
		}
		articles = append(articles, article)
	}

	// Check for errors in iterating over rows
	if err := rows.Err(); err != nil {
		logger.Error("Error iterating over rows:", err)
		return nil, err
	}

	// Log success message
	if articles != nil && len(articles) > 0 {
		logger.Success("Success retrieving articles")
	} else {
		logger.Warning("No articles satisfying the conditions")
		return nil, errors.New("No articles satisfying the conditions")
	}
	// Return the retrieved articles

	return articles, nil
}

func GetEstablishmentEvents(establishmentID string, dbID string) ([]*model.Event, error) {

	currentDate := time.Now()

	// panic(fmt.Errorf("not implemented: Events - events"))
	// return &model.Establishment{DbID: obj.DbID, Name: "DbID " + obj.DbID}, nil

	tableName := fmt.Sprint("Event" + dbID)

	// Construct the SELECT query
	query := `SELECT id, title FROM` + tableName + `WHERE establishmentId = ? ORDER BY ABS(DATEDIFF(start, ? )) ASC LIMIT 10`

	database, err := db.Connect() // Call the Connect function
	if err != nil {
		logger.Error("Error connecting to db query:", err)
	}

	// Execute the query and get the result
	rows, err := database.Query(query, establishmentID, currentDate)

	if err != nil {
		logger.Error("Error executing query:", err)
		return nil, err
	} else {
		logger.Success("Success getting rows")
	}
	defer rows.Close()

	// Create a slice to store the retrieved events
	events := []*model.Event{}

	// Iterate through the rows and scan each into an Event struct
	for rows.Next() {
		event := &model.Event{}
		err := rows.Scan(&event.ID, &event.Title, &event.Date)
		if err != nil {
			logger.Error("Error scanning row:", err)
			return nil, err
		}
		events = append(events, event)
	}

	// Check for errors in iterating over rows
	if err := rows.Err(); err != nil {
		logger.Error("Error iterating over rows:", err)
		return nil, err
	}

	// Log success message
	if events != nil && len(events) > 0 {
		logger.Success("Success retrieving events")
	} else {
		logger.Warning("No events satisfying the conditions")
		return nil, errors.New("No events satisfying the conditions")
	}
	// Return the retrieved events

	return events, nil
}

func GetEstablishment(id *string) (*model.Establishment, error) {

	// Construct the SELECT query
	query := `SELECT id, publicId, name, userId, type, dbId FROM Establishment WHERE id = ?`

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

	// Create an empty model.Establishment to populate with the result
	establishment := &model.Establishment{}

	// Scan the result into the establishment model
	err = row.Scan(&establishment.ID, &establishment.PublicID, &establishment.Name, &establishment.UserID, &establishment.Type, &establishment.DbID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Handle the case when the record doesn't exist
			logger.Info("Establishment with ID %s not found", id)
			// establishment = nil // or set establishment to some default values if needed
			// return establishment, nil // or return a custom error if you prefer
		}
		// Handle other errors
		return nil, err
	}

	logger.Success("Success retrieving establishment with ID %s", id)

	// Return the retrieved establishment
	return establishment, nil

}
