package Controller

import (
	"net/http"

	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Model"
)

type CentralController struct {
	DBProcessor  Model.DBProcessor
	RDBProcessor Model.RdbProcessor
}

func NewCentralController(DBProc *Model.DBProcessor, RDBProc *Model.RdbProcessor) *CentralController {
	return &CentralController{
		DBProcessor:  *DBProc,
		RDBProcessor: *RDBProc,
	}
}

func (CentralControl *CentralController) AddUser(Writer http.ResponseWriter, Req *http.Request) {

}

func (CentralControl *CentralController) EditUser(Writer http.ResponseWriter, Req *http.Request) {}

func (CentralControl *CentralController) EditUserCred(Writer http.ResponseWriter, Req *http.Request) {
}

func (CentralControl *CentralController) DeleteUser(Writer http.ResponseWriter, Req *http.Request) {}
