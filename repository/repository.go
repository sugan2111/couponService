package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/sugan2111/couponService/repository/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	DeleteCoupon(id string) (int64, error)
	UpdateCoupon(r *http.Request, id string) (model.Coupon, error)
	CreateCoupon(coupon model.Coupon) (interface{}, error)
	ListCoupons() ([]model.Coupon, error)
	RetrieveCoupon(id string) (model.Coupon, error)
}

func RetrieveCoupon(id string) (model.Coupon, error) {

	coupon := model.Coupon{}
	idVal, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return coupon, fmt.Errorf("unable to convert id into hex:%v", err)
	}

	collection := ConnectDB()

	err = collection.FindOne(context.TODO(), model.Coupon{ID: idVal}).Decode(&coupon)
	if err != nil {
		return coupon, fmt.Errorf("unable to retrieve item:%v", err)
	}
	return coupon, nil
}

func ListCoupons() ([]model.Coupon, error) {
	collection := ConnectDB()
	var coupons []model.Coupon

	// bson.M{} to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("unable to list coupons:%v", err)
	}

	// Close the cursor once finished
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var coupon model.Coupon
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

	return coupons, nil
}

func CreateCoupon(coupon model.Coupon) (interface{}, error) {

	collection := ConnectDB()

	result, err := collection.InsertOne(context.TODO(), coupon)
	if err != nil {
		return coupon, fmt.Errorf("unable to insert item:%v", err)
	}

	return result.InsertedID, nil
}

func UpdateCoupon(coupon model.Coupon, id string) (model.Coupon, error) {

	idVal, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return coupon, fmt.Errorf("unable to convert id into hex:%v", err)
	}

	collection := ConnectDB()

	// Create filter
	filter := bson.M{"_id": idVal}

	update := bson.D{
		{"$set", bson.D{
			{"name", coupon.Name},
			{"brand", coupon.Brand},
			{"value", coupon.Value},
			{"createdAt", coupon.CreatedAt},
			{"expiry", coupon.Expiry},
		}},
	}

	err = collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&coupon)
	if err != nil {
		return coupon, fmt.Errorf("unable to update item:%v", err)
	}

	coupon.ID = idVal

	return coupon, nil
}

func DeleteCoupon(id string) (int64, error) {

	idVal, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	collection := ConnectDB()

	filter := bson.M{"_id": idVal}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, errors.New("500. Internal Server Error")
	}

	return result.DeletedCount, nil
}
