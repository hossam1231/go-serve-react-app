package graph

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is a struct that holds the state of the application.
// In this case, it holds a slice of pointers to Todo models.
// This slice is used to store all the Todo items created in the application.
// The Resolver struct is used in the resolvers to access and manipulate the application state.
type Resolver struct{

}
