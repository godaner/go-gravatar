package main

import (
	"github.com/godaner/go-route/route"
	"go-gravatar/handler"
)

const(
	ADDR=":80"
)



func main() {

	//routes
	handler.Routes()

	//run
	route.Start(ADDR)
}

