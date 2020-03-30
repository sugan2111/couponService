package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sugan2111/couponService/repository"
)

func ListProcess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	coupons, err := repository.ListCoupons()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(coupons)
	w.WriteHeader(http.StatusOK)

}
