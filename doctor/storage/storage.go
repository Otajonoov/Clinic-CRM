package storage

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/clinic-crm/doctor/storage/postgres"
	"gitlab.com/clinic-crm/doctor/storage/repo"
)

type StorageI interface {
	Doctor() repo.DoctorStorageI
}

type storagePg struct {
	doctorRepo repo.DoctorStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		doctorRepo: postgres.NewDoctor(db),
	}
}

func (s *storagePg) Doctor() repo.DoctorStorageI {
	return s.doctorRepo
}