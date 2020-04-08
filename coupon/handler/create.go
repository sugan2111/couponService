package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sugan2111/couponService/coupon/repository"

	"github.com/sugan2111/couponService/coupon/model"
)

// CreateProcess adds a new coupon
func CreateProcess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var coupon model.Coupon
	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&coupon)

	result, err := repository.DB.CreateCoupon(coupon)
	if err != nil {
		http.Error(w, "Cannot create coupon!", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
	w.WriteHeader(http.StatusOK)
}
