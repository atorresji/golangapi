package endpoints

import (
	"log"
	"net/http"
	"packages/api/handlers"

	"github.com/gorilla/mux"
)

func CreateProductEndpoint() {

	mux := mux.NewRouter()

	mux.HandleFunc("/api/products", handlers.GetProducts).Methods("GET")
	mux.HandleFunc("/api/products/{id:[0-9]+}", handlers.GetProductById).Methods("GET")
	mux.HandleFunc("/api/products", handlers.CreateProduct).Methods("POST")
	mux.HandleFunc("/api/products/{id:[0-9]+}", handlers.UpdateProduct).Methods("PUT")
	mux.HandleFunc("/api/products/{id:[0-9]+}", handlers.DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":1234", mux))

}
