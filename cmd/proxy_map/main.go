package main

import (
	"proxy_map/core"
	"proxy_map/transport/rest"
)

func main() {

	storage := core.NewProxyMap()

	proxyHandler := rest.NewProxyHandler(storage)

	server := rest.NewHTTPServer(proxyHandler)
	err := server.Start("7777")
	if err != nil {
		panic(err)
	}

}
