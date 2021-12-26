package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"strconv"
)

type RedisConnectionParams struct {
	URL, Password, Host, Port, DB string
}

type Client *redis.Client

type Options *redis.Options

func convertStringToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Not parse int")
	}
	return result
}

func BuildRedisOptions(p *RedisConnectionParams) *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", p.Host, p.Port),
		Password: p.Password,
		DB:       convertStringToInt(p.DB),
	}
}

func CheckAlive(c *redis.Client) bool {
	_, err := c.Ping().Result()
	if err != nil {
		log.Fatal("Error connect to Redis")
		return false
	}
	return true
}

func Connect(p *redis.Options) *Client {
	var client Client = redis.NewClient(p)

	statusOK := CheckAlive(client)

	if !statusOK {
		panic("Redis not alive")
	}

	return &client
}
