package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Server will start at http://localhost:2000/")

	connectDatabse()

	route := mux.NewRouter()

	AddApproutes(route)

	log.Fatal(http.ListenAndServe(":2000", route))
}
