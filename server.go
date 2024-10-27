package main

import (
	"encoding/json"
	"fmt"
	pb "github/challenge/protos/distributor/v1"
	"github/challenge/service"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {
	ds := &service.DistributorService{}
	if err := ds.LoadCSVRegions("cities.csv"); err != nil {
		log.Fatalf("Error loading CSV: %v", err)
	}
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(corsOptions.Handler)
	router.Post("/create-distributor", func(w http.ResponseWriter, r *http.Request) {
		var req pb.CreateDistributorRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		resp, err := ds.CreateDistributor(r.Context(), &req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to create distributor: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	})
	router.Post("/check-permissions", func(w http.ResponseWriter, r *http.Request) {
		var req pb.CheckPermissionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		resp, err := ds.CheckPermission(r.Context(), &req)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to check distributor: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	})
	port := ":8000"
	log.Printf("HTTP server listening on %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	// grpcServer := grpc.NewServer()
	// lis, err := net.Listen("tcp", ":5293")
	// if err != nil {
	// 	log.Fatalf("Failed to listen on port %s: %v", "5293", err)
	// }
	// pb.RegisterCreateDistributorServiceServer(grpcServer, &service.DistributorService{})
	// ds := &service.DistributorService{}
	// loadCSVErr := ds.LoadCSVRegions("cities.csv")
	// if loadCSVErr != nil {
	// 	fmt.Println("error in loading csv", loadCSVErr)
	// }
	// log.Printf("gRPC server listening on %s", "5293")
	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("Failed to serve: %v", err)
	// }
}
