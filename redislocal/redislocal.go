package redislocal

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	Client *redis.Client
	client *redis.Client
	once   sync.Once
	err    error
	ctx    = context.Background()
)

func GetInstance(address, password string, db int) *redis.Client {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       db,
		})
	})

	return client
}

func InitClient(address, password string, db int) error {
	Client = GetInstance(address, password, db)
	statusCmd := Client.Ping(ctx)
	fmt.Printf("redis ping result: %s, %s, %d, %s\r\n",
		address, password, db, statusCmd.String())

	return statusCmd.Err()
}

func InitDefaultClient() error {
	return InitClient("localhost:6379", "", 0)
}
