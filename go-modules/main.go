package main

import (
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logger, _ := zap.NewProduction()
	logger.Info("successfully performed http request")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
