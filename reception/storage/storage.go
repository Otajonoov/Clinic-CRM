package storage

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/clinic-crm/reception/storage/postgres"
	"gitlab.com/clinic-crm/reception/storage/repo"
)

type StorageI interface {
	Patient() repo.PatientStorageI
}

type storagePg struct {
	patientRepo repo.PatientStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		patientRepo: postgres.NewPatient(db),
	}
}

func (s *storagePg) Patient() repo.PatientStorageI {
	return s.patientRepo
}
