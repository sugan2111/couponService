package main

import (
	"fmt"
	"github.com/sugan2111/couponService/repository"
	"log"
	"net/http"

	"github.com/sugan2111/couponService/handlers"

	"github.com/gorilla/mux"
)

const (
	// better to read this from environment variables
	dbURI = "mongodb://localhost:27017"
)


func main() {

	fmt.Println("Building the coupon service...")
	store := repository.NewClient(dbURI)
	r := NewRouter(mux.NewRouter(), store)
	r.HandleFunc("/coupon/{id}", handlers.DeleteProcess).Methods("DELETE")
	r.HandleFunc("/coupon/{id}", handlers.RetrieveProcess).Methods("GET")
	r.HandleFunc("/coupon", handlers.CreateProcess).Methods("POST")
	r.HandleFunc("/coupons", handlers.ListProcess).Methods("GET")
	r.HandleFunc("/coupon/{id}", handlers.UpdateProcess).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8001", r))
}
