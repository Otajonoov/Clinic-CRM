package storage

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/clinic-crm/labs/storage/postgres"
	"gitlab.com/clinic-crm/labs/storage/repo"
)

type StorageI interface {
	Lab() repo.LabStorageI
}

type storagePg struct {
	labRepo repo.LabStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		labRepo: postgres.NewLab(db),
	}
}

func (s *storagePg) Lab() repo.LabStorageI {
	return s.labRepo
}
