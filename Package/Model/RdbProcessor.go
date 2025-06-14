package Model

import (
	"sync"

	"github.com/FACELESS-GOD/EcommerceMonkay.git/Package/Utility"
	"github.com/redis/go-redis/v9"
)

type RedisInterface interface {
	AddJWT()
	AddCred()
	ValidateCred()
} 


type RdbProcessor struct {
	RdbProc *redis.Client
}

func (Rdb *RdbProcessor) NewRdbProcessor(Wg *sync.WaitGroup) *RdbProcessor {
	defer Wg.Done()

	Rdb.RdbProc = Utility.RdInitiator()

	return Rdb
}

func (Rdb *RdbProcessor) AddJWT() {}

func (Rdb *RdbProcessor) AddCred() {}

func (Rdb *RdbProcessor) ValidateCred() {}
