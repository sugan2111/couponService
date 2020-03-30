package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sugan2111/couponService/repository"

	"github.com/gorilla/mux"
)

func DeleteProcess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		/*errorResponse(w, http.StatusBadRequest,
			NewCustomError(InvalidRequest, err))
		return*/
	}

	deleteResult, err := repository.DeleteCoupon(id)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}
