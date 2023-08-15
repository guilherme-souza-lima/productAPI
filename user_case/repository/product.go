package repository

import (
	"gorm.io/gorm"
	"productAPI/entitie"
	"productAPI/migration/model"
)

type InterfaceProductRepository interface {
	Create(data entitie.DtoProduct) error
	ConsultAll() ([]entitie.DtoProductList, error)
	ConsultBySKU(sku string) ([]entitie.DtoProductList, error)
	Update(data entitie.DtoProduct) error
	Delete(data entitie.DtoProductForDelete) error
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (p ProductRepository) Create(data entitie.DtoProduct) error {
	var entity model.Product
	entity.FromDomain(data)
	err := p.db.Create(&entity)
	if err != nil {
		return err.Error
	}
	return nil
}

func (p ProductRepository) ConsultAll() ([]entitie.DtoProductList, error) {
	var entity model.Product

	filter := p.db.Model(&entity)
	result, err := filter.Rows()
	if err != nil {
		return nil, err
	}

	var listProduct []entitie.DtoProductList
	for result.Next() {
		var list model.Product
		result.Scan(
			&list.ID,
			&list.UUid,
			&list.SKU,
			&list.Name,
			&list.StockTotal,
			&list.StockCutting,
			&list.StockAvailable,
			&list.PriceFrom,
			&list.PricePer,
		)
		listProduct = append(listProduct, list.ToDomain())
	}

	return listProduct, nil
}

func (p ProductRepository) ConsultBySKU(sku string) ([]entitie.DtoProductList, error) {
	var entity model.Product

	filter := p.db.Model(&entity).Where("sku = ?", sku)
	result, err := filter.Rows()
	if err != nil {
		return nil, err
	}

	var listProduct []entitie.DtoProductList
	for result.Next() {
		var list model.Product
		result.Scan(
			&list.ID,
			&list.UUid,
			&list.SKU,
			&list.Name,
			&list.StockTotal,
			&list.StockCutting,
			&list.StockAvailable,
			&list.PriceFrom,
			&list.PricePer,
		)
		listProduct = append(listProduct, list.ToDomain())
	}

	return listProduct, nil
}

func (p ProductRepository) Update(data entitie.DtoProduct) error {
	var entity model.Product

	err := p.db.Model(&entity).Where("uuid = ? and sku = ?", data.UUid, data.SKU).Updates(map[string]interface{}{
		"name":            data.Name,
		"stock_total":     data.StockTotal,
		"stock_cutting":   data.StockCutting,
		"stock_available": data.StockAvailable,
		"price_from":      data.PriceFrom,
		"price_per":       data.PricePer,
	})
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (p ProductRepository) Delete(data entitie.DtoProductForDelete) error {
	var entity model.Product

	err := p.db.Where("uuid = ? and sku = ?", data.UUid, data.SKU).Delete(&entity)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
