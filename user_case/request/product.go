package request

type CreateProduct struct {
	SKU          string  `json:"sku"`
	Name         string  `json:"name"`
	StockTotal   int64   `json:"stock_total"`
	StockCutting int64   `json:"stock_cutting"`
	PriceFrom    float64 `json:"price_from"`
	PricePer     float64 `json:"price_per"`
}

type UpdateProduct struct {
	UUid         string  `json:"uuid"`
	SKU          string  `json:"sku"`
	Name         string  `json:"name"`
	StockTotal   int64   `json:"stock_total"`
	StockCutting int64   `json:"stock_cutting"`
	PriceFrom    float64 `json:"price_from"`
	PricePer     float64 `json:"price_per"`
}

type DeleteProduct struct {
	UUid string `json:"uuid"`
	SKU  string `json:"sku"`
}
