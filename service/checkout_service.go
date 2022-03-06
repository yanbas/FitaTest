package service

import (
	"errors"
	"fita/model/entity"
	"fita/repository"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

type CheckoutService interface {
	Checkout(param graphql.ResolveParams) (interface{},error)
}

type checkoutServiceImpl struct {
	Repository repository.CheckoutRepository 
}

func NewCheckoutService(
	pr repository.CheckoutRepository,
) CheckoutService {
	return &checkoutServiceImpl{pr}
}

func(s *checkoutServiceImpl) Checkout(param graphql.ResolveParams) (interface{},error)  {
	idProduct := param.Args["ID_PRO"].(string)
	
	rawQty, _ := strconv.Atoi(fmt.Sprintf("%v",param.Args["QTY"]))
	qty := uint16(rawQty)

	// check stock
	product, err := s.Repository.GetProduct(idProduct)
	if err != nil {
		return nil, err
	}

	if product.Qty < qty {
		return nil, errors.New("Stock not enought")
	}

	// check promo
	promoCheck, err := s.Repository.GetPromo(idProduct)
	if err != nil {
		return nil, err
	}

	var order []entity.Order
	orderId := uuid.NewString()

	if promoCheck != nil && qty >= promoCheck.Qty{
		if promoCheck.TypePromo == "item" {
			// get product promo
			productPromo, _ := s.Repository.GetProduct(promoCheck.Item)
			order = []entity.Order{
				{
					OrderID: orderId,
					ProductID: idProduct,
					PromoID: promoCheck.ID,
					Qty: qty,
					Price: product.Price,
					GrandTotal: product.Price,
				},
				{
					OrderID: orderId,
					ProductID: productPromo.SKU,
					Qty: 1,
					Price: productPromo.Price,
					GrandTotal: productPromo.Price * float64(qty),
					PromoID: promoCheck.ID,
				},
			}
		}else if promoCheck.TypePromo == "amount" {
			order = []entity.Order{
				{
					OrderID: orderId,
					ProductID: idProduct,
					PromoID: promoCheck.ID,
					Qty: qty,
					Price: product.Price,
					GrandTotal: (product.Price * float64(qty)) - promoCheck.Amount,
					PromoAmount: promoCheck.Amount,
				},
			}
		}else {
			promoAmount := ((product.Price * float64(qty)) * 100) / float64(promoCheck.Percentage)
			order = []entity.Order{
				{
					OrderID: orderId,
					ProductID: idProduct,
					PromoID: promoCheck.ID,
					Qty: qty,
					Price: product.Price,
					GrandTotal: (product.Price * float64(qty)) - promoAmount,
					PromoAmount: promoCheck.Amount,
				},
			}
		}

		
	}else {
		order = []entity.Order{
			{
				OrderID: orderId,
				ProductID: idProduct,
				Price: product.Price,
				Qty: qty,
				GrandTotal: product.Price * float64(qty),
			},
		}
	}

	// store order
	err = s.Repository.Proceed(&order)
	if err != nil {
		return nil, err
	}	

	// update stock
	err = s.Repository.UpdateStock(idProduct, qty)
	if err != nil {
		return nil, err
	}	

	return order, nil
}



