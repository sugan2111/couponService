package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sugan2111/couponService/repository"

	"github.com/gorilla/mux"
)

func RetrieveProcess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		//log something
	}

	coupon, err := repository.RetrieveCoupon(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(coupon)
	w.WriteHeader(http.StatusOK)

}
