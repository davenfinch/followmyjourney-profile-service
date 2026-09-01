package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/davenfinch/followmyjourney-profile-service/internal/api"
	"github.com/davenfinch/followmyjourney-profile-service/internal/service"
	"github.com/davenfinch/followmyjourney-profile-service/internal/store"
)

func main() {
	addr := ":8081"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}

	// In-memory store for initial implementation
	st := store.NewInMemoryStore()
	svc := service.NewProfileService(st)
	h := api.NewHandler(svc)

	r := mux.NewRouter()
	h.RegisterRoutes(r)

	log.Printf("starting server on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
