package main

import (
	"log"

	micro "github.com/micro/go-micro"

	"github.com/pijalu/go.hands.two/frinsultdata/handler"
	"github.com/pijalu/go.hands.two/frinsultproto"

	// Register k8s specific
	_ "github.com/micro/go-plugins/registry/kubernetes"
	_ "github.com/micro/go-plugins/selector/static"
)

func main() {
	service := micro.NewService(
		micro.Name("frinsult.srv.micro"),
		micro.Version("latest"),
	)

	// Register
	frinsultproto.RegisterFrinsultServiceHandler(
		service.Server(),
		new(handler.FrinsultHandler))

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
