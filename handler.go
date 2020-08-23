package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func HandleMain(w http.ResponseWriter, r *http.Request)  {
	//start := time.Now()
	if _, ok := r.URL.Query()["username"]; !ok {
		SetResponseJSON(w, http.StatusBadRequest, "", 0, "Required username ")
		return
	}

	username := r.URL.Query()["username"][0]

	//Redis get
	val, err := RedisClient.Get(username).Result()
	if err == redis.Nil {
		fmt.Printf("key %s in redis not exists", username)
		url := fmt.Sprintf("https://api.github.com/users/%v/followers", username)
		response, err := http.Get(url)
		if err != nil {
			SetResponseJSON(w, http.StatusBadRequest, "", 0, "Account not found")
			return
		}

		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			SetResponseJSON(w, http.StatusInternalServerError, "", 0, "Error When get data")
			return

		}

		var followers []Follower
		err = json.Unmarshal(body, &followers)
		if err != nil {
			SetResponseJSON(w, http.StatusBadRequest, "", 0, "Error When parsing data")
			return
		}

		res := len(followers)
		SetResponseJSON(w, http.StatusOK,"GITHUBAPI", res, "none")
		err = RedisClient.Set(username, res, 60 * time.Second).Err()
		if err != nil {
			w.Write([]byte(`{"message":"error when save data to redis"}`))
			return
		}
		return
	}
	res, _ := strconv.Atoi(val)
	SetResponseJSON(w, http.StatusOK,"REDIS", res, "none")
	return

	//elapsed := time.Since(start).Seconds()
	//fmt.Println(elapsed)
}
