package main

import (
	server2 "github.com/parijatpurohit/vaccinepass/server"
	"github.com/parijatpurohit/vaccinepass/zz_generated/go/server"
)

func main() {
	go server.Serve()
	server2.Serve()
}
