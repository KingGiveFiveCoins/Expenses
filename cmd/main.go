package main

import (
	"log"
	"net/http"

	"github.com/KingGiveFiveCoins/Expenses/internal/config"
	"github.com/KingGiveFiveCoins/Expenses/internal/handlers"
	"github.com/KingGiveFiveCoins/Expenses/internal/middleware"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Expenses API
// @version 1.0
// @description This is a Expense microserivce.
// @host localhost:8080
// @BasePath /Expenses
func main() {
	cfg := config.Load()

	r := mux.NewRouter()

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	r.Handle("/expenses", middleware.JWTAuth(http.HandlerFunc(handlers.Expenses))).Methods("GET")
	r.Handle("/add-expense", middleware.JWTAuth(http.HandlerFunc(handlers.CreateExpenses))).Methods("POST")
	r.Handle("/update-expense", middleware.JWTAuth(http.HandlerFunc(handlers.UpdateExpense))).Methods("PUT")
	r.Handle("/delete-expense", middleware.JWTAuth(http.HandlerFunc(handlers.DeleteExpense))).Methods("DELETE")

	log.Println("Server is running on port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
