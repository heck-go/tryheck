package main

import (
	"fmt"
	"github.com/heck-go/heck"
)

func main() {
	server := heck.NewServer(":10000")
	server.GetFor("/api/name", func(ctx *heck.Context) {
		ctx.Response = heck.Response{
			StatusCode: 200,
			Value:      "Teja",
		}
	}, nil).Before(func(ctx *heck.Context) {
		fmt.Println("Received route: ", ctx.Method())
	})
	if err := server.Start(); err != nil {
		panic(err)
	}
}
