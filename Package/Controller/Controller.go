package Controller

import (
	"net/http"

	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Logger"
	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Model"
)

type CentralController struct {
	//DBProcessor    Model.DBProcessor
	DatabaseInterface Model.DatabaseInterface
	RedisInterface    Model.RedisInterface
	//	RDBProcessor Model.RdbProcessor//
	Logger Logger.LoggerInterface
}

func NewCentralController(DBProc Model.DatabaseInterface, RDBProc Model.RedisInterface, logger Logger.LoggerInterface) *CentralController {
	return &CentralController{
		DatabaseInterface: DBProc,
		RedisInterface:    RDBProc,
		Logger:            logger,
	}
}

func (CentralControl *CentralController) AddUser(Writer http.ResponseWriter, Req *http.Request) {

}

func (CentralControl *CentralController) EditUser(Writer http.ResponseWriter, Req *http.Request) {}

func (CentralControl *CentralController) EditUserCred(Writer http.ResponseWriter, Req *http.Request) {
}

func (CentralControl *CentralController) DeleteUser(Writer http.ResponseWriter, Req *http.Request) {}
