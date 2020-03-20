package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	helper "github.com/sugan2111/restapi/helpers"
	"github.com/sugan2111/restapi/models"

	"github.com/gorilla/mux"
)

func createCoupon(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var coupon models.Coupon
	fmt.Println("Inside create coupon...")
	// we decode our body request params
	_ = json.NewDecoder(request.Body).Decode(&coupon)

	fmt.Println("Before connecting DB...")
	// connect db
	collection := helper.ConnectDB()

	// insert our coupon model.
	result, err := collection.InsertOne(context.TODO(), coupon)

	if err != nil {
		helper.GetError(err, response)
		return
	}

	json.NewEncoder(response).Encode(result)
}

func main() {
	fmt.Println("Starting the application...")

	//Init Router
	router := mux.NewRouter()
	// arrange our route
	router.HandleFunc("/coupon", createCoupon).Methods("POST")
	//router.HandleFunc("/api/coupons/id", updateCoupon).Methods("PUT")
	//router.HandleFunc("/api/coupons/{id}", retrieveCoupon).Methods("GET")
	//router.HandleFunc("/api/coupons", listCoupons).Methods("GET")
	// set our port address
	log.Fatal(http.ListenAndServe(":8001", router))
}
