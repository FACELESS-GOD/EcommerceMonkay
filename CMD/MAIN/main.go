package main

import (
	"net/http"
	"sync"

	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Controller"
	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Logger"
	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Middleware"
	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Model"
	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Router"
	"github.com/gorilla/mux"
)

func main() {

	logger := Logger.NewLoggerDev()
	//logger := Logger.NewLoggerProd()

	logMiddleWare := Middleware.NewCustomLogger(logger)

	waitgroup := sync.WaitGroup{}

	dbInst := Model.DBProcessor{}

	rdbInst := Model.RdbProcessor{}

	go dbInst.NewDBProcessor(&waitgroup)
	waitgroup.Add(1)

	go rdbInst.NewRdbProcessor(&waitgroup)
	waitgroup.Add(1)

	waitgroup.Wait()

	centralControl := Controller.CentralController{
		DatabaseInterface: &dbInst,
		RedisInterface:    &rdbInst,
	}

	muxRouter := mux.NewRouter()

	Router.CustomRouter(muxRouter, &centralControl, logMiddleWare)

	http.Handle("/", muxRouter)

	http.ListenAndServe(":8080", muxRouter)

}
