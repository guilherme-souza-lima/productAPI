package entitie

type DtoProduct struct {
	UUid           string
	SKU            string
	Name           string
	StockTotal     int64
	StockCutting   int64
	StockAvailable int64
	PriceFrom      float64
	PricePer       float64
}

type DtoProductList struct {
	ID             int64
	UUid           string
	SKU            string
	Name           string
	StockTotal     int64
	StockCutting   int64
	StockAvailable int64
	PriceFrom      float64
	PricePer       float64
}

type DtoProductForDelete struct {
	UUid string
	SKU  string
}
