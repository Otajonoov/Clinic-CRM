package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"gitlab.com/clinic-crm/doctor/genproto/doctor"
	"gitlab.com/clinic-crm/doctor/storage/repo"
)

type doctorRepo struct {
	db *sqlx.DB
}

func NewDoctor(db *sqlx.DB) repo.DoctorStorageI {
	return &doctorRepo{
		db: db,
	}
}

func (dr *doctorRepo) DoctorCreate(req *doctor.Doctor) (*doctor.Doctor, error) {
	var result doctor.Doctor
	query := `
		INSERT INTO doctors(
			id,
			first_name,
			last_name,
			gender,
			work_time,
			price,
			cpecialety,
			room_number,
			phone_number
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING 
			id,
			COALESCE(first_name,'') as first_name,
			COALESCE(last_name,'') as last_name,
			gender,
			COALESCE(work_time,'') as work_time,
			price,
			COALESCE(cpecialety,'') as cpecialety,
			room_number,
			COALESCE(phone_number,'') as phone_number,
			created_at,
			updated_at
		`

	if err := dr.db.DB.QueryRow(query,
		req.Id,
		req.FirstName,
		req.LastName,
		req.Gender,
		req.WorkTime,
		req.Price,
		req.Cpecialety,
		req.RoomNumber,
		req.PhoneNumber,
	).Scan(
		&result.Id,
		&result.FirstName,
		&result.LastName,
		&result.Gender,
		&result.WorkTime,
		&result.Price,
		&result.Cpecialety,
		&result.RoomNumber,
		&result.PhoneNumber,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &doctor.Doctor{}, err
	}

	return &result, nil
}

func (dr *doctorRepo) DoctorGet(req *doctor.GetDoctorReq) (*doctor.Doctor, error) {
	var result doctor.Doctor

	_, err := uuid.Parse(req.Value)

	if (err == nil && req.Field != "id") || (err != nil && req.Field == "id") {
		return &doctor.Doctor{}, sql.ErrNoRows
	}

	query := `
		SELECT
			id,
			COALESCE(first_name,'') as first_name,
			COALESCE(last_name,'') as last_name,
			gender,
			COALESCE(work_time,'') as work_time,
			price,
			COALESCE(cpecialety,'') as cpecialety,
			room_number,
			COALESCE(phone_number,'') as phone_number,
			created_at,
			updated_at
		FROM doctors
		WHERE ` + req.Field + `= '` + req.Value + `' and deleted_at IS NULL
	`

	err = dr.db.DB.QueryRow(query).Scan(
		&result.Id,
		&result.FirstName,
		&result.LastName,
		&result.Gender,
		&result.WorkTime,
		&result.Price,
		&result.Cpecialety,
		&result.RoomNumber,
		&result.PhoneNumber,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &doctor.Doctor{}, err
	} else if err != nil {
		log.Fatalln(err.Error())
		return &doctor.Doctor{}, err
	}

	return &result, nil
}

func (dr *doctorRepo) DoctorsFind(req *doctor.DoctorsFindReq) (*doctor.DoctorsResp, error) {
	result := doctor.DoctorsResp{
		Doctors: make([]*doctor.Doctor, 0),
	}

	offset := (req.Page - 1) * req.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)

	filter := "WHERE deleted_at IS NULL"
	if req.Search != "" {
		str := "%" + req.Search + "%"
		filter += fmt.Sprintf(`
			AND (cpecialety ILIKE '%s')
				`,
			str,
		)
	}

	query := `
	SELECT 
		id,
		COALESCE(first_name,'') as first_name,
		COALESCE(last_name,'') as last_name,
		gender,
		COALESCE(work_time,'') as work_time,
		price,
		COALESCE(cpecialety,'') as cpecialety,
		room_number,
		COALESCE(phone_number,'') as phone_number,
		created_at,
		updated_at
	FROM doctors 
	` + filter + `
	ORDER BY created_at desc
	` + limit
	rows, err := dr.db.Query(query)
	if err != nil {
		return &doctor.DoctorsResp{}, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := doctor.Doctor{}
		err = rows.Scan(
			&temp.Id,
			&temp.FirstName,
			&temp.LastName,
			&temp.Gender,
			&temp.WorkTime,
			&temp.Price,
			&temp.Cpecialety,
			&temp.RoomNumber,
			&temp.PhoneNumber,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return &doctor.DoctorsResp{}, err
		}
		result.Doctors = append(result.Doctors, &temp)
	}
	queryCount := `SELECT count(1) FROM doctors ` + filter
	err = dr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return &doctor.DoctorsResp{}, err
	}

	return &result, nil
}

