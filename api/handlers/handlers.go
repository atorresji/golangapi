package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"packages/database/dto"
	"packages/database/repositories"
	"strconv"

	"github.com/gorilla/mux"
)

func GetProducts(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	products := repositories.GetProducts()
	data, _ := json.Marshal(products)

	fmt.Fprintln(rw, string(data))
}

func GetProductById(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var product = repositories.GetProductById(id)
	data, _ := json.Marshal(product)

	fmt.Fprintln(rw, string(data))
}

func CreateProduct(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)

	productCreateDto := dto.ProductCreateDto{}
	productDto := dto.ProductDto{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&productCreateDto); err != nil {
		fmt.Println(rw, http.StatusUnprocessableEntity)
	} else {
		productDto = repositories.AddProduct(&productCreateDto)
	}

	data, _ := json.Marshal(productDto)

	fmt.Fprintln(rw, string(data))

}

func UpdateProduct(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//fmt.Println("Se actualizara el producto " + strconv.Itoa(id))

	product := dto.ProductUpdateDto{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&product); err != nil {
		fmt.Println(rw, http.StatusUnprocessableEntity)
	} else {
		fmt.Println("Vamo a actualizar el registro")
		repositories.UpdateProduct(&product, id)
	}

	data, _ := json.Marshal(product)

	fmt.Fprintln(rw, string(data))
}

func DeleteProduct(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusNoContent)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	repositories.DeleteProduct(id)

	fmt.Fprintln(rw)
}
