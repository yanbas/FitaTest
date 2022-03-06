package graphql

import (
	"fita/DB"
	"fita/repository"
	"fita/service"

	"github.com/graphql-go/graphql"
)

var (
	checkout = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Checkout",
			Fields: graphql.Fields{
				"product_id": &graphql.Field{Type: graphql.ID},
				"qty":     &graphql.Field{Type: graphql.String},
				"price":     &graphql.Field{Type: graphql.String},
				"promo":     &graphql.Field{Type: graphql.String},				
			},
			Description: "Checkout",
		},
	)
	products = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Products",
			Fields: graphql.Fields{
				"name":        &graphql.Field{Type: graphql.ID},
				"sku":         &graphql.Field{Type: graphql.String},
				"price": &graphql.Field{Type: graphql.String},
				"qty":       &graphql.Field{Type: graphql.String},
			},
			Description: "Products",
		},
	)
	promotion = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Promotion",
			Fields: graphql.Fields{
				"product_id":        &graphql.Field{Type: graphql.ID},
				"type_promo":         &graphql.Field{Type: graphql.String},
				"amount": &graphql.Field{Type: graphql.String},
				"item":       &graphql.Field{Type: graphql.String},
				"qty":       &graphql.Field{Type: graphql.String},
				"percentage":       &graphql.Field{Type: graphql.String},				
			},
			Description: "Promotion",
		},
	)
)

// Register Repository
var productRepository = repository.NewProductRepository(DB.DB)
// Register Services
var productSvc = service.NewProductService(productRepository)

func Product() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(products),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			listLevel, err := productSvc.Products()
			return listLevel, err
		},
		Description: "Products",
		Args: graphql.FieldConfigArgument{
			"keywords": &graphql.ArgumentConfig{Type: graphql.String},
		},
	}
}

func Promotion() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(promotion),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			listUser, err := productSvc.Promotion()
			return listUser, err
		},
		Description: "List Promotion",
		Args: graphql.FieldConfigArgument{
			"keywords": &graphql.ArgumentConfig{Type: graphql.String},
		},
	}
}

func newQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"Products": Product(),
			"Promotion": Promotion(),
		},
	})
}