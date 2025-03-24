package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

const validToken = "d4f7e8b9c6a3e2d1f0b5c4d3a2e1f8d7" //Token given by the electricity provider to the user

type Bill struct {
	Amount       float64 `json:"amount"`
	DueDate      string  `json:"due_date"`
	Status       string  `json:"status"`
	UserID       string  `json:"user_id"`
	ProviderID   string  `json:"provider_id"`
	ProviderName string  `json:"provider_name"`
	PaidDate     string  `json:"paid_date,omitempty"`
}

var electricityBills = []Bill{
	{Amount: 147.67, DueDate: "2023-05-22T00:00:00Z", Status: "paid", UserID: "user-3", ProviderID: "provider-2", ProviderName: "Electricity Provider", PaidDate: "2024-03-04T00:00:00Z"},
	{Amount: 146.43, DueDate: "2023-05-06T00:00:00Z", Status: "paid", UserID: "user-9", ProviderID: "provider-1", ProviderName: "Internet Provider", PaidDate: "2024-04-21T00:00:00Z"},
	{Amount: 148.69, DueDate: "2023-06-17T00:00:00Z", Status: "paid", UserID: "user-10", ProviderID: "provider-3", ProviderName: "Electricity Provider", PaidDate: "2025-02-16T00:00:00Z"},
	{Amount: 229.91, DueDate: "2024-06-17T00:00:00Z", Status: "pending", UserID: "user-9", ProviderID: "provider-1", ProviderName: "Electricity Provider"},
	{Amount: 125.22, DueDate: "2023-07-24T00:00:00Z", Status: "pending", UserID: "user-4", ProviderID: "provider-2", ProviderName: "Water Provider"},
	{Amount: 80.26, DueDate: "2024-11-07T00:00:00Z", Status: "pending", UserID: "user-9", ProviderID: "provider-2", ProviderName: "Electricity Provider"},
	{Amount: 187.32, DueDate: "2023-09-19T00:00:00Z", Status: "paid", UserID: "user-10", ProviderID: "provider-2", ProviderName: "Electricity Provider", PaidDate: "2025-03-30T00:00:00Z"},
	{Amount: 248.10, DueDate: "2023-11-06T00:00:00Z", Status: "paid", UserID: "user-9", ProviderID: "provider-2", ProviderName: "Electricity Provider", PaidDate: "2024-10-05T00:00:00Z"},
	{Amount: 152.86, DueDate: "2023-07-14T00:00:00Z", Status: "paid", UserID: "user-3", ProviderID: "provider-2", ProviderName: "Water Provider", PaidDate: "2023-12-27T00:00:00Z"},
	{Amount: 99.40, DueDate: "2024-02-05T00:00:00Z", Status: "pending", UserID: "user-4", ProviderID: "provider-2", ProviderName: "Internet Provider"},
	{Amount: 216.40, DueDate: "2024-04-03T00:00:00Z", Status: "paid", UserID: "user-2", ProviderID: "provider-2", ProviderName: "Internet Provider", PaidDate: "2024-09-20T00:00:00Z"},
	{Amount: 144.84, DueDate: "2023-10-29T00:00:00Z", Status: "pending", UserID: "user-4", ProviderID: "provider-1", ProviderName: "Water Provider"},
	{Amount: 87.01, DueDate: "2024-01-20T00:00:00Z", Status: "pending", UserID: "user-9", ProviderID: "provider-3", ProviderName: "Water Provider"},
	{Amount: 118.47, DueDate: "2024-04-22T00:00:00Z", Status: "pending", UserID: "user-4", ProviderID: "provider-3", ProviderName: "Internet Provider"},
	{Amount: 277.20, DueDate: "2023-10-25T00:00:00Z", Status: "paid", UserID: "user-2", ProviderID: "provider-3", ProviderName: "Internet Provider", PaidDate: "2025-03-12T00:00:00Z"},
	{Amount: 219.21, DueDate: "2023-10-22T00:00:00Z", Status: "paid", UserID: "user-6", ProviderID: "provider-3", ProviderName: "Internet Provider", PaidDate: "2024-11-03T00:00:00Z"},
	{Amount: 138.40, DueDate: "2023-01-02T00:00:00Z", Status: "paid", UserID: "user-5", ProviderID: "provider-1", ProviderName: "Electricity Provider", PaidDate: "2024-07-02T00:00:00Z"},
	{Amount: 147.22, DueDate: "2023-02-07T00:00:00Z", Status: "pending", UserID: "user-6", ProviderID: "provider-3", ProviderName: "Internet Provider"},
	{Amount: 196.50, DueDate: "2023-04-06T00:00:00Z", Status: "pending", UserID: "user-10", ProviderID: "provider-2", ProviderName: "Electricity Provider"},
	{Amount: 87.68, DueDate: "2023-07-15T00:00:00Z", Status: "paid", UserID: "user-5", ProviderID: "provider-2", ProviderName: "Internet Provider", PaidDate: "2024-04-08T00:00:00Z"},
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" || tokenParts[1] != validToken {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func getBills(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(electricityBills)
}

func main() {
	http.HandleFunc("/api/v1/bills", authMiddleware(getBills))

	log.Println("Electtic mock server running on http://localhost:5001")
	log.Fatal(http.ListenAndServe(":5001", nil))
}
