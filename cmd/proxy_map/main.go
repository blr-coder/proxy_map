package main

import (
	"context"
	"github.com/Shopify/sarama"
	"log"
	"proxy_map/Internal/controllers/rest"
	"proxy_map/Internal/domain/usecases"
	"proxy_map/Internal/infrastructure/repository/redis_store"
	"proxy_map/pkg/kafka"
)

func main() {
	if err := runProxy(); err != nil {
		log.Fatalln(err)
	}
}

func runProxy() error {

	ctx := context.TODO()

	//storage := map_store.NewProxyMap()

	storage, err := redis_store.NewProxyRedisStore(ctx, "localhost:6379")

	conn, err := kafka.Open(&kafka.Config{
		Config: sarama.NewConfig(),
		Details: map[string]*sarama.TopicDetail{
			"quickstart": {
				NumPartitions:     1,
				ReplicationFactor: 1,
			},
		},
		Addr: []string{"localhost:9092"},
	})
	if err != nil {
		return err
	}
	defer conn.Close()

	proxyUC := usecases.NewProxyUseCase(storage, conn)

	proxyHandler := rest.NewProxyHandler(proxyUC)

	server := rest.NewHTTPServer(proxyHandler)

	// TODO: CONFIGGGG!!!
	return server.Start("7777")
}
