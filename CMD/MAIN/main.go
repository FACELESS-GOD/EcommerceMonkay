package main

import (
	"net/http"
	"sync"

	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Controller"
	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Middleware"
	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Model"
	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Router"
	"github.com/gorilla/mux"
)

func main() {

	logger := Middleware.NewLoggerDev()

	waitgroup := sync.WaitGroup{}

	dbInst := Model.DBProcessor{}

	rdbInst := Model.RdbProcessor{}

	go dbInst.NewDBProcessor(&waitgroup)
	waitgroup.Add(1)

	go rdbInst.NewRdbProcessor(&waitgroup)
	waitgroup.Add(1)

	waitgroup.Wait()

	centralControl := Controller.CentralController{
		DBProcessor:  dbInst,
		RDBProcessor: rdbInst,
	}

	muxRouter := mux.NewRouter()

	Router.CustomRouter(muxRouter, &centralControl, logger)

	http.Handle("/", muxRouter)

	http.ListenAndServe(":8080", muxRouter)

}
