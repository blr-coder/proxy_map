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
		Addr:     "6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
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
	ctx := context.TODO()
	redisTestValue := struct {
		*models.ProxyRequest
		*models.ProxyResponse
	}{
		k,
		v,
	}

	//   - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
	/*err := s.Client.HSet(ctx, k.URL, map[*models.ProxyRequest]*models.ProxyResponse{k: v}).Err()
	if err != nil {
		return err
	}*/

	//   - HSet("myhash", MyHash{"value1", "value2"})
	return s.Client.HSet(ctx, k.URL, redisTestValue).Err()
}

func (s *ProxyRedisStore) Get(k *models.ProxyRequest) (*models.ProxyResponse, error) {

	return nil, nil
}

func (s *ProxyRedisStore) All() (map[*models.ProxyRequest]*models.ProxyResponse, error) {
	ctx := context.TODO()

	res := make(map[*models.ProxyRequest]*models.ProxyResponse)

	// TODO: Get all keys, for range by keys, add struct to map, return map

	var redisTestValue struct {
		*models.ProxyRequest
		*models.ProxyResponse
	}
	err := s.Client.HGetAll(ctx, "https://stackoverflow.com").Scan(&redisTestValue)
	if err != nil {
		return nil, err
	}
	res[redisTestValue.ProxyRequest] = redisTestValue.ProxyResponse

	return res, nil
}
