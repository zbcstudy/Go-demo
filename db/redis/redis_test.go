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

	reply, _ := dial.Do("HSET", "person", "name", "zbc")
	log.Println("redis hset reply", reply)
	reply2, _ := dial.Do("HSET", "person", "age", 18)
	log.Println("redis hset reply", reply2)

	response, _ := redis.String(dial.Do("HGet", "person", "name"))
	log.Println("redis hget return:", response)
	response2, _ := redis.Int(dial.Do("HGet", "person", "age"))
	log.Println("redis hget return:", response2)

	fmt.Println(result)
}

// redis pool connect
func redisPoolConnect() {
	pool := &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", address, redis.DialPassword(password))
		},
	}
	conn := pool.Get()
	defer conn.Close()
	reply, err := conn.Do("PING")
	if err != nil {
		fmt.Println("ping error")
	}
	fmt.Println(reply)
}

func TestRedisConnect(t *testing.T) {
	redisConnect()
}

func TestRedisPoolConnect(t *testing.T) {
	redisPoolConnect()
}
