package request

type Checkout struct {
	ProductId string `json:"product_id"`
	Qty uint16 `json:"qty"`
}