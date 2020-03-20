package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create Struct
type Coupon struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Brand string             `json:"brand,omitempty" bson:"brand,omitempty"`
	Value string             `json:"value,omitempty" bson:"value,omitempty"`
	//CreatedAt string             `json:"createdat,omitempty" bson:"createdat,omitempty"`
	//Expiry    string             `json:"expiry,omitempty" bson:"expiry,omitempty"`
}
