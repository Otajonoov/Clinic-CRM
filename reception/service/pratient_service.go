package service

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"gitlab.com/clinic-crm/reception/genproto/doctor"
	"gitlab.com/clinic-crm/reception/genproto/lab"
	"gitlab.com/clinic-crm/reception/genproto/patient"
	"gitlab.com/clinic-crm/reception/pkg/grpc_client"
	"gitlab.com/clinic-crm/reception/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PatientService struct {
	service grpc_client.ServiceManager
	storage storage.StorageI
	patient.UnimplementedPatientServiceServer
}

func NewPatientService(strg storage.StorageI, service *grpc_client.ServiceManager) *PatientService {
	return &PatientService{
		service:                           *service,
		storage:                           strg,
		UnimplementedPatientServiceServer: patient.UnimplementedPatientServiceServer{},
	}
}

func (s *PatientService) PatientCreate(ctx context.Context, req *patient.Patient) (*patient.Patient, error) {
	resp, err := s.storage.Patient().PatientCreate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.Patient{}, status.Error(codes.NotFound, "something went wrong, please not found this patient")
		}
		return &patient.Patient{}, status.Error(codes.Internal, "something went wrong, please check patient info")
	}

	return resp, nil
}

func (s *PatientService) PatientGet(ctx context.Context, req *patient.GetPatientReq) (*patient.Patient, error) {
	resp, err := s.storage.Patient().PatientGet(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.Patient{}, status.Error(codes.NotFound, "something went wrong, please not found this patient")
		}
		return &patient.Patient{}, status.Error(codes.Internal, "something went wrong, please check patient info")
	}

	return resp, nil
}

func (s *PatientService) PatientUpdate(ctx context.Context, req *patient.Patient) (*patient.Patient, error) {
	resp, err := s.storage.Patient().PatientUpdate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.Patient{}, status.Error(codes.NotFound, "something went wrong, please not found this patient")
		}
		return &patient.Patient{}, status.Error(codes.Internal, "something went wrong, please check patient info")
	}

	return resp, nil
}

func (s *PatientService) PatientDelete(ctx context.Context, req *patient.PatientId) (*empty.Empty, error) {
	err := s.storage.Patient().PatientDelete(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "something went wrong, please not found this patient")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "something went wrong, please check patient info")
	}

	return &emptypb.Empty{}, nil
}

func (s *PatientService) PatientsFind(ctx context.Context, req *patient.PatientsFindReq) (*patient.PatientsResp, error) {
	resp, err := s.storage.Patient().PatientsFind(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.PatientsResp{}, status.Error(codes.NotFound, "something went wrong, please not found this patient")
		}
		return &patient.PatientsResp{}, status.Error(codes.Internal, "something went wrong, please check patient info")
	}

	return resp, nil
}

func (s *PatientService) PatientsGetInfo(ctx context.Context, req *patient.PatientsGetInfoFilter) (*patient.PatientsResp, error) {
	resp, err := s.storage.Patient().PatientsGetInfo(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.PatientsResp{}, status.Error(codes.NotFound, "something went wrong, please not found this patient")
		}
		return &patient.PatientsResp{}, status.Error(codes.Internal, "something went wrong, please check patient info")
	}

	return resp, nil
}

// Patient analusiss
func (s *PatientService) PatientAnalysisCreate(ctx context.Context, req *patient.AnalysisInfo) (*patient.AnalysisInfo, error) {
	resp, err := s.storage.Patient().PatientAnalysisCreate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.AnalysisInfo{}, status.Error(codes.NotFound, "something went wrong, please not found this analysis")
		}
		return &patient.AnalysisInfo{}, status.Error(codes.Internal, "something went wrong, please check analysis info")
	}

	return resp, nil
}

func (s *PatientService) PatientAnalysisGet(ctx context.Context, req *patient.PatientPhoneNumber) (*patient.AnalysisInfo, error) {
	resp, err := s.storage.Patient().PatientAnalysisGet(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.AnalysisInfo{}, status.Error(codes.NotFound, "something went wrong, please not found this analysis")
		}
		return &patient.AnalysisInfo{}, status.Error(codes.Internal, "something went wrong, please check analisys info")
	}

	return resp, nil
}

// Doctor reports...
func (s *PatientService) DoctorReportCreate(ctx context.Context, req *patient.DoctorReportInfo) (*patient.DoctorReportInfo, error) {
	resp, err := s.storage.Patient().DoctorReportCreate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.DoctorReportInfo{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor report")
		}
		return &patient.DoctorReportInfo{}, status.Error(codes.Internal, "something went wrong, please check doctor report info")
	}

	return resp, nil
}

