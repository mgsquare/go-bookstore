package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mgsquare/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	// Print message before starting the server
	fmt.Println("Server is running on port 9010...")

	// Start the server and use r as the handler
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
