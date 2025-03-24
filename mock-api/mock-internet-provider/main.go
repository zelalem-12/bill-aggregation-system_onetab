package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

const validToken = "a1d2f3b4c5e6d7a8b9c0d1f2a3b4c5d6" //Token given by the Internet provider to the user

type Bill struct {
	Amount       float64 `json:"amount"`
	DueDate      string  `json:"due_date"`
	Status       string  `json:"status"`
	UserID       string  `json:"user_id"`
	ProviderID   string  `json:"provider_id"`
	ProviderName string  `json:"provider_name"`
	PaidDate     string  `json:"paid_date,omitempty"`
}

var internetBills = []Bill{
	{Amount: 122.33, DueDate: "2023-03-05T00:00:00Z", Status: "paid", UserID: "user-11", ProviderID: "provider-2", ProviderName: "Water Provider", PaidDate: "2023-06-12T00:00:00Z"},
	{Amount: 135.48, DueDate: "2023-04-14T00:00:00Z", Status: "pending", UserID: "user-2", ProviderID: "provider-2", ProviderName: "Water Provider"},
	{Amount: 98.75, DueDate: "2023-06-25T00:00:00Z", Status: "paid", UserID: "user-12", ProviderID: "provider-2", ProviderName: "Water Provider", PaidDate: "2023-08-05T00:00:00Z"},
	{Amount: 105.50, DueDate: "2023-09-08T00:00:00Z", Status: "pending", UserID: "user-13", ProviderID: "provider-2", ProviderName: "Water Provider"},
	{Amount: 175.95, DueDate: "2023-11-22T00:00:00Z", Status: "paid", UserID: "user-4", ProviderID: "provider-2", ProviderName: "Water Provider", PaidDate: "2023-12-15T00:00:00Z"},
	{Amount: 149.62, DueDate: "2024-01-10T00:00:00Z", Status: "pending", UserID: "user-15", ProviderID: "provider-2", ProviderName: "Water Provider"},
	{Amount: 129.85, DueDate: "2024-02-20T00:00:00Z", Status: "paid", UserID: "user-16", ProviderID: "provider-2", ProviderName: "Water Provider", PaidDate: "2024-03-10T00:00:00Z"},
	{Amount: 119.75, DueDate: "2024-04-18T00:00:00Z", Status: "paid", UserID: "user-17", ProviderID: "provider-2", ProviderName: "Water Provider", PaidDate: "2024-05-01T00:00:00Z"},
	{Amount: 83.20, DueDate: "2024-05-25T00:00:00Z", Status: "pending", UserID: "user-18", ProviderID: "provider-2", ProviderName: "Water Provider"},
	{Amount: 155.40, DueDate: "2024-06-14T00:00:00Z", Status: "paid", UserID: "user-19", ProviderID: "provider-2", ProviderName: "Water Provider", PaidDate: "2024-06-28T00:00:00Z"},
	{Amount: 112.65, DueDate: "2024-08-03T00:00:00Z", Status: "pending", UserID: "user-20", ProviderID: "provider-2", ProviderName: "Water Provider"},
	{Amount: 101.50, DueDate: "2024-10-07T00:00:00Z", Status: "paid", UserID: "user-21", ProviderID: "provider-2", ProviderName: "Water Provider", PaidDate: "2024-11-02T00:00:00Z"},
	{Amount: 189.90, DueDate: "2024-11-19T00:00:00Z", Status: "pending", UserID: "user-22", ProviderID: "provider-2", ProviderName: "Water Provider"},
	{Amount: 140.60, DueDate: "2025-01-25T00:00:00Z", Status: "paid", UserID: "user-23", ProviderID: "provider-2", ProviderName: "Water Provider", PaidDate: "2025-02-12T00:00:00Z"},
	{Amount: 174.20, DueDate: "2025-02-12T00:00:00Z", Status: "paid", UserID: "user-24", ProviderID: "provider-2", ProviderName: "Water Provider", PaidDate: "2025-02-20T00:00:00Z"},
	{Amount: 97.90, DueDate: "2025-03-02T00:00:00Z", Status: "pending", UserID: "user-25", ProviderID: "provider-2", ProviderName: "Water Provider"},
	{Amount: 138.00, DueDate: "2025-03-15T00:00:00Z", Status: "pending", UserID: "user-26", ProviderID: "provider-2", ProviderName: "Water Provider"},
	{Amount: 150.80, DueDate: "2025-03-22T00:00:00Z", Status: "paid", UserID: "user-27", ProviderID: "provider-2", ProviderName: "Water Provider", PaidDate: "2025-04-01T00:00:00Z"},
	{Amount: 102.50, DueDate: "2025-03-30T00:00:00Z", Status: "pending", UserID: "user-28", ProviderID: "provider-2", ProviderName: "Water Provider"},
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
	json.NewEncoder(w).Encode(internetBills)
}

func main() {
	http.HandleFunc("/api/v1/bills", authMiddleware(getBills))

	log.Println("Internet mock server running on http://localhost:5003")
	log.Fatal(http.ListenAndServe(":5003", nil))
}
