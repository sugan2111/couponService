package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sugan2111/couponService/coupon/repository"
)

//ListProcess lists all coupons
func ListProcess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	coupons, err := repository.DB.ListCoupons()
	if err != nil {
		http.Error(w, "No coupons exist.", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(coupons)
	w.WriteHeader(http.StatusOK)

}
