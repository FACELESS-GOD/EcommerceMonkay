package Router

import (
	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Controller"
	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Middleware"
	"github.com/gorilla/mux"
)

func CustomRouter(MuxRouter *mux.Router, CentralController *Controller.CentralController, Logger *Middleware.CustomLogger) {

	MuxRouter.Use(Logger.LoggingMiddleware)

	MuxRouter.HandleFunc("/AddUser", CentralController.AddUser).Methods("POST")
	MuxRouter.HandleFunc("/EditUser", CentralController.EditUser).Methods("PUT")
	MuxRouter.HandleFunc("/EditUserCred", CentralController.EditUserCred).Methods("PUT")
	MuxRouter.HandleFunc("/DeleteUser", CentralController.EditUserCred).Methods("DELTEUser")

}