func (dr *doctorRepo) DoctorUpdate(req *doctor.Doctor) (*doctor.Doctor, error) {
	var result doctor.Doctor
	query := `
		UPDATE
			doctors
		SET
			first_name=$1,
			last_name=$2,
			gender=$3,
			work_time=$4,
			price=$5,
			cpecialety=$6,
			room_number=$7,
			phone_number=$8,
			updated_at= ()
		WHERE
			id=$9 AND deleted_at IS NULL
		RETURNING
			id,
			COALESCE(first_name,'') as first_name,
			COALESCE(last_name,'') as last_name,
			gender,
			COALESCE(work_time,'') as work_time,
			price,
			COALESCE(cpecialety,'') as cpecialety,
			room_number,
			COALESCE(phone_number,'') as phone_number,
			created_at,
			updated_at
	`
	err := dr.db.DB.QueryRow(query,
		req.FirstName,
		req.LastName,
		req.Gender,
		req.WorkTime,
		req.Price,
		req.Cpecialety,
		req.RoomNumber,
		req.PhoneNumber,
		req.Id,
	).Scan(
		&result.Id,
		&result.FirstName,
		&result.LastName,
		&result.Gender,
		&result.WorkTime,
		&result.Price,
		&result.Cpecialety,
		&result.RoomNumber,
		&result.PhoneNumber,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		log.Println(err.Error())
		return &doctor.Doctor{}, err
	}

	return &result, nil
}

func (dr *doctorRepo) DoctorDelete(req *doctor.DoctorId) error {
	query := `
		UPDATE
			doctors
		SET
			deleted_at = NOW()
		WHERE id = $1
	`
	effect, err := dr.db.Exec(query, req.DoctorId)
	if err != nil {
		return err
	}
	rowsCount, err := effect.RowsAffected()
	if err != nil {
		return err
	}
	if rowsCount == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Doctor report
func (dr *doctorRepo) DoctorReportCreate(req *doctor.DoctorReport) (*doctor.DoctorReportRes, error) {
	var result doctor.DoctorReportRes
	query := `
		INSERT INTO doctor_reports(
			id,
			client_id,
			doctor_id,
			text
		) VALUES($1, $2, $3, $4)
		RETURNING 
			id,
			client_id,
			doctor_id,
			COALESCE(text,'') as text,
			created_at,
			updated_at
		`

	if err := dr.db.DB.QueryRow(query,
		req.Id,
		req.ClientId,
		req.DoctorId,
		req.Text,
	).Scan(
		&result.Id,
		&result.ClientId,
		&result.DoctorId,
		&result.Text,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &doctor.DoctorReportRes{}, err
	}

	return &result, nil
}

func (dr *doctorRepo) DoctorReportGet(req *doctor.GetDoctorReport) (*doctor.DoctorReportRes, error) {
	var result doctor.DoctorReportRes
	query := `
		SELECT
			id,
			client_id,
			doctor_id,
			COALESCE(text,'') as text,
			created_at,
			updated_at
		FROM doctor_reports
		WHERE ` + req.Field + `= '` + req.Value + `' and deleted_at IS NULL
	`

	err := dr.db.DB.QueryRow(query).Scan(
		&result.Id,
		&result.ClientId,
		&result.DoctorId,
		&result.Text,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &doctor.DoctorReportRes{}, err
	} else if err != nil {
		log.Fatalln(err.Error())
		return &doctor.DoctorReportRes{}, err
	}

	return &result, nil
}

func (dr *doctorRepo) DoctorReportsFind(req *doctor.DoctorReportsFindReq) (*doctor.DoctorReportsResp, error) {
	result := doctor.DoctorReportsResp{
		DoctorReports: make([]*doctor.DoctorReportRes, 0),
	}

	offset := (req.Page - 1) * req.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)

	filter := "WHERE deleted_at IS NULL"
	if req.Search != "" {
		str := "%" + req.Search + "%"
		filter += fmt.Sprintf(`
			AND (client_id ILIKE '%s' OR doctor_id ILIKE '%s')
			`,
			str, str,
		)
	}

	query := `
	SELECT 
		id,
		client_id,
		doctor_id,
		COALESCE(text,'') as text,
		created_at,
		updated_at
	FROM doctor_reports 
	` + filter + `
	ORDER BY created_at desc
	` + limit
	rows, err := dr.db.Query(query)
	if err != nil {
		return &doctor.DoctorReportsResp{}, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := doctor.DoctorReportRes{}
		err = rows.Scan(
			&temp.Id,
			&temp.ClientId,
			&temp.DoctorId,
			&temp.Text,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return &doctor.DoctorReportsResp{}, err
		}
		result.DoctorReports = append(result.DoctorReports, &temp)
	}
	queryCount := `SELECT count(1) FROM doctor_reports ` + filter
	err = dr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return &doctor.DoctorReportsResp{}, err
	}

	return &result, nil
}

