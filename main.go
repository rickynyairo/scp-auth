package main

import (
	"fmt"
	"github.com/micro/go-micro"
	// etcd "github.com/micro/go-micro/registry/etcd"
	pb "github.com/rickynyairo/scp-auth/proto/auth"
	"log"
)

func main() {

	// Creates a database connection and handles
	// closing it again before exit.
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(&pb.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("auth"),
		micro.Version("latest"),
		// micro.Registry(etcd.NewRegistry()),
	)

	// Init will parse the command line flags.
	srv.Init()
	publisher := micro.NewPublisher("user.created", srv.Client())
	// // pubsub := srv.Server().Options().Broker
	// client := pb.NewUserServiceClient("go.micro.api.greeter", srv.Client())
	// srv.Server().Handle(
	// 	srv.Server().NewHandler(
	// 		&Handler{Client: client},
	// 	),
	// )
	// Register handler
	pb.RegisterAuthHandler(srv.Server(), &Handler{repo, tokenService, publisher})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
