package repositories

import (
	"packages/database/dto"
	"packages/database/models"
)

type ProductRepository interface {
	AddProduct(p models.Product) bool

	GetProducts() []dto.ProductDto

	GetProductById(id int) dto.ProductDto

	UpadteProduct(p models.Product) bool

	DeleteProduct(id int) bool
}
