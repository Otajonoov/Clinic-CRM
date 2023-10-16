package service

import (
	//pb "gitlab.com/clinic-crm/labs/genproto/lab"

	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"gitlab.com/clinic-crm/labs/genproto/lab"
	"gitlab.com/clinic-crm/labs/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LabService struct {
	storage storage.StorageI
	lab.UnimplementedLabServiceServer
}

func NewLabService(strg storage.StorageI) *LabService {
	return &LabService{
		storage:                       strg,
		UnimplementedLabServiceServer: lab.UnimplementedLabServiceServer{},
	}
}

// Labs
func (s *LabService) LabCreate(ctx context.Context, req *lab.LabCreateReq) (*lab.LabCreateRes, error) {
	resp, err := s.storage.Lab().LabCreate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.LabCreateRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.LabCreateRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) LabGet(ctx context.Context, req *lab.LabGetReq) (*lab.LabCreateRes, error) {
	resp, err := s.storage.Lab().LabGet(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.LabCreateRes{}, status.Error(codes.NotFound, "something went wrong, please not found this lab")
		}
		return &lab.LabCreateRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) LabsFind(ctx context.Context, req *lab.LabsFindReq) (*lab.LabsRes, error) {
	resp, err := s.storage.Lab().LabsFind(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.LabsRes{}, status.Error(codes.NotFound, "something went wrong, please not found this labs")
		}
		return &lab.LabsRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) LabUpdate(ctx context.Context, req *lab.LabUpdateReq) (*lab.LabCreateRes, error) {
	resp, err := s.storage.Lab().LabUpdate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.LabCreateRes{}, status.Error(codes.NotFound, "something went wrong, please not found this labs")
		}
		return &lab.LabCreateRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) LabDelete(ctx context.Context, req *lab.LabId) (*empty.Empty, error) {
	err := s.storage.Lab().LabDelete(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "something went wrong, please not found this lab")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "something went wrong, please check lab info")
	}

	return &emptypb.Empty{}, nil
}

// Aparats
func (s *LabService) AparatCreate(ctx context.Context, req *lab.AparatCreateReq) (*lab.AparatCreateRes, error) {
	resp, err := s.storage.Lab().AparatCreate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.AparatCreateRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.AparatCreateRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) AparatGet(ctx context.Context, req *lab.AparatGetReq) (*lab.AparatCreateRes, error) {
	resp, err := s.storage.Lab().AparatGet(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.AparatCreateRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.AparatCreateRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) AparatsFind(ctx context.Context, req *lab.AparatsFindReq) (*lab.AparatsRes, error) {
	resp, err := s.storage.Lab().AparatsFind(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.AparatsRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.AparatsRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) AparatsUpdate(ctx context.Context, req *lab.AparatUpdateReq) (*lab.AparatCreateRes, error) {
	resp, err := s.storage.Lab().AparatsUpdate(req)
	fmt.Println(err)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.AparatCreateRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.AparatCreateRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) AparatsDelete(ctx context.Context, req *lab.AparatId) (*empty.Empty, error) {
	err := s.storage.Lab().AparatsDelete(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "something went wrong, please not found this lab")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "something went wrong, please check lab info")
	}

	return &emptypb.Empty{}, nil
}

// Lab category
func (s *LabService) LabCategoryCreate(ctx context.Context, req *lab.Category) (*lab.CategoryRes, error) {
	resp, err := s.storage.Lab().LabCategoryCreate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.CategoryRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.CategoryRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) LabCategoryGet(ctx context.Context, req *lab.CategoryGetReq) (*lab.CategoryRes, error) {
	resp, err := s.storage.Lab().LabCategoryGet(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.CategoryRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.CategoryRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) LabCategoryFind(ctx context.Context, req *lab.CategoryFindReq) (*lab.CategoriesRes, error) {
	resp, err := s.storage.Lab().LabCategoryFind(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.CategoriesRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.CategoriesRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) LabCategoryUpdate(ctx context.Context, req *lab.Category) (*lab.CategoryRes, error) {
	resp, err := s.storage.Lab().LabCategoryUpdate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.CategoryRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.CategoryRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) LabCategoryDelete(ctx context.Context, req *lab.CategoryId) (*empty.Empty, error) {
	err := s.storage.Lab().LabCategoryDelete(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "something went wrong, please not found this lab")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "something went wrong, please check lab info")
	}

	return &emptypb.Empty{}, nil
}

// Aparat category
func (s *LabService) AparatCategoryCreate(ctx context.Context, req *lab.Category) (*lab.CategoryRes, error) {
	resp, err := s.storage.Lab().AparatCategoryCreate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.CategoryRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.CategoryRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) AparatCategoryGet(ctx context.Context, req *lab.CategoryGetReq) (*lab.CategoryRes, error) {
	resp, err := s.storage.Lab().AparatCategoryGet(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.CategoryRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.CategoryRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) AparatCategoryFind(ctx context.Context, req *lab.CategoryFindReq) (*lab.CategoriesRes, error) {
	resp, err := s.storage.Lab().AparatCategoryFind(req)
	fmt.Println(resp)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.CategoriesRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.CategoriesRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) AparatCategoryUpdate(ctx context.Context, req *lab.Category) (*lab.CategoryRes, error) {
	resp, err := s.storage.Lab().AparatCategoryUpdate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.CategoryRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.CategoryRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) AparatCategoryDelete(ctx context.Context, req *lab.CategoryId) (*empty.Empty, error) {
	err := s.storage.Lab().AparatCategoryDelete(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "something went wrong, please not found this lab")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "something went wrong, please check lab info")
	}

	return &emptypb.Empty{}, nil
}

