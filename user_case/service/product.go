package service

import (
	"productAPI/entitie"
	"productAPI/user_case/repository"
	"productAPI/user_case/request"
	"productAPI/utils"
)

type InterfaceProductService interface {
	Create(data request.CreateProduct) error
	ConsultAll() ([]entitie.DtoProductList, error)
	Update(data request.UpdateProduct) error
	Delete(data request.DeleteProduct) error
}

type ProductService struct {
	InterfaceProductRepository repository.InterfaceProductRepository
}

func NewProductService(InterfaceProductRepository repository.InterfaceProductRepository) ProductService {
	return ProductService{InterfaceProductRepository}
}

func (p ProductService) Create(data request.CreateProduct) error {
	err := p.checkPrice(data.PriceFrom, data.PricePer)
	if err != nil {
		return err
	}

	uuid, err := utils.GeneratorUUid()
	if err != nil {
		return err
	}

	available := p.checkAvailableStock(data.StockTotal, data.StockCutting)
	product := entitie.DtoProduct{
		UUid:           uuid,
		SKU:            data.SKU,
		Name:           data.Name,
		StockTotal:     data.StockTotal,
		StockCutting:   data.StockCutting,
		StockAvailable: available,
		PriceFrom:      data.PriceFrom,
		PricePer:       data.PricePer,
	}
	return p.InterfaceProductRepository.Create(product)
}

func (p ProductService) ConsultAll() ([]entitie.DtoProductList, error) {
	return p.InterfaceProductRepository.ConsultAll()
}

func (p ProductService) Update(data request.UpdateProduct) error {
	err := p.checkPrice(data.PriceFrom, data.PricePer)
	if err != nil {
		return err
	}
	available := p.checkAvailableStock(data.StockTotal, data.StockCutting)
	product := entitie.DtoProduct{
		UUid:           data.UUid,
		SKU:            data.SKU,
		Name:           data.Name,
		StockTotal:     data.StockTotal,
		StockCutting:   data.StockCutting,
		StockAvailable: available,
		PriceFrom:      data.PriceFrom,
		PricePer:       data.PricePer,
	}
	return p.InterfaceProductRepository.Update(product)
}

func (p ProductService) Delete(data request.DeleteProduct) error {
	product := entitie.DtoProductForDelete{
		UUid: data.UUid,
		SKU:  data.SKU,
	}
	return p.InterfaceProductRepository.Delete(product)
}
