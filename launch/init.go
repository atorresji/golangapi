package launch

import (
	"fmt"
	"log"
	"net/http"
	"packages/api/endpoints"
)

func LaunchApi() {
	endpoints.CreateProductEndpoint()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HomePage endpoint hit!")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
