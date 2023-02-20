package redis_store

import (
	"context"
	"github.com/redis/go-redis/v9"
	"proxy_map/Internal/domain/models"
)

type ProxyRedisStore struct {
	Client *redis.Client
}

// Верхний уровень - метод
// Слкдующий уровень - URL

func NewProxyRedisStore(ctx context.Context, redisAddr string) (*ProxyRedisStore, error) {
	// TODO: Move NewClient to app
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &ProxyRedisStore{
		Client: client,
	}, nil
}

func (s *ProxyRedisStore) Save(k *models.ProxyRequest, v *models.ProxyResponse) error {

	//err := s.Client.Do(context.TODO(), k, v).Err()

	return nil
}

func (s *ProxyRedisStore) Get(k *models.ProxyRequest) (*models.ProxyResponse, error) {

	return nil, nil
}

func (s *ProxyRedisStore) All() (map[*models.ProxyRequest]*models.ProxyResponse, error) {

	return nil, nil
}
