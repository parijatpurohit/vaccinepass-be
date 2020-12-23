package logic

import (
	"github.com/parijatpurohit/vaccinepass/server/transport"
	"github.com/parijatpurohit/vaccinepass/zz_generated/go/server"
	"github.com/parijatpurohit/vaccinepass/zz_generated/go/server/handlers"
	"sync"
)

const (
	address = "localhost:50010"
)

type ILogic interface {
	GetUserDetails(req *transport.GetUserDetailsRequest) (*transport.GetUserDetailsResponse, error)
	GetUserVaccineDetails(req *transport.GetUserVaccineDetailsRequest) (*transport.GetUserVaccineDetailsResponse, error)
	GetRequiredVaccines(req *transport.GetRequiredVaccineRequest) (*transport.GetRequiredVaccineResponse, error)
}

var (
	logicClient *logicImpl
	once        sync.Once
)

type logicImpl struct {
	address string
	client  *handlers.Service
}

func initialise() {
	logicClient = &logicImpl{ client : server.LocalService}
}

func GetClient() ILogic {
	once.Do(func() {
		initialise()
	})
	return logicClient
}
