package main

import (
	"github.com/Shopify/sarama"
	"log"
	"proxy_map/Internal/controllers/rest"
	"proxy_map/Internal/domain/usecases"
	"proxy_map/Internal/infrastructure/repository/map_store"
	"proxy_map/pkg/events"
)

func main() {
	if err := runProxy(); err != nil {
		log.Fatalln(err)
	}
}

func runProxy() error {

	storage := map_store.NewProxyMap()

	// TODO: Add func for kafka like "setupSender(config *Config) error"
	saramaConfig := sarama.NewConfig()
	saramaConfig.Consumer.Return.Errors = true
	// TODO: CONFIGGGG!!!
	saramaClient, err := sarama.NewClient([]string{"localhost:9092"}, saramaConfig)
	if err != nil {
		return err
	}
	producer, err := sarama.NewSyncProducerFromClient(saramaClient)
	if err != nil {
		return err
	}

	sender := events.NewKafkaSender(producer)

	proxyUC := usecases.NewProxyUseCase(storage, sender)

	proxyHandler := rest.NewProxyHandler(proxyUC)

	server := rest.NewHTTPServer(proxyHandler)

	// TODO: CONFIGGGG!!!
	return server.Start("7777")
}
