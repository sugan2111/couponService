package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sugan2111/couponService/repository"

	"github.com/gorilla/mux"
)

// RetrieveProcess returns a coupon
func RetrieveProcess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	idVal, ok := params["id"]
	if !ok {
		http.Error(w, fmt.Sprintf("%s is not a valid ID", params["id"]), http.StatusBadRequest)
		return
	}

	coupon, err := repository.DB.RetrieveCoupon(idVal)
	if err != nil {
		http.Error(w, "The coupon you requested does not exist.", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(coupon)
	w.WriteHeader(http.StatusOK)
}
