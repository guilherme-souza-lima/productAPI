package service

import "errors"

func (p ProductService) checkPrice(priceFrom, pricePer float64) error {
	if pricePer > priceFrom {
		return errors.New("wrong price")
	}
	return nil
}
func (p ProductService) checkAvailableStock(StockTotal, StockCutting int64) int64 {
	return StockTotal - StockCutting
}
