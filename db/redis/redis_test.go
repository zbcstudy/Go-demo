package redis

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

const password string = "ediieolol"
const address string = "10.0.0.73:5396"

func redisConnect() {
	dial, err := redis.Dial("tcp", address, redis.DialPassword(password))
	if err != nil {
		log.Fatalln("redis connect error")
	}
	defer dial.Close()

	dial.Do("Get", "name")

}
