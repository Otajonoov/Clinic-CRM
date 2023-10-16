package repo

import (
	"gitlab.com/clinic-crm/labs/genproto/lab"
)

type LabStorageI interface {
	// Labs
	LabCreate(*lab.LabCreateReq) (*lab.LabCreateRes, error)
	LabGet(*lab.LabGetReq) (*lab.LabCreateRes, error)
	LabsFind(*lab.LabsFindReq) (*lab.LabsRes, error)
	LabUpdate(*lab.LabUpdateReq) (*lab.LabCreateRes, error)
	LabDelete(*lab.LabId) error

	// Aparats
	AparatCreate(*lab.AparatCreateReq) (*lab.AparatCreateRes, error)
	AparatGet(*lab.AparatGetReq) (*lab.AparatCreateRes, error) 
	AparatsFind(*lab.AparatsFindReq) (*lab.AparatsRes, error)
	AparatsUpdate(*lab.AparatUpdateReq) (*lab.AparatCreateRes, error)
	AparatsDelete(*lab.AparatId) error

	// Lab category
	LabCategoryCreate(*lab.Category) (*lab.CategoryRes, error)
	LabCategoryGet(*lab.CategoryGetReq) (*lab.CategoryRes, error)
	LabCategoryFind(*lab.CategoryFindReq) (*lab.CategoriesRes, error)
	LabCategoryUpdate(*lab.Category) (*lab.CategoryRes, error)
	LabCategoryDelete(*lab.CategoryId) error

	// Aparat category
	AparatCategoryCreate(*lab.Category) (*lab.CategoryRes, error)
	AparatCategoryGet(*lab.CategoryGetReq) (*lab.CategoryRes, error)
	AparatCategoryFind(*lab.CategoryFindReq) (*lab.CategoriesRes, error)
	AparatCategoryUpdate(*lab.Category) (*lab.CategoryRes, error)
	AparatCategoryDelete(*lab.CategoryId) error

	// Lab sub category
	LabSubCategoryCreate(*lab.SubCategory) (*lab.SubCategoryRes, error)
	LabSubCategoryGet(*lab.CategoryGetReq) (*lab.SubCategoryRes, error) 
	LabSubCategoryFind(*lab.SubCategoryFindReq) (*lab.SubCategoriesRes, error)
	LabSubCategoryUpdate(*lab.SubCategory) (*lab.SubCategoryRes, error)
	LabSubCategoryDelete(*lab.CategoryId) error

	// Aparat sub category
	AparatSubCategoryCreate(*lab.SubCategory) (*lab.SubCategoryRes, error)
	AparatSubCategoryGet(*lab.CategoryGetReq) (*lab.SubCategoryRes, error) 
	AparatSubCategoryFind(*lab.SubCategoryFindReq) (*lab.SubCategoriesRes, error)
	AparatSubCategoryUpdate(*lab.SubCategory) (*lab.SubCategoryRes, error)
	AparatSubCategoryDelete(*lab.CategoryId) error

	// Aparat analysis
	AparatAnalysisCreate(*lab.AnalysisReq) (*lab.AnalysisResp, error)
	AparatAnalysisGet(*lab.AnalysisGetReq) (*lab.AnalysisResp, error)
	AparatAnalysisDelete(*lab.AparatId) error

	// // Lab analysis
	LabAnalysisCreate(*lab.AnalysisReq) (*lab.AnalysisResp, error)
	LabAnalysisGet(*lab.AnalysisGetReq) (*lab.AnalysisResp, error)
	LabAnalysisDelete(*lab.AparatId) error
}