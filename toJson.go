package main

import (
	"encoding/json"
	_ "github.com/go-redis/redis"
	"log"
	"net/http"
)

func SetResponseJSON(w http.ResponseWriter, code int, source string, follower int, error string) {
	msg := Message {
		Status: code,
		Error: error,
		Follower: follower,
		Source: source,
	}
	b, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
