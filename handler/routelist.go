package handler

import (
	"github.com/godaner/go-route/route"
)

func Routes() {
	route.RegistRoutes(
		route.MakeAnyRoute("/head",HeadHandler),)
}

