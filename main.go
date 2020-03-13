package main

import (
	"fmt"
	"log"
	"net/http"
	"tinycache/lru"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	lru.NewGroup("scores", 2<<10, lru.GetterFunc(
		func(key string) ([]byte, error) {
			log.Printf("[slowdb] search key", key)

			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"

	peers := lru.NewHTTPPool(addr)
	log.Printf("cache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
