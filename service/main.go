package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/wataly/micro/service/handler"

	user "github.com/wataly/micro/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("top.hfjy.service.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler, ignore exception
	_ = user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
