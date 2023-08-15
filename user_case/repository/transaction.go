package repository

import (
	"gorm.io/gorm"
	"productAPI/entitie"
	"productAPI/migration/model"
)

type InterfaceTransactionRepository interface {
	Create(data entitie.DtoTransaction) error
	List() ([]entitie.DtoListTransaction, error)
}

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return TransactionRepository{db}
}

func (t TransactionRepository) Create(data entitie.DtoTransaction) error {
	var entity model.Transaction
	entity.FromDomain(data)
	err := t.db.Create(&entity)
	if err != nil {
		return err.Error
	}
	return nil
}

func (t TransactionRepository) List() ([]entitie.DtoListTransaction, error) {
	var entity model.Transaction

	filter := t.db.Model(&entity)
	result, err := filter.Rows()
	if err != nil {
		return nil, err
	}

	var listTransaction []entitie.DtoListTransaction
	for result.Next() {
		var list model.Transaction
		result.Scan(
			&list.ID,
			&list.UUid,
			&list.UUidProduct,
			&list.SKUProduct,
			&list.Amount,
			&list.Price,
			&list.Date,
		)
		listTransaction = append(listTransaction, list.ToDomain())
	}
	return listTransaction, nil
}
