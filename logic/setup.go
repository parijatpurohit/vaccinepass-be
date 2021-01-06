package logic

import (
	"sync"

	"github.com/parijatpurohit/vaccinepass/server/transport"
	"github.com/parijatpurohit/vaccinepass/zz_generated/go/server"
	"github.com/parijatpurohit/vaccinepass/zz_generated/go/server/handlers"
)

const (
	address = "localhost:50010"
)

type ILogic interface {
	GetUserSummary(req *transport.GetUserSummaryRequest) (*transport.GetUserSummaryResponse, error)
	GetUserVaccineDetails(req *transport.GetUserVaccineDetailsRequest) (*transport.GetUserVaccineDetailsResponse, error)
	GetCountryVaccines(req *transport.GetCountryVaccinesRequest) (*transport.GetCountryVaccinesResponse, error)
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
	logicClient = &logicImpl{client: server.LocalService}
}

func GetClient() ILogic {
	once.Do(func() {
		initialise()
	})
	return logicClient
}
