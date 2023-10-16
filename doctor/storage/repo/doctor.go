package repo

import (
	pb "gitlab.com/clinic-crm/doctor/genproto/doctor"
)

type DoctorStorageI interface {
	// Doctor
	DoctorCreate(*pb.Doctor) (*pb.Doctor, error)
	DoctorGet(*pb.GetDoctorReq) (*pb.Doctor, error)
	DoctorsFind(*pb.DoctorsFindReq) (*pb.DoctorsResp, error)
	DoctorUpdate(*pb.Doctor) (*pb.Doctor, error)
	DoctorDelete(*pb.DoctorId) error

	DoctorTypeGet() (*pb.DoctorTypes, error)

	// Doctor reports
	DoctorReportCreate(*pb.DoctorReport) (*pb.DoctorReportRes, error)
	DoctorReportGet(*pb.GetDoctorReport) (*pb.DoctorReportRes, error)
	DoctorReportsFind(*pb.DoctorReportsFindReq) (*pb.DoctorReportsResp, error)
	DoctorReportDelete(*pb.ReportId) error

	// Sqlad
	SqladCreate(*pb.SqladReq) (*pb.SqladRes, error)
	SqladGet(*pb.SqladGetReq) (*pb.SqladRes, error)
	SqladUpdate(*pb.SqladReq) (*pb.SqladRes, error)
	SqladDelete(*pb.SqladId) error
	LowStock() (*pb.LowStockRes, error)

	// Doctor page
	DoctorPageFilter(*pb.DocPageFilter) (*pb.DocPageFilterRes, error)
}