func (s *PatientService) DoctorReportGet(ctx context.Context, req *patient.PatientId) (*patient.DoctorReportInfo, error) {
	resp, err := s.storage.Patient().DoctorReportGet(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.DoctorReportInfo{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor report")
		}
		return &patient.DoctorReportInfo{}, status.Error(codes.Internal, "something went wrong, please check doctor report info")
	}

	return resp, nil
} //

func (s *PatientService) CreatePatientQueue(ctx context.Context, req *patient.CreatePatientQueueReq) (*patient.PatientQueueResp, error) {
	queue, err := s.storage.Patient().CheckServiceQueue(&patient.CheckQueueReq{
		ServiceId:   req.ServiceId,
		ServiceType: req.ServiceType,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println(err.Error())
		return &patient.PatientQueueResp{}, status.Error(codes.Internal, "something went wrong, please check doctor report info")
	}
	if queue == nil {
		queue.QueueNumber = 0
	}
	resp, err := s.storage.Patient().CreatePatientQueue(&patient.CreatePatientQueueReq{
		Id:          req.Id,
		ClientId:    req.ClientId,
		QueueNumber: queue.QueueNumber + 1,
		ServiceId:   req.ServiceId,
		ServiceType: req.ServiceType,
	})

	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.PatientQueueResp{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor report")
		}
		return &patient.PatientQueueResp{}, status.Error(codes.Internal, "something went wrong, please check doctor report info")
	}

	return resp, nil
}

func (s *PatientService) GetPatientQueue(ctx context.Context, req *patient.PaymentHistoryId) (*patient.PatientQueueResp, error) {
	resp, err := s.storage.Patient().GetPatientQueue(req)
	if err != nil {
		return &patient.PatientQueueResp{}, err
	}
	return resp, nil
}

func (s *PatientService) CheckServiceQueue(ctx context.Context, req *patient.CheckQueueReq) (*patient.QueueNumber, error) {
	resp, err := s.storage.Patient().CheckServiceQueue(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.QueueNumber{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor report")
		}
		return &patient.QueueNumber{}, status.Error(codes.Internal, "something went wrong, please check doctor report info")
	}

	return resp, nil
}

func (s *PatientService) UpdateQueue(ctx context.Context, req *patient.UpdateQueueReq) (*patient.PatientQueueResp, error) {
	resp, err := s.storage.Patient().UpdateQueue(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.PatientQueueResp{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor report")
		}
		return &patient.PatientQueueResp{}, status.Error(codes.Internal, "something went wrong, please check doctor report info")
	}

	return resp, nil
}

func (s *PatientService) FindQueue(ctx context.Context, req *patient.QueueFilter) (*patient.QueuesResp, error) {
	resp, err := s.storage.Patient().FindQueue(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.QueuesResp{}, status.Error(codes.NotFound, "something went wrong, please not found this find queue")
		}
		return &patient.QueuesResp{}, status.Error(codes.Internal, "something went wrong, please check find queue info")
	}

	return resp, nil
}

func (s *PatientService) CreateCashbox(ctx context.Context, req *patient.CreateCashboxReq) (*patient.CashboxResp, error) {
	Summa := 0
	for _, doctorId := range req.DoctorsIds {
		doktor, err := s.service.DoctorService().DoctorGet(context.Background(), &doctor.GetDoctorReq{Field: "id", Value: doctorId})
		if err != nil {
			log.Println(err.Error())
			if errors.Is(err, sql.ErrNoRows) {
				return &patient.CashboxResp{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor report")
			}
			return &patient.CashboxResp{}, status.Error(codes.Internal, "something went wrong, please check doctor report info")
		}
		Summa += int(doktor.Price)
	}

	for _, labId := range req.LabsIds {
		labaratoriya, err := s.service.LabService().LabGet(context.Background(), &lab.LabGetReq{Field: "id", Value: labId})
		if err != nil {
			log.Println(err.Error())
			if errors.Is(err, sql.ErrNoRows) {
				return &patient.CashboxResp{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor report")
			}
			return &patient.CashboxResp{}, status.Error(codes.Internal, "something went wrong, please check doctor report info")
		}
		Summa += int(labaratoriya.Price)
	}

	for _, aparatId := range req.AparatsIds {
		aparat, err := s.service.LabService().AparatGet(context.Background(), &lab.AparatGetReq{Field: "id", Value: aparatId})
		if err != nil {
			log.Println(err.Error())
			if errors.Is(err, sql.ErrNoRows) {
				return &patient.CashboxResp{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor report")
			}
			return &patient.CashboxResp{}, status.Error(codes.Internal, "something went wrong, please check doctor report info")
		}
		Summa += int(aparat.Price)
	}

	resp, err := s.storage.Patient().CreateCashbox(&patient.CashboxResp{
		Id:          req.Id,
		ClientId:    req.ClientId,
		Summa:       int64(Summa),
		IsPayed:     req.IsPayed,
		CashCount:   req.CashCount,
		PaymentType: req.PaymentType,
		DoctorsIds:  req.DoctorsIds,
		LabsIds:     req.LabsIds,
		AparatsIds:  req.AparatsIds,
	})
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.CashboxResp{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor report")
		}
		return &patient.CashboxResp{}, status.Error(codes.Internal, "something went wrong, please check doctor report info")
	}
	return resp, nil
}

func (s *PatientService) FindCashbox(ctx context.Context, req *patient.FindCashboxReq) (*patient.FindCashboxResp, error) {
	var (
		ClientIDs []int
	)
	if req.Search != "" {
		clientResp, err := s.storage.Patient().PatientsFind(&patient.PatientsFindReq{
			Limit:  9999999,
			Page:   1,
			Search: req.Search,
		})
		if err != nil {
			log.Println(err.Error())
			if errors.Is(err, sql.ErrNoRows) {
				return &patient.FindCashboxResp{}, status.Error(codes.NotFound, "something went wrong, please not found this find queue")
			}
			return &patient.FindCashboxResp{}, status.Error(codes.Internal, "something went wrong, please check find queue info")
		}

		for _, client := range clientResp.Patients {
			ClientIDs = append(ClientIDs, int(client.ClientId))
		}
	}

	resp, err := s.storage.Patient().FindCashbox(req, ClientIDs)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.FindCashboxResp{}, status.Error(codes.NotFound, "something went wrong, please not found this find queue")
		}
		return &patient.FindCashboxResp{}, status.Error(codes.Internal, "something went wrong, please check find queue info")
	}
	return resp, nil
}

