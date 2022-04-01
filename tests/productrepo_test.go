package tests

import (
	"packages/database/dto"
	"packages/database/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProductById(t *testing.T) {

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

	repoProductMock.On("GetProducts").Return(products)

	data := repoProductMock.GetProducts()

	assert.NotNil(t, data)
	assert.Equal(t, data[0].Name, "Matebook 13")

}
