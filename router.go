package main

import (
	"github.com/gorilla/mux"
	"github.com/sugan2111/couponService/coupon/repository"
)

type Router struct {
	*mux.Router
	store repository.MongoStore
}

func NewRouter(r *mux.Router, store repository.MongoStore) Router {
	return Router{Router: r, store: store}
}
