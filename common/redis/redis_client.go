package redis

import (
	"context"
	"errors"
	"github.com/gomodule/redigo/redis"
	"log"
)

var Handler *RedisClient

type RedisClient struct{
	Client redis.Conn
}


func InitRedisHandler(sections []string) {
	handler, err := redis.Dial("tcp", "localhost:6379")
	if err != nil{
		log.Fatal("init redis client fail, err:"+ err.Error())
	}
	defer handler.Close()
	Handler = &RedisClient{
		Client: handler,
	}
}

func (r *RedisClient) Set(ctx context.Context, key, val string) error {
	_, err := r.Client.Do("SET", key, val)
    return err
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := redis.String(r.Client.Do("GET", key))
	if err !=nil{
		if errors.Is(err, redis.ErrNil) {
			return "", errors.New("key not found")
		}else{
			return "", errors.New(err.Error())
		}
	}
	
	return val, nil
}