// Lab sub category
func (s *LabService) LabSubCategoryCreate(ctx context.Context, req *lab.SubCategory) (*lab.SubCategoryRes, error) {
	resp, err := s.storage.Lab().LabSubCategoryCreate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.SubCategoryRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.SubCategoryRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) LabSubCategoryGet(ctx context.Context, req *lab.CategoryGetReq) (*lab.SubCategoryRes, error) {
	resp, err := s.storage.Lab().LabSubCategoryGet(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.SubCategoryRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.SubCategoryRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) LabSubCategoryFind(ctx context.Context, req *lab.SubCategoryFindReq) (*lab.SubCategoriesRes, error) {
	resp, err := s.storage.Lab().LabSubCategoryFind(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.SubCategoriesRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.SubCategoriesRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) LabSubCategoryUpdate(ctx context.Context, req *lab.SubCategory) (*lab.SubCategoryRes, error) {
	resp, err := s.storage.Lab().LabSubCategoryUpdate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.SubCategoryRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.SubCategoryRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) LabSubCategoryDelete(ctx context.Context, req *lab.CategoryId) (*empty.Empty, error) {
	err := s.storage.Lab().LabSubCategoryDelete(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "something went wrong, please not found this lab")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "something went wrong, please check lab info")
	}

	return &emptypb.Empty{}, nil
}

// Aparat sub category
func (s *LabService) AparatSubCategoryCreate(ctx context.Context, req *lab.SubCategory) (*lab.SubCategoryRes, error) {
	resp, err := s.storage.Lab().AparatSubCategoryCreate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.SubCategoryRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.SubCategoryRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) AparatSubCategoryGet(ctx context.Context, req *lab.CategoryGetReq) (*lab.SubCategoryRes, error) {
	resp, err := s.storage.Lab().AparatSubCategoryGet(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.SubCategoryRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.SubCategoryRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) AparatSubCategoryFind(ctx context.Context, req *lab.SubCategoryFindReq) (*lab.SubCategoriesRes, error) {
	resp, err := s.storage.Lab().AparatSubCategoryFind(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.SubCategoriesRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.SubCategoriesRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) AparatSubCategoryUpdate(ctx context.Context, req *lab.SubCategory) (*lab.SubCategoryRes, error) {
	resp, err := s.storage.Lab().AparatSubCategoryUpdate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.SubCategoryRes{}, status.Error(codes.NotFound, "something went wrong, please not found this")
		}
		return &lab.SubCategoryRes{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) AparatSubCategoryDelete(ctx context.Context, req *lab.CategoryId) (*empty.Empty, error) {
	err := s.storage.Lab().AparatSubCategoryDelete(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "something went wrong, please not found this lab")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "something went wrong, please check lab info")
	}

	return &emptypb.Empty{}, nil
}


// Aparat analysis
func (s *LabService) AparatAnalysisCreate(ctx context.Context, req *lab.AnalysisReq) (*lab.AnalysisResp, error) {
	resp, err := s.storage.Lab().AparatAnalysisCreate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.AnalysisResp{}, status.Error(codes.NotFound, "something went wrong, please not found this aparat analysis")
		}
		return &lab.AnalysisResp{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) AparatAnalysisGet(ctx context.Context, req *lab.AnalysisGetReq) (*lab.AnalysisResp, error) {
	resp, err := s.storage.Lab().AparatAnalysisGet(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.AnalysisResp{}, status.Error(codes.NotFound, "something went wrong, please not found this aparat analysis")
		}
		return &lab.AnalysisResp{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) AparatAnalysisDelete(ctx context.Context, req *lab.AparatId) (*empty.Empty, error) {
	err := s.storage.Lab().AparatAnalysisDelete(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "something went wrong, please not found this aparat analysis")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "something went wrong, please check lab info")
	}

	return &emptypb.Empty{}, nil
}


// Lab analysis
func (s *LabService) LabAnalysisCreate(ctx context.Context, req *lab.AnalysisReq) (*lab.AnalysisResp, error) {
	resp, err := s.storage.Lab().LabAnalysisCreate(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.AnalysisResp{}, status.Error(codes.NotFound, "something went wrong, please not found this lab analysis")
		}
		return &lab.AnalysisResp{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) LabAnalysisGet(ctx context.Context, req *lab.AnalysisGetReq) (*lab.AnalysisResp, error) {
	resp, err := s.storage.Lab().LabAnalysisGet(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &lab.AnalysisResp{}, status.Error(codes.NotFound, "something went wrong, please not found this lab analysis")
		}
		return &lab.AnalysisResp{}, status.Error(codes.Internal, "something went wrong, please check info")
	}

	return resp, nil
}

func (s *LabService) LabAnalysisDelete(ctx context.Context, req *lab.AparatId) (*empty.Empty, error) {
	err := s.storage.Lab().LabAnalysisDelete(req)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &emptypb.Empty{}, status.Error(codes.NotFound, "something went wrong, please not found this lab analysis")
		}
		return &emptypb.Empty{}, status.Error(codes.Internal, "something went wrong, please check lab info")
	}

	return &emptypb.Empty{}, nil
}
