package main

import (
	"flag"
	"log"

	engine "github.com/muidea/magicEngine"
	magicprotal "github.com/muidea/magicProtal/core"
)

var bindPort = "8866"
var centerServer = "127.0.0.1:8888"
var endpointName = "magicProtal"
var endpointID = "8a30313b-50ed-4a3a-b52b-e0d3e000d400"
var authToken = "1xlUfmwNZ6nIsIpDUTax17TbgpaSI1Xj"

func main() {
	flag.StringVar(&bindPort, "ListenPort", bindPort, "magicProtal listen address")
	flag.StringVar(&centerServer, "CenterSvr", centerServer, "magicCenter server")
	flag.StringVar(&endpointName, "EndpointName", endpointName, "magicProtal endpoint name.")
	flag.StringVar(&endpointID, "EndpointID", endpointID, "magicProtal endpoint id")
	flag.StringVar(&authToken, "AuthToken", authToken, "magicProtal authtoken")
	flag.Parse()

	router := engine.NewRouter()

	protal, ok := magicprotal.New(centerServer, endpointName, endpointID, authToken)
	if ok {
		protal.Startup(router)

		svr := engine.NewHTTPServer(bindPort)
		svr.Bind(router)

		svr.Run()
	} else {
		log.Printf("new Protal failed.")
	}

	protal.Teardown()
}
