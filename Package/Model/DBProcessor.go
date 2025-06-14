package Model

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	Cred "github.com/FACELESS-GOD/EcommerceMonkay.git/Helper/CredStore"
	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Utility"
)

type DatabaseInterface interface {
	AddUser(http.ResponseWriter, *http.Request)
	EditUser(http.ResponseWriter, *http.Request)
	EditUserCred(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
}

type DBProcessor struct {
	DB *sql.DB
}

func (DBProc *DBProcessor) NewDBProcessor(Wg *sync.WaitGroup) *DBProcessor {

	defer Wg.Done()

	db, err := Utility.InitiateDBInstance(Cred.DBConnString)

	if err != nil {
		fmt.Println(err)
	}

	DBProc.DB = db

	return DBProc
}

func (DBProc *DBProcessor) AddUser(Writer http.ResponseWriter, Req *http.Request) {}

func (DBProc *DBProcessor) EditUser(Writer http.ResponseWriter, Req *http.Request) {}

func (DBProc *DBProcessor) EditUserCred(Writer http.ResponseWriter, Req *http.Request) {}

func (DBProc *DBProcessor) DeleteUser(Writer http.ResponseWriter, Req *http.Request) {}
