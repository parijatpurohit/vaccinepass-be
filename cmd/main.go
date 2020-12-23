package main

import (
	"github.com/parijatpurohit/vaccinepass/server"
	server2 "github.com/parijatpurohit/vaccinepass/zz_generated/go/server"
)

func main() {
	go server2.Serve()
	server.Serve()
}
