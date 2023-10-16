package services

import (
	"fmt"

	"gitlab.com/clinic-crm/api-gateway/config"

	d "gitlab.com/clinic-crm/api-gateway/genproto/doctor"
	l "gitlab.com/clinic-crm/api-gateway/genproto/lab"
	p "gitlab.com/clinic-crm/api-gateway/genproto/patient"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	PatientService() p.PatientServiceClient
	DoctorService() d.DoctorServiceClient
	LabService() l.LabServiceClient
}

type serviceManager struct {
	patientService p.PatientServiceClient
	doctorService  d.DoctorServiceClient
	labService     l.LabServiceClient
}

func (s *serviceManager) PatientService() p.PatientServiceClient {
	return s.patientService
}

func (s *serviceManager) DoctorService() d.DoctorServiceClient {
	return s.doctorService
}

func (s *serviceManager) LabService() l.LabServiceClient {
	return s.labService
}

func NewServiceManager(conf config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connPatient, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PatientServiceHost, conf.PatientServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connDoctor, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.DoctorServiceHost, conf.DoctorServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connLab, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.LabServiceHost, conf.LabServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Error at services/services.go 64  %v\n", err)
		return nil, err
	}

	serviceManager := &serviceManager{
		patientService: p.NewPatientServiceClient(connPatient),
		doctorService:  d.NewDoctorServiceClient(connDoctor),
		labService:     l.NewLabServiceClient(connLab),
	}

	return serviceManager, nil
}
