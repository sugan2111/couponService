package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Coupon struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Brand     string             `json:"brand,omitempty" bson:"brand,omitempty"`
	Value     int                `json:"value,omitempty" bson:"value,omitempty"`
	CreatedAt string             `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	Expiry    string             `json:"expiry,omitempty" bson:"expiry,omitempty"`
}
