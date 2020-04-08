package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/sugan2111/couponService/coupon/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	bson2 "gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	*mongo.Collection
}

// ConnectDB is a function to connect mongoDB
func ConnectDB() *mongo.Collection {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("myrestapi").Collection("coupons")

	return collection
}

func NewClient(uri string) MongoStore {
	return MongoStore{ConnectDB()}

}

func (d MongoStore) RetrieveCoupon(id string) (model.Coupon, error) {
	var coupon model.Coupon
	idVal := bson2.ObjectIdHex(id)

	err := d.FindOne(context.TODO(), model.Coupon{ID: idVal}).Decode(&coupon)
	if err != nil {
		return coupon, fmt.Errorf("unable to retrieve item:%v", err)
	}
	return coupon, nil
}

func (d MongoStore) ListCoupons() ([]model.Coupon, error) {
	var coupons []model.Coupon

	// bson.M{} to get all data.
	cur, err := d.Find(context.TODO(), bson.M{})
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

func (d MongoStore) CreateCoupon(coupon model.Coupon) (string, error) {

	result, err := d.InsertOne(context.TODO(), coupon)
	if err != nil {
		return coupon, fmt.Errorf("unable to insert item:%v", err)
	}

	return result.InsertedID, nil
}

func (d MongoStore) UpdateCoupon(coupon model.Coupon, id string) (model.Coupon, error) {

	idVal, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return coupon, fmt.Errorf("unable to convert id into hex:%v", err)
	}
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

	err = d.FindOneAndUpdate(context.TODO(), filter, update).Decode(&coupon)
	if err != nil {
		return coupon, fmt.Errorf("unable to update item:%v", err)
	}

	coupon.ID = idVal

	return coupon, nil
}

func (d MongoStore) DeleteCoupon(id string) (int64, error) {

	idVal, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	filter := bson.M{"_id": idVal}

	result, err := d.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0, errors.New("500. Internal Server Error")
	}

	return result.DeletedCount, nil
}
