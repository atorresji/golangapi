package entities

type Product struct {
	Name         string `json:"Name"`
	StockUnits   uint   `json:"Stock"`
	Discontinued bool   `json:"Discontinued"`
	CategoryId   uint   `json:"CategoryId"`
	SupplierId   uint   `json:"SupplierId"`
}
