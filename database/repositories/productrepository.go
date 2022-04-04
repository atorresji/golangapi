package repositories

import (
	"database/sql"
	"fmt"
	"packages/database"
	"packages/database/dto"
	"packages/database/models"
	"packages/utils"
	"strconv"
)

type ProductRepository interface {
	AddProduct(p models.Product) models.Product

	GetProducts() []dto.ProductDto

	GetProductById(id int) dto.ProductDto

	UpadteProduct(p models.Product) bool

	DeleteProduct(id int) bool
}

func GetProducts() []dto.ProductDto {

	var rows *sql.Rows = database.QueryStatement(utils.ProductDetailQuery)

	products := []dto.ProductDto{}

	for rows.Next() {

		product := dto.ProductDto{}
		rows.Scan(&product.Id, &product.Name, &product.Stock, &product.Discontinued, &product.Category, &product.Supplier)
		products = append(products, product)
	}

	fmt.Println(products)

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

func AddProduct(pr *dto.ProductCreateDto) dto.ProductDto {

	insert := utils.CreateInsertProduct(pr.Name, pr.Stock, pr.Discontinued, pr.SupplierId, pr.CategoryId)
	fmt.Println(insert)
	database.NonQueryStatement(insert)
	productCreated := GetProductById(getLastId())

	return productCreated
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

func getLastId() int {
	query := utils.GetLastInsertedId
	fmt.Println(query)
	var rows *sql.Rows = database.QueryStatement(query)

	id := 0

	if rows.Next() {
		rows.Scan(&id)
	}
	return id

}
