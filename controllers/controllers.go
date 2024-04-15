package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PkMs7/api-session-finance-dio/models"
	"github.com/PkMs7/api-session-finance-dio/utils"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transaction = models.Transactions{
		models.Transaction{
			Title:     "Salario",
			Amount:    1250.90,
			Type:      0,
			CreatedAt: utils.StringToDate("2024-02-10T15:59:48"),
		},
	}

	json.NewEncoder(w).Encode(transaction)

}

func CreateATransaction(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = models.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &res)

	fmt.Println(res)

}
