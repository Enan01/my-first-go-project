package main

import (
	"fmt"
	"github.com/alphazero/Go-Redis"
)

func main() {
	spec := redis.DefaultSpec().Host("localhost").Port(6379).Db(0).Password("admin")

	client, err := redis.NewSynchClientWithSpec(spec)

	if err != nil {
		fmt.Println("Connect redis server fail")
		return
	}

	dbkey := "session:392"

	getValue, err := client.Get(dbkey)

	if err != nil {
		fmt.Println("get key fail")
		return
	} else {
		str := string(getValue)
		fmt.Println(str)
	}

}
