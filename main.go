package main

import (
	"fmt"
	"github.com/heck-go/heck"
)

func main() {
	server := heck.NewServer(":15000")
	server.GetFor("/api/hello", func(ctx *heck.Context) {
		ctx.Response = &heck.Response{
			StatusCode: 200,
			Value:      "Hello stranger!",
		}
	}, nil).Before(func(ctx *heck.Context) {
		fmt.Println("Received route: ", ctx.Method())
		ctx.Before(func(ctx *heck.Context) {
			fmt.Println("Path: ", ctx.Path())
		})
	})
	server.GetFor("/api/hello/:name", func(ctx *heck.Context) {
		name, _ := ctx.PathParams.Get("name")
		ctx.Response = &heck.Response{
			StatusCode: 200,
			Value:      "Hello " + name + "!",
		}
	}, nil)
	server.GetFor("/api/math/add", func(ctx *heck.Context) {
		a, _, _ := ctx.Query.Int("a")
		b, _, _ := ctx.Query.Int("b")
		ctx.Response = &heck.Response{
			StatusCode: 200,
			Value:      a + b,
		}
	}, nil)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
