package redis_store

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
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
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       0,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	fmt.Println("REDIS PING OK")

	return &ProxyRedisStore{
		Client: client,
	}, nil
}

func (s *ProxyRedisStore) Save(k *models.ProxyRequest, v *models.ProxyResponse) error {
	fmt.Println("REDIS Save")

	ctx := context.TODO()

	return s.Client.HSet(ctx, k.URL, k, v).Err()
}

func (s *ProxyRedisStore) Get(k *models.ProxyRequest) (*models.ProxyResponse, error) {

	return nil, nil
}

func (s *ProxyRedisStore) All() (map[*models.ProxyRequest]*models.ProxyResponse, error) {
	ctx := context.TODO()

	res := make(map[*models.ProxyRequest]*models.ProxyResponse)

	// TODO: Get all keys, for range by keys, add struct to map, return map

	all := s.Client.HGetAll(ctx, "https://go.dev")
	if all.Err() != nil {
		return nil, all.Err()
	}

	rrr, err := all.Result()
	if err != nil {
		return nil, err
	}

	var key *models.ProxyRequest
	var value *models.ProxyResponse

	for rrrKey, rrrVal := range rrr {
		spew.Dump(rrrKey)
		spew.Dump(rrrVal)

		err = json.Unmarshal([]byte(rrrKey), &key)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(rrrVal), &value)
		if err != nil {
			return nil, err
		}

		res[key] = value
	}

	return res, nil
}
