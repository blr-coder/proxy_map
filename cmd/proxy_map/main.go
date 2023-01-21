package main

import (
	"log"
	"proxy_map/Internal/controllers/rest"
	"proxy_map/Internal/domain/usecases"
	"proxy_map/Internal/infrastructure/repository/map_store"
)

func main() {
	if err := runProxy(); err != nil {
		log.Fatalln(err)
	}
}

func runProxy() error {
	storage := map_store.NewProxyMap()

	proxyUC := usecases.NewProxyUseCase(storage)

	proxyHandler := rest.NewProxyHandler(proxyUC)

	server := rest.NewHTTPServer(proxyHandler)

	return server.Start("7777")
}
