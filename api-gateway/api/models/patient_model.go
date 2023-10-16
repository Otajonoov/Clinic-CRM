package models

type PatientModel struct {
	Id                 string `json:"id"`
	ClientId           int64  `json:""`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Patronymic         string `json:"patronomic"`
	DateOfBirth        string `json:"date_of_birth"`
	MainPhoneNumber    string `json:"main_phone_number"`
	OtherPhoneNumber   string `json:"other_phone_number"`
	AdvertisingChannel string `json:"advertising_chanel"`
	Respublic          string `json:"respublic"`
	Region             string `json:"region"`
	District           string `json:"district"`
	PassportInfo       string `json:"passport_info"`
	Discount           string `json:"discount"`
	Condition          string `json:"condition"`
	Gender             string `json:"gender"`
	DoctorId           string `json:"doctor_id"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

type CheckQueueReq struct {
	ServiceId   string `json:"service_id"`
	ServiceType string `json:"service_type"`
}

type QueueNumber struct {
	QueueNumber int64 `json:"queue_number"`
}

type PatientQueueId struct {
	Id string `json:"id"`
}

type CreatePatientModel struct {
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Patronymic         string `json:"patronomic"`
	DateOfBirth        string `json:"date_of_birth"`
	MainPhoneNumber    string `json:"main_phone_number"`
	OtherPhoneNumber   string `json:"other_phone_number"`
	AdvertisingChannel string `json:"advertising_chanel"`
	Respublic          string `json:"respublic"`
	Region             string `json:"region"`
	District           string `json:"district"`
	PassportInfo       string `json:"passport_info"`
	Discount           string `json:"discount"`
	Condition          string `json:"condition"`
	Gender             string `json:"gender"`
	DoctorId           string `json:"doctor_id"`
}

type PatientsResp struct {
	Patients []*PatientModel `json:"patients"`
	Count    int64           `json:"count"`
}

type GetPatientReqModel struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type PatientId struct {
	Id int64 `json:"id"`
}

type PatientsFindModel struct {
	Limit  int64  `json:"limit"`
	Page   int64  `json:"page"`
	Search string `json:"search"`
}

type AnalysisInfo struct {
	Id                string `json:"id"`
	ClientPhoneNumber string `json:"client_phone_number"`
	AnalysName        string `json:"analys_name"`
	AnalysUrl         string `json:"analys_url"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type CreatePatientQueueReq struct {
	ClientId    int64  `json:"client_id"`
	ServiceId   string `json:"service_id"`
	ServiceType string `json:"service_type"`
}

type PatientQueueResp struct {
	Id          string `json:"id"`
	ClientId    int64  `json:"client_id"`
	QueueNumber int64  `json:"queue_number"`
	ServiceId   string `json:"service_id"`
	ServiceType string `json:"service_type"`
	TurnPassed  bool   `json:"turn_passed"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateQueueReq struct {
	ClientId    int64  `json:"client_id"`
	ServiceId   string `json:"service_id"`
	ServiceType string `json:"service_type"`
}

type QueueFilter struct {
	ServiceId   string `json:"service_id"`
	ServiceType string `json:"service_type"`
	ClientId    int64  `json:"client_id"`
	Page        int64  `json:"page"`
	Limit       int64  `json:"limit"`
}

type QueuesResp struct {
	Queues []*PatientQueueResp `json:"queues"`
	Count  int                 `json:"count"`
}

type CreateCashboxReq struct {
	ClientId    int64    `json:"client_id"`
	IsPayed     bool     `json:"is_payed"`
	PaymentType string   `json:"payment_type"`
	DoctorsIds  []string `json:"doctors_ids"`
	LabsIds     []string `json:"labs_ids"`
	AparatsIds  []string `json:"aparats_ids"`
}

type CashboxResp struct {
	Id         string   `json:"id"`
	ClientId   int      `json:"client_id"`
	Summa      int      `json:"summa"`
	IsPayed    bool     `json:"is_payed"`
	CashCount  int      `json:"cash_count"`
	DoctorsIds []string `json:"doctors_ids"`
	LabsIds    []string `json:"labs_ids"`
	AparatsIds []string `json:"aparats_ids"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
}

type CashboxesPrinterResp struct {
	Cashboxes []*CashboxPrinterResp `json:"cashboxes"`
	Count     int                   `json:"count"`
}

type CashboxPrinterResp struct {
	ImageUrl    string `json:"image_url"`
	CashCount   int    `json:"cash_count"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	ServiceType string `json:"service_type"`
	DoctorName  string `json:"doctor_name"`
	RoomNumber  string `json:"room_number"`
	Summa       int    `json:"summa"`
	CreatedAt   string `json:"created_at"`
}

type FindCashboxReq struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Search   string `json:"search"`
	ClientId int    `json:"client_id"`
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
}

type FindCashboxResp struct {
	Cashboxes []*CashboxResp `json:"cashboxes"`
	Count     int            `json:"count"`
}

type UpdateCashboxReq struct {
	Id          string `json:"id"`
	IsPayed     bool   `json:"is_payed"`
	PaymentType string `json:"payment_type"`
}

type CreatePaymentHistoryReq struct {
	ClientId    int64  `json:"client_id"`
	Summa       int64  `json:"summa"`
	PaymentType string `json:"payment_type"`
	CashboxId   string `json:"cashbox_id"`
}

type PaymentHistoryResp struct {
	Id          string `json:"id"`
	ClientId    int64  `json:"client_id"`
	Summa       int64  `json:"summa"`
	PaymentType string `json:"payment_type"`
	CashboxId   string `json:"cashbox_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type PaymentHistoryId struct {
	Id string `json:"id"`
}

type PaymentHistoryFilter struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	ClientId int    `json:"client_id"`
	FromDate string `json:"from_date"`
	ToDate   string `json:"to_date"`
}

type PaymentHistoriesResp struct {
	PaymentHistories []*PaymentHistoryResp `json:"payment_histories"`
	Count            int                   `json:"count"`
}
