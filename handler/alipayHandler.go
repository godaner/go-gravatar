package handler

import (
	"github.com/godaner/go-route/route"
	"github.com/godaner/go-util"
	"github.com/ungerik/go-gravatar"
	"log"
)

type GravatarResponse struct {

	Link string	`bson:link,json:link"`
}

func HeadHandler(response route.RouteResponse, request route.RouteRequest){
	email:=request.Params["email"].(string)
	log.Println("HeadHandler start ! email is : ",email)
	headUrl:=gravatar.SecureUrl(email)
	log.Println("HeadHandler make headUrl success ! email is : ",email," , headUrl is : ",headUrl)
	go_util.RedirectUrl(response.ResponseWriter,headUrl)
}
