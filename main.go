package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

	helper "github.com/sugan2111/couponService/helpers"
	"github.com/sugan2111/couponService/models"

	"github.com/gorilla/mux"
)

func createCoupon(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var coupon models.Coupon
	// we decode our body request params
	_ = json.NewDecoder(request.Body).Decode(&coupon)

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

func retrieveCoupon(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var coupon models.Coupon

	collection := helper.ConnectDB()

	err := collection.FindOne(context.TODO(), models.Coupon{ID: id}).Decode(&coupon)

	if err != nil {
		helper.GetError(err, response)
		return
	}

	json.NewEncoder(response).Encode(coupon)
}

func updateCoupon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var coupon models.Coupon

	collection := helper.ConnectDB()

	// Create filter
	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&coupon)

	update := bson.D{
		{"$set", bson.D{
			{"name", coupon.Name},
			{"brand", coupon.Brand},
			{"value", coupon.Value},
			{"createdAt", coupon.CreatedAt},
			{"expiry", coupon.Expiry},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&coupon)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	coupon.ID = id

	json.NewEncoder(w).Encode(coupon)
}

func listCoupons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var coupons []models.Coupon

	collection := helper.ConnectDB()

	// bson.M{} to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var coupon models.Coupon
		err := cur.Decode(&coupon)
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		coupons = append(coupons, coupon)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(coupons)
}

func deleteCoupon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	collection := helper.ConnectDB()

	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

func main() {
	fmt.Println("Building the coupon service...")

	router := mux.NewRouter()
	router.HandleFunc("/coupon", createCoupon).Methods("POST")
	router.HandleFunc("/coupon/{id}", updateCoupon).Methods("PUT")
	router.HandleFunc("/coupon/{id}", retrieveCoupon).Methods("GET")
	router.HandleFunc("/coupon/{id}", deleteCoupon).Methods("DELETE")
	router.HandleFunc("/coupons", listCoupons).Methods("GET")
	log.Fatal(http.ListenAndServe(":8001", router))
}
