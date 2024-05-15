package graph

import "github.com/DiegoSenaa/go-graphql/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *database.Category
	CoursesDB  *database.Course
}
