package store

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type StorageService interface {
	Connect() error
	Save(k string, val any) error
	Fetch(k string) (any, error)
}

type Redis struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewRedis(ctx context.Context) *Redis {
	return &Redis{
		redisClient: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		}),
		ctx: ctx,
	}
}

func (r *Redis) Connect() error {
	result, err := r.redisClient.Ping(r.ctx).Result()
	if err != nil {
		return err
	}
	log.Println("Successfully Connected storage:", result)
	return nil
}

func (r *Redis) Save(k string, val any) error {
	val, err := r.redisClient.SetEx(r.ctx, k, val, time.Hour*24*365*5).Result()
	if err != nil {
		return err
	}
	log.Println("save command value:", val)
	return nil
}

func (r *Redis) Fetch(k string) (any, error) {
	cmd := r.redisClient.Get(r.ctx, k)
	return cmd.Result()
}
