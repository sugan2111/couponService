package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sugan2111/couponService/repository"

	"github.com/sugan2111/couponService/repository/model"
)

func CreateProcess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var coupon model.Coupon
	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&coupon)

	result, err := repository.CreateCoupon(coupon)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(result)
	w.WriteHeader(http.StatusOK)
}
