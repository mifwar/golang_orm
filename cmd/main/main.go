package main

import (
	"gamestore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterGameStore(r)
	http.Handle("/", r)

	port := "localhost:8080"
	log.Fatal(http.ListenAndServe(port, r))
}
