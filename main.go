package main

import (
	"fmt"
	"github.com/sugan2111/couponService/coupon/repository"
	"log"
	"net/http"

	"github.com/sugan2111/couponService/coupon/handler"

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
	r.HandleFunc("/coupon/{id}", handler.DeleteProcess).Methods("DELETE")
	r.HandleFunc("/coupon/{id}", handler.RetrieveProcess).Methods("GET")
	r.HandleFunc("/coupon", handler.CreateProcess).Methods("POST")
	r.HandleFunc("/coupons", handler.ListProcess).Methods("GET")
	r.HandleFunc("/coupon/{id}", handler.UpdateProcess).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8001", r))
}
