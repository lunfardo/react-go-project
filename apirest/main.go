package main

import (
	"apirest/server"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

type post struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

type myredisclient redis.Client

func fetchStockData() {
	var posts []post
	resp, _ := http.Get("https://jsonplaceholder.typicode.com/posts")
	readitall, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(readitall))
	err := json.Unmarshal(readitall, &posts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(posts[0].Title)
	fmt.Println("done")
}

func ExampleNewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "172.17.0.2:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	return client
}

func test(client *redis.Client) {
	err := client.Set("key", "value", 0).Err()
	err = client.Set("key2", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func main() {
	//client := ExampleNewClient()
	//test(client)
	server.Serve()
	//fetchStockData()
	//scheduler.AddTask(fetchStockData, 1)
	//go scheduler.AddTask(fetchStockData2, 2)
	//time.Sleep(time.Second * 100000)
}
