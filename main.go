package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", getTransactions)

	http.ListenAndServe(":8080", nil)
}

type Transaction struct {
	Title     string
	Amount    float32
	Type      int
	CreatedAt time.Time
}

type Transactions []Transaction

func getTransactions(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var layoutData = "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layoutData, "2024-04-09T15:59:48")

	var transaction = Transactions{
		Transaction{
			Title:     "Salario",
			Amount:    1250.90,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	json.NewEncoder(w).Encode(transaction)

}
