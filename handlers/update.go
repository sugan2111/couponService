package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sugan2111/couponService/repository/model"

	"github.com/sugan2111/couponService/repository"

	"github.com/gorilla/mux"
)

func UpdateProcess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		//log something
	}

	var coupon model.Coupon
	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&coupon)

	coupon, err := repository.DB.UpdateCoupon(coupon, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(coupon)
	w.WriteHeader(http.StatusOK)

}
