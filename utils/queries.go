package utils

import (
	"packages/database/dto"
	"strconv"
	"strings"
)

const ProductDetailQuery string = "SELECT * FROM [GoLangAPI].[dbo].[v_ProductDetail]"

var ProductDetailQueryById string = "SELECT * FROM [GoLangAPI].[dbo].[v_ProductDetail] WHERE Id = "

var DeleteProductById string = "DELETE FROM [GoLangAPI].[dbo].[Product] WHERE Id = "

var GetLastInsertedId = "SELECT MAX(ID) FROM [GoLangAPI].[dbo].[v_ProductDetail]"

func CreateUpdateStatement(updateDto *dto.ProductUpdateDto, id int) string {

	discontinuedValue := "0"

	if updateDto.Discontinued {
		discontinuedValue = "1"
	}

	var sb strings.Builder

	sb.WriteString("UPDATE [GoLangAPI].[dbo].[Product] ")
	sb.WriteString("SET [StockUnits] = " + strconv.Itoa(int(updateDto.Stock)))
	sb.WriteString(", [Discontinued] = " + discontinuedValue)
	sb.WriteString(", [CategoryId] = " + strconv.Itoa(int(updateDto.CategoryId)))
	sb.WriteString(", [SupplierId] = " + strconv.Itoa(int(updateDto.SupplierId)))
	sb.WriteString(" WHERE [Id] = " + strconv.Itoa(int(id)))

	return sb.String()
}

func CreateInsertProduct(name string, stock uint, discontinued bool, categoryId, supplierId uint) string {

	insertProduct := `INSERT INTO [GoLangAPI].[dbo].[Product] (
		[Name]
		,[StockUnits]
		,[Discontinued]
		,[CategoryId]
		,[SupplierId])
	VALUES (`

	discontinuedValue := "0"

	if discontinued {
		discontinuedValue = "1"
	}

	parameters := []string{
		`'` + name + `'`,
		strconv.FormatUint(uint64(stock), 10),
		discontinuedValue,
		strconv.FormatUint(uint64(categoryId), 10),
		strconv.FormatUint(uint64(supplierId), 10)}

	values := strings.Join(parameters, ", ")

	insertProduct = insertProduct + values + ");"

	return insertProduct
}
