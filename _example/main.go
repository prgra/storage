package main

import (
	"fmt"

	mem "github.com/prgra/storage/mem"
	redis "github.com/prgra/storage/redis"
)

func main() {
	mem.Push("test", "hohoho")
	v, ok := mem.Get("test")
	if ok {
		fmt.Println(v)
	}
	redis.New("127.0.0.1:6379", "", 0)
	err := redis.Push("hello", "World")
	if err != nil {
		panic(err)
	}

	i, err := redis.Get("hello")
	if err != nil {
		panic(err)
	}
	fmt.Print(i)
}
