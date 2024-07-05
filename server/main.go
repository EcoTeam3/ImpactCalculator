package main

import (
	"impactcalculator/config"
	pb "impactcalculator/generated/impact"
	"impactcalculator/service"
	"impactcalculator/storage"
	"impactcalculator/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := storage.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	config := config.Load()
	listener, err := net.Listen("tcp", ":"+config.URL_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	s := service.NewImpactRepo(*postgres.NewImpactRepo(db))
	grpc := grpc.NewServer()
	pb.RegisterImpactServer(grpc, s)
	if err = grpc.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
