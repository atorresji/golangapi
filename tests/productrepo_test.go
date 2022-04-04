package tests

import (
	"packages/database/dto"
	"packages/database/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {

	repoProductMock := mocks.ProductRepository{}

	products := []dto.ProductDto{
		{
			Id:           1,
			Name:         "Matebook 13",
			Stock:        4,
			Discontinued: false,
			Category:     "LapTop",
			Supplier:     "Huawei",
		},
		{
			Id:           1,
			Name:         "Matebook 14",
			Stock:        6,
			Discontinued: false,
			Category:     "LapTop",
			Supplier:     "Huawei",
		},
	}

	repoProductMock.On("GetProducts").Return(products).Once()

	data := repoProductMock.GetProducts()

	assert.NotNil(t, data)
	assert.Equal(t, data[0].Name, "Matebook 13")
	assert.Equal(t, data[1].Name, "Matebook 14")

}

func TestGetProductById(t *testing.T) {

	repoProductMock := mocks.ProductRepository{}

	product := dto.ProductDto{
		Id:           10,
		Name:         "Matebook 15",
		Stock:        4,
		Discontinued: false,
		Category:     "LapTop",
		Supplier:     "Huawei",
	}

	productId := int(product.Id)

	repoProductMock.On("GetProductById", productId).Return(product).Once()

	data := repoProductMock.GetProductById(productId)

	assert.NotNil(t, data)
	assert.Equal(t, product.Name, "Matebook 15")
}

// func TestCreateProduct(t *testing.T) {

// 	repoProductMock := mocks.ProductRepository{}

// 	product := dto.ProductDto{
// 		Id:           2,
// 		Name:         "Matebook 15",
// 		Stock:        4,
// 		Discontinued: false,
// 		Category:     "LapTop",
// 		Supplier:     "Huawei",
// 	}

// 	productCreate := dto.ProductCreateDto{
// 		Name:         product.Name,
// 		Stock:        product.Stock,
// 		Discontinued: product.Discontinued,
// 		CategoryId:   1,
// 		SupplierId:   2,
// 	}

// 	repoProductMock.On("AddProduct", productCreate).Return(product).Once()

// 	data := repoProductMock.AddProduct(productCreate)

// 	assert.NotNil(t, data)
// 	assert.Equal(t, product.Id, 2)
// 	assert.Equal(t, product.Name, "Matebook 15")
// }
