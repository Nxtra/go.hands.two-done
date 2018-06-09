package main

import (
	"log"

	micro "github.com/micro/go-micro"
	"github.com/pijalu/go.hands.two/frinsultdata/handler"
	"github.com/pijalu/go.hands.two/frinsultdata/repository"
	"github.com/pijalu/go.hands.two/frinsultproto"
)

func main() {
	// Close DB at end of exec
	defer func() {
		if repository.DB != nil {
			repository.DB.Close()
		}
	}()

	service := micro.NewService(
		micro.Name("frinsult.micro.srv.micro"),
		micro.Version("latest"),
	)

	// Register
	frinsultproto.RegisterFrinsultServiceHandler(service.Server(), new(handler.FrinsultHandler))

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
