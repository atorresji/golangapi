package models

type Product struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Stock        uint   `json:"stock"`
	Discontinued bool   `json:"discontinued"`
	CategoryId   uint   `json:"category_id"`
	SupplierId   uint   `json:"supplier_id"`
}