func (s *PatientService) GetCashbox(ctx context.Context, req *patient.GetCashboxReq) (*patient.CashboxResp, error) {
	resp, err := s.storage.Patient().GetCashbox(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.CashboxResp{}, status.Error(codes.NotFound, "something went wrong, please not found this find queue")
		}
		return &patient.CashboxResp{}, status.Error(codes.Internal, "something went wrong, please check find queue info")
	}
	return resp, nil
}

func (s *PatientService) UpdateCashbox(ctx context.Context, req *patient.UpdateCashboxReq) (*patient.CashboxResp, error) {
	resp, err := s.storage.Patient().UpdateCashbox(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.CashboxResp{}, status.Error(codes.NotFound, "something went wrong, please not found this find queue")
		}
		return &patient.CashboxResp{}, status.Error(codes.Internal, "something went wrong, please check find queue info")
	}
	return resp, nil
}

func (s *PatientService) DeleteCashbox(ctx context.Context, req *patient.GetCashboxReq) (*empty.Empty, error) {
	err := s.storage.Patient().DeleteCashbox(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "something went wrong, please not found this find queue")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "something went wrong, please check find queue info")
	}
	return &emptypb.Empty{}, nil
}

func (s *PatientService) CreatePaymentHistory(ctx context.Context, req *patient.CreatePaymentHistoryReq) (*patient.PaymentHistoryResp, error) {
	resp, err := s.storage.Patient().CreatePaymentHistory(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.PaymentHistoryResp{}, status.Error(codes.NotFound, "something went wrong, please not found this find queue")
		}
		return &patient.PaymentHistoryResp{}, status.Error(codes.Internal, "something went wrong, please check find queue info")
	}
	return resp, nil
}

func (s *PatientService) GetPaymentHistory(ctx context.Context, req *patient.PaymentHistoryId) (*patient.PaymentHistoryResp, error) {
	resp, err := s.storage.Patient().GetPaymentHistory(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.PaymentHistoryResp{}, status.Error(codes.NotFound, "something went wrong, please not found this find queue")
		}
		return &patient.PaymentHistoryResp{}, status.Error(codes.Internal, "something went wrong, please check find queue info")
	}
	return resp, nil
}

func (s *PatientService) FindPaymentHistory(ctx context.Context, req *patient.PaymentHistoryFilter) (*patient.PaymentHistoriesResp, error) {
	resp, err := s.storage.Patient().FindPaymentHistory(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &patient.PaymentHistoriesResp{}, status.Error(codes.NotFound, "something went wrong, please not found this find queue")
		}
		return &patient.PaymentHistoriesResp{}, status.Error(codes.Internal, "something went wrong, please check find queue info")
	}
	return resp, nil
}

func (s *PatientService) DeletePaymentHistory(ctx context.Context, req *patient.PaymentHistoryId) (*empty.Empty, error) {
	err := s.storage.Patient().DeletePaymentHistory(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "something went wrong, please not found this find queue")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "something went wrong, please check find queue info")
	}
	return &emptypb.Empty{}, nil
}
