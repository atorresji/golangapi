package dto

type ProductDto struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Stock        uint   `json:"stock"`
	Discontinued bool   `json:"discontinued"`
	Category     string `json:"category"`
	Supplier     string `json:"supplier"`
}

type ProductCreateDto struct {
	Name         string `json:"name"`
	Stock        uint   `json:"stock"`
	Discontinued bool   `json:"discontinued"`
	CategoryId   uint   `json:"categoryid"`
	SupplierId   uint   `json:"supplierid"`
}

type ProductUpdateDto struct {
	Stock        uint `json:"stock"`
	Discontinued bool `json:"discontinued"`
	CategoryId   uint `json:"categoryid"`
	SupplierId   uint `json:"supplierid"`
}
