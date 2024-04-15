package routes

import (
	"net/http"

	"github.com/PkMs7/api-session-finance-dio/controllers"
)

func HandleRequest() {

	http.HandleFunc("/transactions", controllers.GetTransactions)
	http.HandleFunc("/transactions/create", controllers.CreateATransaction)

	http.ListenAndServe(":8080", nil)

}
