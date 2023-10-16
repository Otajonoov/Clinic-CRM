package models

// Model
type LabModel struct {
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Type          string  `json:"type"`
	SubCategoryId string  `json:"sub_category_id"`
}

type LabModelResp struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Type          string  `json:"type"`
	SubCategoryId string  `json:"sub_category_id"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type LabGetReq struct {
	Field string `json:"field" binding:"required" default:"id"`
	Value string `json:"value"`
}

type LabsFindReq struct {
	Limit  int64  `json:"limit" binding:"required" default:"10"`
	Page   int64  `json:"page" binding:"required" default:"1"`
	Search string `json:"search"`
}

type LabsResp struct {
	Labs  []*LabModelResp `json:"labs"`
	Count int64           `json:"count"`
}

// Aparat
type AparatModel struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Type          string  `json:"type"`
	SubCategoryId string  `json:"sub_category_id"`
}

type CreateAparat struct {
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Type          string  `json:"type"`
	SubCategoryId string  `json:"sub_category_id"`
}

type UpdateAparat struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Type  string  `json:"type"`
}

type AparatModelResp struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Type          string  `json:"type"`
	SubCategoryId string  `json:"sub_category_id"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type AparatGetReq struct {
	Field string `json:"field" binding:"required" default:"id"`
	Value string `json:"value"`
}

type AparatFindReq struct {
	Limit  int64  `json:"limit" binding:"required" default:"10"`
	Page   int64  `json:"page" binding:"required" default:"1"`
	Search string `json:"search"`
}

type AparatsResp struct {
	Aparats []*AparatModelResp `json:"aparats"`
	Count   int64              `json:"count"`
}

// Lab categoty
type CategoryModel struct {
	Name string `json:"name"`
}

type CategoryModelResp struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CategoryGetReqModel struct {
	Field string `json:"field" binding:"required" default:"id"`
	Value string `json:"value"`
}

type CategoryFindReqModel struct {
	Limit      int64  `json:"limit" binding:"required" default:"10"`
	Page       int64  `json:"page" binding:"required" default:"1"`
	CategoryId string `json:"category_id"`
}

type CategoriesResp struct {
	Category []*CategoryModelResp `json:"category"`
	Count    int64                `json:"count"`
}

// Lab sub category
type SubCategoryModel struct {
	Name       string `json:"name"`
	CategoryId string `json:"category_id"`
}

type UpdateSubCategory struct {
	Name string `json:"name"`
}

type SubCategoryResModel struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	CategoryId string `json:"category_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type SubCategoriesResp struct {
	Category []*SubCategoryResModel `json:"category"`
	Count    int64                  `json:"count"`
}

type AnalysisReqModel struct {
	ClientId    int64  `json:"client_id"`
	AparatId    string `json:"aparat_id"`
	AnalysisUrl string `json:"analysis_url"`
}

type AnalysisRespModel struct {
	Id          string `json:"id"`
	ClientId    int64  `json:"client_id"`
	AparatId    string `json:"aparat_id"`
	AnalysisUrl string `json:"analysis_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type AnalysisGetReqModel struct {
	Field string `json:"field" binding:"required" default:"id"`
	Value string `json:"value"`
}

type AparatId struct {
	AparatId string `json:"aparat_id"`
}
