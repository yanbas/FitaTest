package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)


func NewHandler() (*handler.Handler, error) {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: newQuery(),
			Mutation: Mutation,
		},
	)

	if err != nil {
		return nil, err
	}

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}), nil
}