package repositories

import (
	"database/sql"
	"fmt"
	"packages/database"
	"packages/database/dto"
	"packages/utils"
	"strconv"
)

func AddProduct(pr *dto.ProductCreateDto) bool {

	insert := utils.CreateInsertProduct(pr.Name, pr.Stock, pr.Discontinued, pr.SupplierId, pr.CategoryId)
	fmt.Println(insert)
	database.NonQueryStatement(insert)

	return true
}

func GetProducts() []dto.ProductDto {

	var rows *sql.Rows = database.QueryStatement(utils.ProductDetailQuery)

	var products = []dto.ProductDto{}

	for rows.Next() {
		product := dto.ProductDto{}
		rows.Scan(&product.Id, &product.Name, &product.Stock, &product.Discontinued, &product.Category, &product.Supplier)
		products = append(products, product)
	}

	return products
}

func GetProductById(id int) dto.ProductDto {

	query := utils.ProductDetailQueryById + strconv.Itoa(id)

	var rows *sql.Rows = database.QueryStatement(query)

	var product dto.ProductDto

	if rows.Next() {
		rows.Scan(&product.Id, &product.Name, &product.Stock, &product.Discontinued, &product.Category, &product.Supplier)
	}

	return product
}

func UpdateProduct(pr *dto.ProductUpdateDto, id int) bool {

	update := utils.CreateUpdateStatement(pr, id)
	fmt.Println(update)
	database.NonQueryStatement(update)
	return true
}

func DeleteProduct(id int) bool {
	delete := utils.DeleteProductById + strconv.Itoa(id)
	fmt.Println(delete)
	database.NonQueryStatement(delete)
	return true
}
