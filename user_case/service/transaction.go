package service

import (
	"productAPI/entitie"
	"productAPI/user_case/repository"
	"productAPI/user_case/request"
	"productAPI/utils"
	"time"
)

type InterfaceTransactionService interface {
	Create(data request.CreateTransaction) error
	List() ([]entitie.DtoListTransaction, error)
}

type TransactionService struct {
	InterfaceTransactionRepository repository.InterfaceTransactionRepository
	InterfaceProductRepository     repository.InterfaceProductRepository
}

func NewTransactionService(
	InterfaceTransactionRepository repository.InterfaceTransactionRepository,
	InterfaceProductRepository repository.InterfaceProductRepository) TransactionService {
	return TransactionService{InterfaceTransactionRepository, InterfaceProductRepository}
}

func (t TransactionService) Create(data request.CreateTransaction) error {
	uuid, err := utils.GeneratorUUid()
	if err != nil {
		return err
	}

	result, err := t.InterfaceProductRepository.ConsultBySKU(data.SKUProduct)
	if err != nil {
		return err
	}

	priceTotal := result[0].PricePer * float64(data.Amount)
	transaction := entitie.DtoTransaction{
		UUid:        uuid,
		UUidProduct: data.UUidProduct,
		SKUProduct:  data.SKUProduct,
		Amount:      data.Amount,
		Price:       priceTotal,
		Date:        time.Now().UTC(),
	}
	err = t.InterfaceTransactionRepository.Create(transaction)
	if err != nil {
		return err
	}
	return nil
}

func (t TransactionService) List() ([]entitie.DtoListTransaction, error) {
	return t.InterfaceTransactionRepository.List()
}
