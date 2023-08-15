package request

type CreateTransaction struct {
	UUidProduct string `json:"uuid_product"`
	SKUProduct  string `json:"sku"`
	Amount      int64  `json:"amount"`
}
