package main

import (
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

type Follower struct {
	Login string `json:"login"`
	ID    int    `json:"id"`
}

type Message struct {
	Error    string `json:"error"`
	Status   int    `json:"status"`
	Follower int    `json:"follower"`
	Source   string `json:"source"`
	Message  string `json:"message"`
}

var RedisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func main() {
	http.HandleFunc("/", HandleMain)
	log.Println(http.ListenAndServe(":9000", nil))
}
