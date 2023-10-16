package grpc_client

import (
	"fmt"

	"gitlab.com/clinic-crm/reception/config"
	pbd "gitlab.com/clinic-crm/reception/genproto/doctor"
	pbl "gitlab.com/clinic-crm/reception/genproto/lab"
	pbp "gitlab.com/clinic-crm/reception/genproto/patient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	PatientService() pbp.PatientServiceClient
	DoctorService() pbd.DoctorServiceClient
	LabService() pbl.LabServiceClient
}

type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (ServiceManager, error) {
	connPatientService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.PatientServiceHost, cfg.PatientServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("patient service dial host: %s port:%s err: %v",
			cfg.PatientServiceHost, cfg.PatientServicePort, err)
	}

	connDoctorService, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.DoctorServiceHost, cfg.DoctorServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("doctor service dial host: %s port:%s err: %v",
			cfg.PatientServiceHost, cfg.DoctorServicePort, err)
	}

	connLabService, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.LabServiceHost, cfg.LabServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("lab service dial host: %s port:%s err: %v",
			cfg.PatientServiceHost, cfg.LabServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"patient_service": pbp.NewPatientServiceClient(connPatientService),
			"doctor_service":  pbd.NewDoctorServiceClient(connDoctorService),
			"lab_service":     pbl.NewLabServiceClient(connLabService),
		},
	}, nil
}

func (g *GrpcClient) PatientService() pbp.PatientServiceClient {
	return g.connections["patient_service"].(pbp.PatientServiceClient)
}

func (g *GrpcClient) DoctorService() pbd.DoctorServiceClient {
	return g.connections["doctor_service"].(pbd.DoctorServiceClient)
}

func (g *GrpcClient) LabService() pbl.LabServiceClient {
	return g.connections["lab_service"].(pbl.LabServiceClient)
}