func (dr *doctorRepo) DoctorReportDelete(req *doctor.ReportId) error {
	query := `
		UPDATE
			doctor_reports
		SET
			deleted_at = NOW()
		WHERE id = $1
	`
	effect, err := dr.db.Exec(query, req.ReportId)
	if err != nil {
		return err
	}
	rowsCount, err := effect.RowsAffected()
	if err != nil {
		return err
	}
	if rowsCount == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Sqlad
func (dr *doctorRepo) SqladCreate(req *doctor.SqladReq) (*doctor.SqladRes, error) {
	var result doctor.SqladRes
	query := `
		INSERT INTO sqlad(
			id,
			name,
			count,
			price,
			low_stock,
			expiration_date,
			provider
		) VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING 
			id,
			COALESCE(name,'') as name,
			count,
			price,
			low_stock,
			expiration_date,
			COALESCE(provider,'') as provider,
			created_at,
			updated_at
		`

	if err := dr.db.DB.QueryRow(query,
		req.Id,
		req.Name,
		req.Count,
		req.Price,
		req.LowStock,
		req.ExpirationDate,
		req.Provider,
	).Scan(
		&result.Id,
		&result.Name,
		&result.Count,
		&result.Price,
		&result.LowStock,
		&result.ExpirationDate,
		&result.Provider,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &doctor.SqladRes{}, err
	}

	return &result, nil
}

func (dr *doctorRepo) SqladGet(req *doctor.SqladGetReq) (*doctor.SqladRes, error) {
	var result doctor.SqladRes
	query := `
		SELECT
			id,
			COALESCE(name,'') as name,
			count,
			price,
			low_stock,
			expiration_date,
			COALESCE(provider,'') as provider,
			created_at,
			updated_at
		FROM sqlad
		WHERE ` + req.Field + `= '` + req.Value + `' and deleted_at IS NULL
	`

	err := dr.db.DB.QueryRow(query).Scan(
		&result.Id,
		&result.Name,
		&result.Count,
		&result.Price,
		&result.LowStock,
		&result.ExpirationDate,
		&result.Provider,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &doctor.SqladRes{}, err
	} else if err != nil {
		log.Fatalln(err.Error())
		return &doctor.SqladRes{}, err
	}

	return &result, nil
}

func (dr *doctorRepo) SqladUpdate(req *doctor.SqladReq) (*doctor.SqladRes, error) {
	var result doctor.SqladRes
	query := `
		UPDATE
			sqlad
		SET
			name=$1,
			count=$2,
			price=$3,
			low_stock=$4,
			expiration_date=$5,
			provider=$6,	
			updated_at=NOW()
		WHERE
			id=$7 AND deleted_at IS NULL
		RETURNING
			id,
			COALESCE(name,'') as name,
			count,
			price,
			low_stock,
			expiration_date,
			COALESCE(provider,'') as provider,
			created_at,
			updated_at
	`
	err := dr.db.DB.QueryRow(query,
		req.Name,
		req.Count,
		req.Price,
		req.LowStock,
		req.ExpirationDate,
		req.Provider,
		req.Id,
	).Scan(
		&result.Id,
		&result.Name,
		&result.Count,
		&result.Price,
		&result.LowStock,
		&result.ExpirationDate,
		&result.Provider,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		log.Println(err.Error())
		return &doctor.SqladRes{}, err
	}

	return &result, nil
}

func (dr *doctorRepo) SqladDelete(req *doctor.SqladId) error {
	query := `
		UPDATE
			sqlad
		SET
			deleted_at = NOW()
		WHERE id = $1
	`
	effect, err := dr.db.Exec(query, req.Id)
	if err != nil {
		return err
	}
	rowsCount, err := effect.RowsAffected()
	if err != nil {
		return err
	}
	if rowsCount == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (dr *doctorRepo) LowStock() (*doctor.LowStockRes, error) {
	result := doctor.LowStockRes{
		LowStock: make([]*doctor.SqladRes, 0),
	}

	query := `
		SELECT 
			id,
			COALESCE(name,'') as name,
			count,
			price,
			low_stock,
			expiration_date,
			COALESCE(provider,'') as provider,
			created_at,
			updated_at
	FROM sqlad WHERE count < low_stock AND deleted_at IS NULL
	ORDER BY created_at desc
	`
	rows, err := dr.db.Query(query)
	if err != nil {
		return &doctor.LowStockRes{}, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := doctor.SqladRes{}
		err = rows.Scan(
			&temp.Id,
			&temp.Name,
			&temp.Count,
			&temp.Price,
			&temp.LowStock,
			&temp.ExpirationDate,
			&temp.Provider,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return &doctor.LowStockRes{}, err
		}
		result.LowStock = append(result.LowStock, &temp)
	}
	queryCount := `SELECT count(1) FROM sqlad WHERE count < low_stock AND deleted_at IS NULL`
	err = dr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return &doctor.LowStockRes{}, err
	}

	return &result, nil
}

// Doctor page
func (dr *doctorRepo) DoctorPageFilter(req *doctor.DocPageFilter) (*doctor.DocPageFilterRes, error) {
	var (
		result doctor.DocPageFilterRes
	)
	return &result, nil
}

func (dr *doctorRepo) DoctorTypeGet() (*doctor.DoctorTypes, error) {
	result := doctor.DoctorTypes{
		DoctorTypes: make([]*doctor.DoctorType, 0),
	}

	query := `
	SELECT 
    	cpecialety
	FROM 
		doctors 
	WHERE 
		deleted_at IS NULL
	ORDER BY created_at desc
	`
	rows, err := dr.db.Query(query)
	if err != nil {
		return &doctor.DoctorTypes{}, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := doctor.DoctorType{}
		err = rows.Scan(
			&temp.DoctorType,
		)
		if err != nil {
			return &doctor.DoctorTypes{}, err
		}
		result.DoctorTypes = append(result.DoctorTypes, &temp)
	}

	result.DoctorTypes = removeDuplicates(result.DoctorTypes)

	queryCount := `SELECT count(1) FROM doctors WHERE deleted_at IS NULL`
	err = dr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return &doctor.DoctorTypes{}, err
	}

	return &result, nil
}

func removeDuplicates(slice []*doctor.DoctorType) []*doctor.DoctorType {
	encountered := make(map[string]bool)
	result := []*doctor.DoctorType{}
	for _, item := range slice {
		if encountered[item.DoctorType] {
			continue
		}
		encountered[item.DoctorType] = true
		result = append(result, item)
	}
	return result
}
