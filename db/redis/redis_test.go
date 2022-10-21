package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"testing"
)

const password string = "ediieolol"
const address string = "10.0.0.73:5396"

func redisConnect() {
	dial, err := redis.Dial("tcp", address, redis.DialPassword(password))
	if err != nil {
		log.Fatalln("redis connect error")
	}
	defer dial.Close()
	_ = dial.Send("SET", "name", "zbc-test")
	result, _ := redis.String(dial.Do("GET", "name"))
	fmt.Println(result)
}

func TestRedisConnect(t *testing.T) {
	redisConnect()
}
