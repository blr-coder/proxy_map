package main

import (
	"log"
	rest2 "proxy_map/Internal/controllers/rest"
	"proxy_map/Internal/infrastructure/repository/map_store"
)

func main() {
	if err := runProxy(); err != nil {
		log.Fatalln(err)
	}
}

func runProxy() error {
	storage := map_store.NewProxyMap()

	proxyHandler := rest2.NewProxyHandler(storage)

	server := rest2.NewHTTPServer(proxyHandler)

	return server.Start("7777")
}
