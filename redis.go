package redis

import (
	"fmt"
	r "github.com/go-redis/redis"
	"log"
	"strconv"
)

type Redis struct {
	URL, Password, Host, Port, DB string
	Params                        *r.Options
	Client                        *r.Client
}

func convertStringToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Not parse int")
	}
	return result
}

func (redis *Redis) Options() *r.Options {
	return &r.Options{
		Addr:     fmt.Sprintf("%s:%s", redis.Host, redis.Port),
		Password: redis.Password,
		DB:       convertStringToInt(redis.DB),
	}
}

func (redis *Redis) Alive() bool {
	_, err := redis.Client.Ping().Result()
	if err != nil {
		log.Fatal("Error connect to Redis")
		return false
	}
	return true
}

func (redis *Redis) Connect() *r.Client {
	return r.NewClient(redis.Params)
}
