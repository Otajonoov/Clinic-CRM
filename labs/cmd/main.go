package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gitlab.com/clinic-crm/labs/config"
	"gitlab.com/clinic-crm/labs/genproto/lab"
	"gitlab.com/clinic-crm/labs/service"
	"gitlab.com/clinic-crm/labs/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	cfg := config.Load()

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	strg := storage.NewStoragePg(psqlConn)

	labService := service.NewLabService(strg)
	lis, err := net.Listen("tcp", ":"+cfg.LabServicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	lab.RegisterLabServiceServer(s, labService)

	log.Println("lab service started in port ", cfg.LabServicePort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while listening: %v", err)
	}
}
