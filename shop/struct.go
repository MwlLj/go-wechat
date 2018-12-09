package shop

type CAddStockRequest struct {
	ProductId string `json:"product_id"`
	SkuInfo   string `json:"sku_info"`
	Quantity  int    `json:"quantity"`
}
