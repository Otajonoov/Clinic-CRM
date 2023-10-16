package service

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"gitlab.com/clinic-crm/doctor/genproto/doctor"
	"gitlab.com/clinic-crm/doctor/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DoctorService struct {
	doctor.UnimplementedDoctorServiceServer
	storage storage.StorageI
}

func NewDoctorService(strg storage.StorageI) *DoctorService {
	return &DoctorService{
		storage:                          strg,
		UnimplementedDoctorServiceServer: doctor.UnimplementedDoctorServiceServer{},
	}
}

// Doctor...
func (s *DoctorService) DoctorCreate(ctx context.Context, req *doctor.Doctor) (*doctor.Doctor, error) {
	resp, err := s.storage.Doctor().DoctorCreate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &doctor.Doctor{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor")
		}
		return &doctor.Doctor{}, status.Error(codes.Internal, "something went wrong, please not found this doctor")
	}

	return resp, nil
}

func (s *DoctorService) DoctorGet(ctx context.Context, req *doctor.GetDoctorReq) (*doctor.Doctor, error) {
	resp, err := s.storage.Doctor().DoctorGet(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &doctor.Doctor{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor")
		}
		return &doctor.Doctor{}, status.Error(codes.Internal, "something went wrong, please not found this doctor")
	}

	return resp, nil
}

func (s *DoctorService) DoctorsFind(ctx context.Context, req *doctor.DoctorsFindReq) (*doctor.DoctorsResp, error) {
	resp, err := s.storage.Doctor().DoctorsFind(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &doctor.DoctorsResp{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor")
		}
		return &doctor.DoctorsResp{}, status.Error(codes.Internal, "something went wrong, please not found this doctor")
	}

	return resp, nil
}

func (s *DoctorService) DoctorUpdate(ctx context.Context, req *doctor.Doctor) (*doctor.Doctor, error) {
	resp, err := s.storage.Doctor().DoctorUpdate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &doctor.Doctor{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor")
		}
		return &doctor.Doctor{}, status.Error(codes.Internal, "something went wrong, please not found this doctor")
	}

	return resp, nil
}

func (s *DoctorService) DoctorDelete(ctx context.Context, req *doctor.DoctorId) (*empty.Empty, error) {
	err := s.storage.Doctor().DoctorDelete(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "something went wrong, please check doctor info")
	}

	return &emptypb.Empty{}, nil
}

// Doctor reports...
func (s *DoctorService) DoctorReportCreate(ctx context.Context, req *doctor.DoctorReport) (*doctor.DoctorReportRes, error) {
	resp, err := s.storage.Doctor().DoctorReportCreate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &doctor.DoctorReportRes{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor report")
		}
		return &doctor.DoctorReportRes{}, status.Error(codes.Internal, "something went wrong, please not found this doctor report")
	}

	return resp, nil
}

func (s *DoctorService) DoctorReportGet(ctx context.Context, req *doctor.GetDoctorReport) (*doctor.DoctorReportRes, error) {
	resp, err := s.storage.Doctor().DoctorReportGet(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &doctor.DoctorReportRes{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor report")
		}
		return &doctor.DoctorReportRes{}, status.Error(codes.Internal, "something went wrong, please not found this doctor report")
	}

	return resp, nil
}

func (s *DoctorService) DoctorReportsFind(ctx context.Context, req *doctor.DoctorReportsFindReq) (*doctor.DoctorReportsResp, error) {
	resp, err := s.storage.Doctor().DoctorReportsFind(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &doctor.DoctorReportsResp{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor report")
		}
		return &doctor.DoctorReportsResp{}, status.Error(codes.Internal, "something went wrong, please not found this doctor report")
	}

	return resp, nil
}

func (s *DoctorService) DoctorReportDelete(ctx context.Context, req *doctor.ReportId) (*empty.Empty, error) {
	err := s.storage.Doctor().DoctorReportDelete(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "something went wrong, please not found this doctor")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "something went wrong, please check doctor info")
	}

	return &emptypb.Empty{}, nil
}

// Sqlad
func (s *DoctorService) SqladCreate(ctx context.Context, req *doctor.SqladReq) (*doctor.SqladRes, error) {
	resp, err := s.storage.Doctor().SqladCreate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &doctor.SqladRes{}, status.Error(codes.NotFound, "something went wrong, please not found this infi")
		}
		return &doctor.SqladRes{}, status.Error(codes.Internal, "something went wrong, please not found this info")
	}

	return resp, nil
}

func (s *DoctorService) SqladGet(ctx context.Context, req *doctor.SqladGetReq) (*doctor.SqladRes, error) {
	resp, err := s.storage.Doctor().SqladGet(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &doctor.SqladRes{}, status.Error(codes.NotFound, "something went wrong, please not found this info")
		}
		return &doctor.SqladRes{}, status.Error(codes.Internal, "something went wrong, please not found this info")
	}

	return resp, nil
}

func (s *DoctorService) SqladUpdate(ctx context.Context, req *doctor.SqladReq) (*doctor.SqladRes, error) {
	resp, err := s.storage.Doctor().SqladUpdate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &doctor.SqladRes{}, status.Error(codes.NotFound, "something went wrong, please not found this info")
		}
		return &doctor.SqladRes{}, status.Error(codes.Internal, "something went wrong, please not found this info")
	}

	return resp, nil
}

func (s *DoctorService) SqladDelete(ctx context.Context, req *doctor.SqladId) (*empty.Empty, error) {
	err := s.storage.Doctor().SqladDelete(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "something went wrong, please not found this info")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "something went wrong, please check this info")
	}

	return &emptypb.Empty{}, nil
}

func (s *DoctorService) LowStock(context.Context, *emptypb.Empty) (*doctor.LowStockRes, error) {
	resp, err := s.storage.Doctor().LowStock()
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &doctor.LowStockRes{}, status.Error(codes.NotFound, "something went wrong, please not found this info")
		}
		return &doctor.LowStockRes{}, status.Error(codes.Internal, "something went wrong, please not found this info")
	}

	return resp, nil
}

func (s *DoctorService) DoctorPageFilter(ctx context.Context, req *doctor.DocPageFilter) (*doctor.DocPageFilterRes, error) {
	resp, err := s.storage.Doctor().DoctorPageFilter(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &doctor.DocPageFilterRes{}, status.Error(codes.NotFound, "something went wrong, please not found this info")
		}
		return &doctor.DocPageFilterRes{}, status.Error(codes.Internal, "something went wrong, please not found this info")
	}

	return resp, nil
}

func (s *DoctorService) DoctorTypeGet(context.Context, *emptypb.Empty) (*doctor.DoctorTypes, error) {
	resp, err := s.storage.Doctor().DoctorTypeGet()
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &doctor.DoctorTypes{}, status.Error(codes.NotFound, "something went wrong, please not found this info")
		}
		return &doctor.DoctorTypes{}, status.Error(codes.Internal, "something went wrong, please not found this info")
	}

	return resp, nil
}