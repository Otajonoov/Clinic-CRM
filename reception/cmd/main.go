package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	pb "gitlab.com/clinic-crm/reception/genproto/patient"
	"gitlab.com/clinic-crm/reception/pkg/grpc_client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"gitlab.com/clinic-crm/reception/config"
	"gitlab.com/clinic-crm/reception/service"
	"gitlab.com/clinic-crm/reception/storage"
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

	grpcClient, _ := grpc_client.New(cfg)

	strg := storage.NewStoragePg(psqlConn)

	patientService := service.NewPatientService(strg, &grpcClient)
	lis, err := net.Listen("tcp", ":"+cfg.PatientServicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterPatientServiceServer(s, patientService)

	log.Println("patient service started in port ", cfg.PatientServicePort, "host ", cfg.PatientServiceHost)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while listening: %v", err)
	}
}
