package entitie

import "time"

type DtoTransaction struct {
	UUid        string    `json:"uuid"`
	UUidProduct string    `json:"uuid_product"`
	SKUProduct  string    `json:"sku_product"`
	Amount      int64     `json:"amount"`
	Price       float64   `json:"price"`
	Date        time.Time `json:"date"`
}

type DtoListTransaction struct {
	ID          int64     `json:"id"`
	UUid        string    `json:"uuid"`
	UUidProduct string    `json:"uuid_product"`
	SKUProduct  string    `json:"sku_product"`
	Amount      int64     `json:"amount"`
	Price       float64   `json:"price"`
	Date        time.Time `json:"date"`
}
