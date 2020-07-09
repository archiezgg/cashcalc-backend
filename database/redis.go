/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package database

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

var (
	redisURL = os.Getenv("REDIS_URL")
	client   *redis.Client
)

//StartupRedis initiates conenction to Redis and returns with the client
func StartupRedis() *redis.Client {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalln("error parsing URL of Redis: ", err)
	}

	client = redis.NewClient(opt)
	_, err = client.Ping().Result()
	if err != nil {
		log.Fatalln("error connecting to Redis: ", err)
	}

	log.Println("successfully connected to Redis!")
	return client
}

// RedisClient exposes the current connection to Redis
func RedisClient() *redis.Client {
	return client
}
