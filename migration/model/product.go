package model

import "productAPI/entitie"

type Product struct {
	ID             int64   `gorm:"PrimaryKey;autoIncrement"`
	UUid           string  `gorm:"column:uuid"`
	SKU            string  `gorm:"unique"`
	Name           string  `gorm:"column:name"`
	StockTotal     int64   `gorm:"column:stock_total"`
	StockCutting   int64   `gorm:"column:stock_cutting"`
	StockAvailable int64   `gorm:"column:stock_available"`
	PriceFrom      float64 `gorm:"column:price_from"`
	PricePer       float64 `gorm:"column:price_per"`
}

func (ref *Product) FromDomain(data entitie.DtoProduct) {
	ref.UUid = data.UUid
	ref.SKU = data.SKU
	ref.Name = data.Name
	ref.StockTotal = data.StockTotal
	ref.StockCutting = data.StockCutting
	ref.StockAvailable = data.StockAvailable
	ref.PriceFrom = data.PriceFrom
	ref.PricePer = data.PricePer
}

func (ref *Product) ToDomain() entitie.DtoProductList {
	return entitie.DtoProductList{
		ID:             ref.ID,
		UUid:           ref.UUid,
		SKU:            ref.SKU,
		Name:           ref.Name,
		StockTotal:     ref.StockTotal,
		StockCutting:   ref.StockCutting,
		StockAvailable: ref.StockAvailable,
		PriceFrom:      ref.PriceFrom,
		PricePer:       ref.PricePer,
	}
}

func (ref *Product) TableName() string {
	return "products"
}
