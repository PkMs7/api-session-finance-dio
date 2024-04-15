package routes

import (
	"net/http"

	"github.com/PkMs7/api-session-finance-dio/controllers"
	"github.com/PkMs7/api-session-finance-dio/middlewares"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func HandleRequest() {

	http.HandleFunc("/health", middlewares.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/transactions", controllers.GetTransactions)
	http.HandleFunc("/transactions/create", controllers.CreateATransaction)

	http.ListenAndServe(":8080", nil)

}
