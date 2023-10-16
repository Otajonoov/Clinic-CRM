package repo

import (
	"gitlab.com/clinic-crm/reception/genproto/patient"
)

type PatientStorageI interface {
	// Patient
	PatientCreate(*patient.Patient) (*patient.Patient, error)
	PatientGet(*patient.GetPatientReq) (*patient.Patient, error)
	PatientUpdate(*patient.Patient) (*patient.Patient, error)
	PatientDelete(*patient.PatientId) error
	PatientsFind(*patient.PatientsFindReq) (*patient.PatientsResp, error)
	PatientsGetInfo(*patient.PatientsGetInfoFilter) (*patient.PatientsResp, error)

	// Patient analysiss
	PatientAnalysisCreate(*patient.AnalysisInfo) (*patient.AnalysisInfo, error)
	PatientAnalysisGet(*patient.PatientPhoneNumber) (*patient.AnalysisInfo, error)

	// Doctor reports...
	DoctorReportCreate(*patient.DoctorReportInfo) (*patient.DoctorReportInfo, error)
	DoctorReportGet(*patient.PatientId) (*patient.DoctorReportInfo, error)

	CreatePatientQueue(*patient.CreatePatientQueueReq) (*patient.PatientQueueResp, error)
	CheckServiceQueue(*patient.CheckQueueReq) (*patient.QueueNumber, error)
	UpdateQueue(*patient.UpdateQueueReq) (*patient.PatientQueueResp, error)
	FindQueue(*patient.QueueFilter) (*patient.QueuesResp, error)
	GetPatientQueue(*patient.PaymentHistoryId) (*patient.PatientQueueResp, error)

	// CashBox
	CreateCashbox(*patient.CashboxResp) (*patient.CashboxResp, error)

	FindCashbox(*patient.FindCashboxReq, []int) (*patient.FindCashboxResp, error)
	GetCashbox(*patient.GetCashboxReq) (*patient.CashboxResp, error)
	UpdateCashbox(*patient.UpdateCashboxReq) (*patient.CashboxResp, error)
	DeleteCashbox(*patient.GetCashboxReq) error

	// Payment History
	CreatePaymentHistory(*patient.CreatePaymentHistoryReq) (*patient.PaymentHistoryResp, error)
	GetPaymentHistory(req *patient.PaymentHistoryId) (*patient.PaymentHistoryResp, error)
	FindPaymentHistory(req *patient.PaymentHistoryFilter) (*patient.PaymentHistoriesResp, error)
	DeletePaymentHistory(req *patient.PaymentHistoryId) error
}
