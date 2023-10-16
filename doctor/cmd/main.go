package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	pb "gitlab.com/clinic-crm/doctor/genproto/doctor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"gitlab.com/clinic-crm/doctor/config"
	"gitlab.com/clinic-crm/doctor/service"
	"gitlab.com/clinic-crm/doctor/storage"

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

	patientService := service.NewDoctorService(strg)
	lis, err := net.Listen("tcp", ":"+cfg.DoctorServicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterDoctorServiceServer(s, patientService)

	log.Println("doctor service started in port ", cfg.DoctorServicePort,)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while listening: %v", err)
	}
}
