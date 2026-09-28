package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"payment-service/controllers"
	"payment-service/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	godotenv.Load()

	r := mux.NewRouter()

	pc := &controllers.PaymentController{
		JwtKey: []byte(os.Getenv("JWT_SECRET")),
	}

	// 👉 THIS WAS MISSING !
	paymentRouter := r.PathPrefix("/payment").Subrouter()

	fmt.Println("📌 Subrouter created on: /payment")

	// MOUNT ROUTES ON SUBROUTER
	routes.PaymentRoutes(paymentRouter, pc)

	// 👀 print all registered routes
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		fmt.Println("📍 Registered Route:", path, "Methods:", methods)
		return nil
	})

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:5174"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	log.Println("🔥 Payment service running on port 8086")
	log.Fatal(http.ListenAndServe(":8086", corsHandler.Handler(r)))
}
