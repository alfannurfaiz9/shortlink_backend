package config

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type Rdb struct {
	User string
	Pass string
	Host string
	Port string
}

func NewRdb(user, pass, host, port string) *Rdb {
	return &Rdb{
		User: user,
		Pass: pass,
		Host: host,
		Port: port,
	}
}

func (r *Rdb) Connect() *redis.Client {
	redis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.Host, r.Port),
		Username: r.User,
		Password: r.Pass,
	})

	if res, err := redis.Ping(context.Background()).Result(); err != nil {
		log.Printf("Close Redis : %s\n\n", err.Error())
	} else {
		log.Printf("%s ! Conected redis successfully\n\n", res)
	}

	return redis
}
