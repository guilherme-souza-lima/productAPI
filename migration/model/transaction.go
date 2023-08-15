package model

import (
	"productAPI/entitie"
	"time"
)

type Transaction struct {
	ID          int64     `gorm:"PrimaryKey;autoIncrement"`
	UUid        string    `gorm:"column:uuid"`
	UUidProduct string    `gorm:"column:uuid_product"`
	SKUProduct  string    `gorm:"column:sku_product"`
	Amount      int64     `gorm:"column:amount"`
	Price       float64   `gorm:"column:price"`
	Date        time.Time `gorm:"column:date"`
}

func (ref *Transaction) FromDomain(data entitie.DtoTransaction) {
	ref.UUid = data.UUid
	ref.UUidProduct = data.UUidProduct
	ref.SKUProduct = data.SKUProduct
	ref.Amount = data.Amount
	ref.Price = data.Price
	ref.Date = data.Date
}

func (ref *Transaction) ToDomain() entitie.DtoListTransaction {
	return entitie.DtoListTransaction{
		ID:          ref.ID,
		UUid:        ref.UUid,
		UUidProduct: ref.UUidProduct,
		SKUProduct:  ref.SKUProduct,
		Amount:      ref.Amount,
		Price:       ref.Price,
		Date:        ref.Date,
	}
}

func (ref *Transaction) TableName() string {
	return "transactions"
}
