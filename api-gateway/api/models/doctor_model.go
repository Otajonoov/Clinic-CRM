package models

type CreateDoctorModel struct {
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Gender      string  `json:"gender"`
	WorkTime    string  `json:"work_time"`
	Price       float64 `json:"price"`
	Specialty   string  `json:"specialty"`
	RoomNumber  string  `json:"room_number"`
	PhoneNumber string  `json:"phone_number"`
}

type DoctorResp struct {
	Id          string  `json:"id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Gender      string  `json:"gender"`
	WorkTime    string  `json:"work_time"`
	Price       float64 `json:"price"`
	Specialty   string  `json:"specialty"`
	RoomNumber  string  `json:"room_number"`
	PhoneNumber string  `json:"phone_number"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type DoctorsResp struct {
	Doctors []*DoctorResp `json:"doctors"`
	Count   int64         `json:"count"`
}

type GetDoctorReq struct {
	Field string `json:"field" binding:"required" default:"id"`
	Value string `json:"value"`
}

type DoctorsFindReq struct {
	Limit  int64  `json:"limit" binding:"required" default:"10"`
	Page   int64  `json:"page" binding:"required" default:"1"`
	Search string `json:"search"`
}

type DoctorId struct {
	DoctorId string `json:"doctor_id"`
}

type DoctorReportsModel struct {
	ClientId string `json:"client_id"`
	DoctorId string `json:"doctor_id"`
	Text     string `json:"text"`
}

type DoctorReportsModelRes struct {
	Id        string `json:"id"`
	ClientId  string `json:"client_id"`
	DoctorId  string `json:"doctor_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetDoctorReportReq struct {
	Field string `json:"field" binding:"required" default:"id"`
	Value string `json:"value"`
}

type DoctorReportsFindReq struct {
	Limit  int64  `json:"limit" binding:"required" default:"10"`
	Page   int64  `json:"page" binding:"required" default:"1"`
	Search string `json:"search"`
}

type DoctorReportsResp struct {
	Reports []*DoctorReportsModelRes `json:"reports"`
	Count   int64                    `json:"count"`
}

type SqladReqModel struct {
	Name           string  `json:"name"`
	Count          int64   `json:"count"`
	Price          float64 `json:"price"`
	LowStock       int64   `json:"low_stock"`
	ExpirationDate string  `json:"expiration_date"`
	Provider       string  `json:"provider"`
}

type SqladRespModel struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	Count          int64   `json:"count"`
	Price          float64 `json:"price"`
	LowStock       int64   `json:"low_stock"`
	ExpirationDate string  `json:"expiration_date"`
	Provider       string  `json:"provider"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

type SqladGetReqModel struct {
	Field string `json:"field" binding:"required" default:"id"`
	Value string `json:"value"`
}

type SqladId struct {
	SqladId string `json:"sqlad_id"`
}

type LowStocksRespModel struct {
	LowStocks []*SqladRespModel `json:"low_stocks"`
	Count     int64             `json:"count"`
}

type DocPageFilterResModel struct {
	PatientInfo []*PatientInfo
}

type PatientInfo struct {
	QueueNumber    int64  `json:"queue_number"`
	FullName       string `json:"full_name"`
	PhoneNumber    string `json:"phone_number"`
	DateLastVisit  string `json:"date_last_visit"`
	PatientReports []*DoctorReportsModelRes
}

type DocPageFilterReq struct {
	ServiceId   string `json:"service_id"`
	ServiceType string `json:"service_type"`
	ClientId    int64  `json:"client_id"`
	Page        int64  `json:"page"`
	Limit       int64  `json:"limit"`
}

type DoctorTypes struct {
	DoctorTypes []*DoctorType `json:"doctor_types"`
	Count       int64         `json:"count"`
}

type DoctorType struct {
	DoctorType string `json:"doctor_type"`
}
