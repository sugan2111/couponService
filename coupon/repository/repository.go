package repository

import (
	"github.com/sugan2111/couponService/coupon/model"
)

type Repository interface {
	DeleteCoupon(id string) (int64, error)
	UpdateCoupon(coupon model.Coupon, id string) (model.Coupon, error)
	CreateCoupon(coupon model.Coupon) (string, error)
	ListCoupons() ([]model.Coupon, error)
	RetrieveCoupon(id string) (model.Coupon, error)
}

var DB Repository
