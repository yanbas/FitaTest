package graphql

import (
	"fita/DB"
	"fita/repository"
	"fita/service"

	"github.com/graphql-go/graphql"
)

var checkoutRepository = repository.NewCheckoutRepository(DB.DB)
var checkoutSvc = service.NewCheckoutService(checkoutRepository)


var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name:"Mutation",
	Fields:graphql.Fields{
		"CheckoutProduct":&graphql.Field{
			Type:graphql.NewList(CheckoutTypes),
			//config param argument
			Args:graphql.FieldConfigArgument{
				"ID_PRO": &graphql.ArgumentConfig{
					Type:graphql.NewNonNull(graphql.String),
				},
				"QTY": &graphql.ArgumentConfig{
					Type:graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: checkoutSvc.Checkout,
		},
	},
})

var CheckoutTypes = graphql.NewObject(graphql.ObjectConfig{
	Name:"Checkout",
	Fields:graphql.Fields{
		"ID_PRO":&graphql.Field{
			Type:graphql.String,
		},
		"QTY":&graphql.Field{
			Type:graphql.String,
		},
	},
})
