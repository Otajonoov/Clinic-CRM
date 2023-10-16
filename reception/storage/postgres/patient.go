package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"gitlab.com/clinic-crm/reception/genproto/patient"
	"gitlab.com/clinic-crm/reception/storage/repo"
)

type patientRepo struct {
	db *sqlx.DB
}

func NewPatient(db *sqlx.DB) repo.PatientStorageI {
	return &patientRepo{
		db: db,
	}
}

func (pr *patientRepo) PatientCreate(req *patient.Patient) (*patient.Patient, error) {
	var result patient.Patient
	query := `
		INSERT INTO patients(
			id,
			client_id,
			doctor_id,
			first_name,
			last_name,
			patronymic,
			date_of_birth,
			main_phone_number,
			other_phone_number,
			advertising_channel,
			respublic,
			region,
			district,
			passport_info,
			discount,
			condition,
			gender
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
		RETURNING
			id,
			client_id,
			doctor_id,
			COALESCE(first_name,'') as first_name,
			COALESCE(last_name,'') as last_name,
			COALESCE(patronymic,'') as patronymic,
			date_of_birth,
			COALESCE(main_phone_number,'') as main_phone_number,
			COALESCE(other_phone_number,'') as other_phone_number,
			COALESCE(advertising_channel,'') as advertising_channel,
			respublic,
			COALESCE(region,'') as region,
			COALESCE(district,'') as district,
			COALESCE(passport_info,'') as passport_info,
			COALESCE(discount,'') as discount,
			COALESCE(condition,'') as condition,
			gender,
			created_at,
			updated_at
		`
	if err := pr.db.QueryRow(query,
		req.Id,
		req.ClientId,
		req.DoctorId,
		req.FirstName,
		req.LastName,
		req.Patronymic,
		req.DateOfBirth,
		req.MainPhoneNumber,
		req.OtherPhoneNumber,
		req.AdvertisingChannel,
		req.Respublic,
		req.Region,
		req.District,
		req.PassportInfo,
		req.Discount,
		req.Condition,
		req.Gender,
	).Scan(
		&result.Id,
		&result.ClientId,
		&result.DoctorId,
		&result.FirstName,
		&result.LastName,
		&result.Patronymic,
		&result.DateOfBirth,
		&result.MainPhoneNumber,
		&result.OtherPhoneNumber,
		&result.AdvertisingChannel,
		&result.Respublic,
		&result.Region,
		&result.District,
		&result.PassportInfo,
		&result.Discount,
		&result.Condition,
		&result.Gender,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		log.Println(err.Error())
		return &patient.Patient{}, err

	}
	return &result, nil
}

func (pr *patientRepo) PatientGet(req *patient.GetPatientReq) (*patient.Patient, error) {
	var result patient.Patient

	query := `
		SELECT
			id,
			client_id,
			doctor_id,
			COALESCE(first_name,'') as first_name,
			COALESCE(last_name,'') as last_name,
			COALESCE(patronymic,'') as patronymic,
			date_of_birth,
			COALESCE(main_phone_number,'') as main_phone_number,
			COALESCE(other_phone_number,'') as other_phone_number,
			COALESCE(advertising_channel,'') as advertising_channel,
			respublic,
			COALESCE(region,'') as region,
			COALESCE(district,'') as district,
			COALESCE(passport_info,'') as passport_info,
			COALESCE(discount,'') as discount,
			COALESCE(condition,'') as condition,
			gender,
			created_at,
			updated_at
		FROM patients
		WHERE ` + req.Field + `='` + req.Value + `' and deleted_at IS NULL
	`

	err := pr.db.QueryRow(query).Scan(
		&result.Id,
		&result.ClientId,
		&result.DoctorId,
		&result.FirstName,
		&result.LastName,
		&result.Patronymic,
		&result.DateOfBirth,
		&result.MainPhoneNumber,
		&result.OtherPhoneNumber,
		&result.AdvertisingChannel,
		&result.Respublic,
		&result.Region,
		&result.District,
		&result.PassportInfo,
		&result.Discount,
		&result.Condition,
		&result.Gender,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &patient.Patient{}, err
	} else if err != nil {
		log.Println(err.Error())
		return &patient.Patient{}, err
	}

	return &result, nil
}

func (r *patientRepo) PatientUpdate(req *patient.Patient) (*patient.Patient, error) {
	var result patient.Patient
	query := `
		UPDATE
			patients
		SET
			first_name=$1,
			last_name=$2,
			patronymic=$3,
			date_of_birth=$4,
			main_phone_number=$5,
			other_phone_number=$6,
			advertising_channel=$7,
			respublic=$8,
			region=$9,
			district=$10,
			passport_info=$11,
			discount=$12,
			condition=$13,
			gender=$14,
			updated_at=NOW()
		WHERE
			id=$15 AND deleted_at IS NULL
		RETURNING
			id,
			client_id,
			doctor_id,
			COALESCE(first_name,'') as first_name,
			COALESCE(last_name,'') as last_name,
			COALESCE(patronymic,'') as patronymic,
			date_of_birth,
			COALESCE(main_phone_number,'') as main_phone_number,
			COALESCE(other_phone_number,'') as other_phone_number,
			COALESCE(advertising_channel,'') as advertising_channel,
			respublic,
			COALESCE(region,'') as region,
			COALESCE(district,'') as district,
			COALESCE(passport_info,'') as passport_info,
			COALESCE(discount,'') as discount,
			COALESCE(condition,'') as condition,
			gender,
			created_at,
			updated_at
	`
	err := r.db.QueryRow(query,
		req.FirstName,
		req.LastName,
		req.Patronymic,
		req.DateOfBirth,
		req.MainPhoneNumber,
		req.OtherPhoneNumber,
		req.AdvertisingChannel,
		req.Respublic,
		req.Region,
		req.District,
		req.PassportInfo,
		req.Discount,
		req.Condition,
		req.Gender,
		req.Id,
	).Scan(
		&result.Id,
		&result.ClientId,
		&result.DoctorId,
		&result.FirstName,
		&result.LastName,
		&result.Patronymic,
		&result.DateOfBirth,
		&result.MainPhoneNumber,
		&result.OtherPhoneNumber,
		&result.AdvertisingChannel,
		&result.Respublic,
		&result.Region,
		&result.District,
		&result.PassportInfo,
		&result.Discount,
		&result.Condition,
		&result.Gender,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		log.Println(err.Error())
		return &patient.Patient{}, err
	}

	return &result, nil
}

func (r *patientRepo) PatientDelete(req *patient.PatientId) error {
	query := `
		UPDATE 
			patients
		SET
			deleted_at = NOW()
		WHERE client_id = $1
	`
	effect, err := r.db.Exec(query, req.ClientId)
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

func (pr *patientRepo) PatientsFind(req *patient.PatientsFindReq) (*patient.PatientsResp, error) {
	result := patient.PatientsResp{
		Patients: make([]*patient.Patient, 0),
	}

	offset := (req.Page - 1) * req.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)

	filter := "WHERE deleted_at IS NULL"
	if req.Search != "" {
		str := "%" + req.Search + "%"
		filter += fmt.Sprintf(`
			AND (first_name ILIKE '%s' OR last_name ILIKE '%s' OR patronymic ILIKE '%s' 
				OR main_phone_number ILIKE '%s' OR other_phone_number ILIKE '%s'
				OR advertising_channel ILIKE '%s' OR respublic ILIKE '%s' OR region ILIKE '%s' 
				OR district ILIKE '%s' OR passport_info ILIKE '%s')
				`,
			str, str, str, str, str, str, str, str, str, str,
		)
	}

	query := `
	SELECT 
		id,
		client_id,
		doctor_id,
		COALESCE(first_name,'') as first_name,
		COALESCE(last_name,'') as last_name,
		COALESCE(patronymic,'') as patronymic,
		date_of_birth,
		COALESCE(main_phone_number,'') as main_phone_number,
		COALESCE(other_phone_number,'') as other_phone_number,
		COALESCE(advertising_channel,'') as advertising_channel,
		respublic,
		COALESCE(region,'') as region,
		COALESCE(district,'') as district,
		COALESCE(passport_info,'') as passport_info,
		COALESCE(discount,'') as discount,
		COALESCE(condition,'') as condition,
		gender,
		created_at,
		updated_at
	FROM
		patients 
	` + filter + `
	ORDER BY created_at desc
	` + limit
	rows, err := pr.db.Query(query)
	if err != nil {
		return &patient.PatientsResp{}, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := patient.Patient{}
		err = rows.Scan(
			&temp.Id,
			&temp.ClientId,
			&temp.DoctorId,
			&temp.FirstName,
			&temp.LastName,
			&temp.Patronymic,
			&temp.DateOfBirth,
			&temp.MainPhoneNumber,
			&temp.OtherPhoneNumber,
			&temp.AdvertisingChannel,
			&temp.Respublic,
			&temp.Region,
			&temp.District,
			&temp.PassportInfo,
			&temp.Discount,
			&temp.Condition,
			&temp.Gender,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return &patient.PatientsResp{}, err
		}

		result.Patients = append(result.Patients, &temp)
	}
	queryCount := `SELECT count(1) FROM patients ` + filter
	err = pr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return &patient.PatientsResp{}, err
	}

	return &result, nil
}

func (pr *patientRepo) PatientsGetInfo(req *patient.PatientsGetInfoFilter) (*patient.PatientsResp, error) {
	result := patient.PatientsResp{
		Patients: make([]*patient.Patient, 0),
	}

	offset := (req.Page - 1) * req.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)

	filter := "WHERE deleted_at IS NULL"
	strPhone := "%" + req.PhoneNumber + "%"
	strName := "%" + req.Fullname + "%"

	filter += fmt.Sprintf(`
			AND (client_id ILIKE '%d' OR main_phone_number ILIKE '%s' OR first_name ILIKE '%s')
			`,
		req.ClientId, strPhone, strName,
	)

	query := `
	SELECT 
		id,
		client_id,
		doctor_id,
		COALESCE(first_name,'') as first_name,
		COALESCE(last_name,'') as last_name,
		COALESCE(patronymic,'') as patronymic,
		date_of_birth,
		COALESCE(main_phone_number,'') as main_phone_number,
		COALESCE(other_phone_number,'') as other_phone_number,
		COALESCE(advertising_channel,'') as advertising_channel,
		respublic,
		COALESCE(region,'') as region,
		COALESCE(district,'') as district,
		COALESCE(passport_info,'') as passport_info,
		COALESCE(discount,'') as discount,
		COALESCE(condition,'') as condition,
		gender,
		created_at,
		updated_at
	FROM
		patients 
	` + filter + `
	ORDER BY created_at desc
	` + limit
	rows, err := pr.db.Query(query)
	if err != nil {
		return &patient.PatientsResp{}, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := patient.Patient{}
		err = rows.Scan(
			&temp.Id,
			&temp.ClientId,
			&temp.DoctorId,
			&temp.FirstName,
			&temp.LastName,
			&temp.Patronymic,
			&temp.DateOfBirth,
			&temp.MainPhoneNumber,
			&temp.OtherPhoneNumber,
			&temp.AdvertisingChannel,
			&temp.Respublic,
			&temp.Region,
			&temp.District,
			&temp.PassportInfo,
			&temp.Discount,
			&temp.Condition,
			&temp.Gender,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return &patient.PatientsResp{}, err
		}

		result.Patients = append(result.Patients, &temp)
	}
	queryCount := `SELECT count(1) FROM patients ` + filter
	err = pr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return &patient.PatientsResp{}, err
	}

	return &result, nil
}

// Patient analysis
func (r *patientRepo) PatientAnalysisCreate(req *patient.AnalysisInfo) (*patient.AnalysisInfo, error) {
	var result patient.AnalysisInfo
	query := `
		INSERT INTO analysiss(
			id,
			client_phone_number,
			analysiss_name,
			analysiss_url
		) VALUES($1, $2, $3, $4)
		RETURNING
			id,
			COALESCE(client_phone_number,'') as client_phone_number,
			COALESCE(analysiss_name,'') as analysiss_name,
			COALESCE(analysiss_url,'') as analysiss_url,
			created_at,
			updated_at
		`
	if err := r.db.QueryRow(
		query, req.Id, req.ClientPhoneNumber,
		req.AnalysName, req.AnalysUrl,
	).Scan(
		&result.Id,
		&result.ClientPhoneNumber,
		&result.AnalysName,
		&result.AnalysUrl,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		log.Println(err.Error())
		return &patient.AnalysisInfo{}, err
	}
	return &result, nil
}

func (r *patientRepo) PatientAnalysisGet(req *patient.PatientPhoneNumber) (*patient.AnalysisInfo, error) {
	var result patient.AnalysisInfo

	query := `
		SELECT
			id,
			COALESCE(client_phone_number,'') as client_phone_number,
			COALESCE(analysiss_name,'') as analysiss_name,
			COALESCE(analysiss_url,'') as analysiss_url,
			created_at,
			updated_at
		FROM analysiss
		WHERE client_phone_number=$1 and deleted_at IS NULL
	`

	err := r.db.QueryRow(query, req.PhoneNumber).Scan(
		&result.Id,
		&result.ClientPhoneNumber,
		&result.AnalysName,
		&result.AnalysUrl,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &patient.AnalysisInfo{}, err
	} else if err != nil {
		log.Println(err.Error())
		return &patient.AnalysisInfo{}, err
	}

	return &result, nil
}

// Doctor reports
func (r *patientRepo) DoctorReportCreate(req *patient.DoctorReportInfo) (*patient.DoctorReportInfo, error) {
	var result patient.DoctorReportInfo
	query := `
		INSERT INTO doctor_reports(
			id,
			doctor_id,
			text,
			patient_id
		) VALUES($1, $2, $3, $4)
		RETURNING
			id,
			COALESCE(doctor_id,'') as doctor_id,
			COALESCE(text,'') as text,
			COALESCE(patient_id,'') as patient_id,
			created_at,
			updated_at
		`

	if err := r.db.QueryRow(query, req.Id, req.DoctorId, req.Text, req.ClientId).Scan(
		&result.Id,
		&result.DoctorId,
		&result.Text,
		&result.ClientId,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		log.Println(err.Error())
		return &patient.DoctorReportInfo{}, err
	}
	return &result, nil
}

func (r *patientRepo) DoctorReportGet(req *patient.PatientId) (*patient.DoctorReportInfo, error) {
	var result patient.DoctorReportInfo

	query := `
		SELECT
			id,
			COALESCE(doctor_id,'') as doctor_id,
			COALESCE(text,'') as text,
			COALESCE(patient_id,'') as patient_id,			
			created_at,
			updated_at
		FROM doctor_reports
		WHERE patient_id=$1 and deleted_at IS NULL
	`
	err := r.db.QueryRow(query, req.ClientId).Scan(
		&result.Id,
		&result.DoctorId,
		&result.Text,
		&result.ClientId,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &patient.DoctorReportInfo{}, err
	} else if err != nil {
		log.Println(err.Error())
		return &patient.DoctorReportInfo{}, err
	}

	return &result, nil
}

// Queue
func (r *patientRepo) CreatePatientQueue(req *patient.CreatePatientQueueReq) (*patient.PatientQueueResp, error) {
	var (
		result patient.PatientQueueResp
	)
	query := `
		INSERT INTO queues(
			id,
			client_id,
			queue_number,
			service_id,
			service_type,
			turn_passed
		)values($1,$2,$3,$4,$5,false)
		RETURNING
			id,
			client_id,
			queue_number,
			service_id,
			service_type,
			turn_passed,
			created_at,
			updated_at
	`

	if err := r.db.QueryRow(query, req.Id, req.ClientId, req.QueueNumber, req.ServiceId, req.ServiceType).Scan(
		&result.Id,
		&result.ClientId,
		&result.QueueNumber,
		&result.ServiceId,
		&result.ServiceType,
		&result.TurnPassed,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &patient.PatientQueueResp{}, err
	}
	return &result, nil
}

func (r *patientRepo) GetPatientQueue(req *patient.PaymentHistoryId) (*patient.PatientQueueResp, error) {
	var (
		result patient.PatientQueueResp
	)
	query := `
		SELECT
			id,
			client_id,
			queue_number,
			service_id,
			service_type,
			turn_passed,
			created_at,
			updated_at
		FROM queues 
		WHERE id=$1
	`
	if err := r.db.QueryRow(query, req.Id).Scan(
		&result.Id,
		&result.ClientId,
		&result.QueueNumber,
		&result.ServiceId,
		&result.ServiceType,
		&result.TurnPassed,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &patient.PatientQueueResp{}, err
	}

	return &result, nil
}

func (r *patientRepo) CheckServiceQueue(req *patient.CheckQueueReq) (*patient.QueueNumber, error) {
	var (
		result int64
	)

	query := `
		SELECT 
			count(1) 
		FROM queues
		WHERE
			service_id = $1 AND service_type = $2 AND
			DATE(NOW())-DATE(created_at)=0 AND deleted_at IS NULL
	`
	row := r.db.QueryRow(query, req.ServiceId, req.ServiceType)

	err := row.Scan(
		&result,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &patient.QueueNumber{}, err
	} else if err != nil {
		log.Println(err.Error())
		return &patient.QueueNumber{}, err
	}

	return &patient.QueueNumber{QueueNumber: result}, nil
}

func (r *patientRepo) UpdateQueue(req *patient.UpdateQueueReq) (*patient.PatientQueueResp, error) {
	var (
		result patient.PatientQueueResp
	)
	query := `
		UPDATE queues SET
			turn_passed = NOT turn_passed,
			updated_at = NOW()
		WHERE
			client_id = $1 AND service_id = $2 AND service_type = $3
		RETURNING
			id,
			client_id,
			queue_number,
			service_id,
			service_type,
			turn_passed,
			created_at,
			updated_at
	`
	row := r.db.QueryRow(query, req.ClientId, req.ServiceId, req.ServiceType)
	if err := row.Scan(
		&result.Id,
		&result.ClientId,
		&result.QueueNumber,
		&result.ServiceId,
		&result.ServiceType,
		&result.TurnPassed,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &patient.PatientQueueResp{}, err
	}
	return &result, nil
}

func (r *patientRepo) FindQueue(req *patient.QueueFilter) (*patient.QueuesResp, error) {
	var (
		result patient.QueuesResp
	)
	offset := (req.Page - 1) * req.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)

	filter := fmt.Sprintf("WHERE service_id='%s' AND service_type='%s' AND turn_passed=FALSE AND deleted_at IS NULL", req.ServiceId, req.ServiceType)
	if req.ClientId != 0 {
		filter += fmt.Sprintf(` AND client_id = %d`, req.ClientId)
	}

	query := `
	SELECT 
		COALESCE(id,'95233a87-add9-4da7-a788-901ff3170ecd') as id,
		COALESCE(client_id,0) as client_id,
		COALESCE(queue_number,0) as queue_number,
		COALESCE(service_id,'95233a87-add9-4da7-a788-901ff3170ecd') as service_id,
		COALESCE(service_type,'')as service_id,
		COALESCE(turn_passed,false) as turn_passed,
		created_at,
		updated_at
	FROM queues
		` + filter + ` ORDER BY created_at asc ` + limit

	rows, err := r.db.Query(query)
	if err != nil {
		if err.Error() == "pq: function coalece(uuid, unknown) does not exist" {
			return &patient.QueuesResp{}, nil
		}
		return &patient.QueuesResp{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var temp patient.PatientQueueResp
		if err := rows.Scan(
			&temp.Id,
			&temp.ClientId,
			&temp.QueueNumber,
			&temp.ServiceId,
			&temp.ServiceType,
			&temp.TurnPassed,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		); err != nil {
			return &patient.QueuesResp{}, err
		}
		result.Queues = append(result.Queues, &temp)
	}

	queryCount := `SELECT count(1) FROM queues ` + filter
	err = r.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return &patient.QueuesResp{}, err
	}

	return &result, nil
}

func (r *patientRepo) CreateCashbox(req *patient.CashboxResp) (*patient.CashboxResp, error) {
	var (
		result                          patient.CashboxResp
		DoctorsIds, LabsIds, AparatsIds string
	)

	query := `
		INSERT INTO cashbox(
			id,
			client_id,
			summa,
			is_payed,
			cash_count,
			payment_type,
			doctors_ids,
			aparats_ids,
			labs_ids
		)values($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING
			COALESCE(id,'95233a87-add9-4da7-a788-901ff3170ecd') as id,
			COALESCE(client_id,0) as client_id,
			COALESCE(summa,0) as summa,
			COALESCE(is_payed,false) as is_payed,
			COALESCE(cash_count,0) as cash_count,
			COALESCE(payment_type,'') as payment_type,
			doctors_ids,
			labs_ids,
			aparats_ids,
			COALESCE(created_at,'1900-01-01') as created_at,
			COALESCE(updated_at,'1900-01-01') as created_at
	`

	row := r.db.DB.QueryRow(query, req.Id, req.ClientId, req.Summa,
		req.IsPayed, req.CashCount, req.PaymentType,
		pq.Array(req.DoctorsIds), pq.Array(req.AparatsIds), pq.Array(req.LabsIds))

	if err := row.Scan(
		&result.Id,
		&result.ClientId,
		&result.Summa,
		&result.IsPayed,
		&result.CashCount,
		&result.PaymentType,
		&DoctorsIds,
		&LabsIds,
		&AparatsIds,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &patient.CashboxResp{}, err
	}

	result.DoctorsIds = strings.Split(string(DoctorsIds), ",")
	result.AparatsIds = strings.Split(string(AparatsIds), ",")
	result.LabsIds = strings.Split(string(LabsIds), ",")

	return &result, nil
}

func (r *patientRepo) FindCashbox(req *patient.FindCashboxReq, ClientIds []int) (*patient.FindCashboxResp, error) {

	var (
		result patient.FindCashboxResp
	)

	offset := (req.Page - 1) * req.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)

	filter := " WHERE deleted_at IS NULL "
	if req.ClientId != 0 || len(ClientIds) != 0 {
		if req.ClientId != 0 {
			ClientIds = append(ClientIds, int(req.ClientId))
		}
		for index, client_id := range ClientIds {
			if index == 0 {
				filter += fmt.Sprintf(` AND (client_id=%d `, client_id)
			} else {
				filter += fmt.Sprintf(` OR client_id=%d `, client_id)
			}
		}
		filter += ` ) `
	}

	if req.FromDate != "" || req.ToDate != "" {
		if req.FromDate == "" {
			req.FromDate = "1000-01-01"
		}
		if req.ToDate == "" {
			req.ToDate = "3000-01-01"
		}
		filter += fmt.Sprintf(" AND (%s <= created_at AND created_at<= %s) ", req.FromDate, req.ToDate)
	}

	query := `
		SELECT 
			id,
			client_id,
			summa,
			is_payed,
			cash_count,
			payment_type,
			doctors_ids,
			labs_ids,
			aparats_ids,
			created_at,
			updated_at
		FROM cashbox
	` + filter + ` ORDER BY created_at asc ` + limit

	rows, err := r.db.Query(query)
	if err != nil {
		if err.Error() == "pq: function coalece(uuid, unknown) does not exist" {
			return &patient.FindCashboxResp{}, nil
		}
		return &patient.FindCashboxResp{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var DoctorsIds, LabsIds, AparatsIds string
		var temp patient.CashboxResp
		if err := rows.Scan(
			&temp.Id,
			&temp.ClientId,
			&temp.Summa,
			&temp.IsPayed,
			&temp.CashCount,
			&temp.PaymentType,
			&DoctorsIds,
			&LabsIds,
			&AparatsIds,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		); err != nil {
			return &patient.FindCashboxResp{}, err
		}

		temp.DoctorsIds = strings.Split(string(DoctorsIds), ",")
		temp.AparatsIds = strings.Split(string(AparatsIds), ",")
		temp.LabsIds = strings.Split(string(LabsIds), ",")

		result.Cashboxes = append(result.Cashboxes, &temp)
	}

	queryCount := `SELECT count(1) FROM cashbox ` + filter
	err = r.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		fmt.Println(743)
		return &patient.FindCashboxResp{}, err
	}
	return &result, nil
}

func (r *patientRepo) GetCashbox(req *patient.GetCashboxReq) (*patient.CashboxResp, error) {
	var (
		result                          patient.CashboxResp
		DoctorsIds, LabsIds, AparatsIds string
	)

	query := `
	SELECT 
		id,
		client_id,
		summa,
		is_payed,
		cash_count,
		payment_type,
		doctors_ids,
		labs_ids,
		aparats_ids,
		created_at,
		updated_at
	FROM cashbox
	WHERE id=$1 AND deleted_at IS NULL
	`
	row := r.db.QueryRow(query, req.CashboxId)
	if err := row.Scan(
		&result.Id,
		&result.ClientId,
		&result.Summa,
		&result.IsPayed,
		&result.CashCount,
		&result.PaymentType,
		&DoctorsIds,
		&LabsIds,
		&AparatsIds,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &patient.CashboxResp{}, err
	}

	result.DoctorsIds = strings.Split(string(DoctorsIds), ",")
	result.AparatsIds = strings.Split(string(AparatsIds), ",")
	result.LabsIds = strings.Split(string(LabsIds), ",")

	return &result, nil
}

func (r *patientRepo) UpdateCashbox(req *patient.UpdateCashboxReq) (*patient.CashboxResp, error) {
	var (
		result                          patient.CashboxResp
		DoctorsIds, LabsIds, AparatsIds string
	)

	query := `
	UPDATE cashbox SET 
		is_payed=$1,
		payment_type=$2,
		updated_at=NOW()
	WHERE id=$3 AND deleted_at IS NULL
	RETURNING
		id,
		client_id,
		summa,
		is_payed,
		cash_count,
		payment_type,
		doctors_ids,
		labs_ids,
		aparats_ids,
		created_at,
		updated_at
	`

	row := r.db.QueryRow(query, req.IsPayed, req.PaymentType, req.Id)
	if err := row.Scan(
		&result.Id,
		&result.ClientId,
		&result.Summa,
		&result.IsPayed,
		&result.CashCount,
		&result.PaymentType,
		&DoctorsIds,
		&LabsIds,
		&AparatsIds,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &patient.CashboxResp{}, err
	}
	result.DoctorsIds = strings.Split(string(DoctorsIds), ",")
	result.AparatsIds = strings.Split(string(AparatsIds), ",")
	result.LabsIds = strings.Split(string(LabsIds), ",")

	return &result, nil
}

func (r *patientRepo) DeleteCashbox(req *patient.GetCashboxReq) error {
	query := `
		UPDATE 
			cashbox
		SET
			deleted_at = NOW()
		WHERE id = $1
	`
	effect, err := r.db.Exec(query, req.CashboxId)
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

func (r *patientRepo) CreatePaymentHistory(req *patient.CreatePaymentHistoryReq) (*patient.PaymentHistoryResp, error) {
	var (
		result patient.PaymentHistoryResp
	)

	query := `
		INSERT INTO payment_history(
			id,
			client_id,
			summa,
			payment_type,
			cashbox_id
		)VALUES($1,$2,$3,$4,$5)
		RETURNING
			COALESCE(id,'aaa6662a-523a-45ca-bb93-53198f8077fe') as id,
			COALESCE(client_id,0) as client_id,
			COALESCE(summa,0) as summa,
			COALESCE(payment_type,'') as payment_type,
			COALESCE(cashbox_id,'aaa6662a-523a-45ca-bb93-53198f8077fe') as cashbox_id,
			created_at,
			updated_at
	`

	row := r.db.QueryRow(query, req.Id, req.ClientId, req.Summa, req.PaymentType, req.CashboxId)
	if err := row.Scan(
		&result.Id,
		&result.ClientId,
		&result.Summa,
		&result.PaymentType,
		&result.CashboxId,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &patient.PaymentHistoryResp{}, nil
	}

	return &result, nil
}

func (r *patientRepo) GetPaymentHistory(req *patient.PaymentHistoryId) (*patient.PaymentHistoryResp, error) {
	var (
		result patient.PaymentHistoryResp
	)

	query := `
		SELECT
			id,
			client_id,
			summa,
			payment_type,
			cashbox_id,
			created_at,
			updated_at
		FROM payment_history
		WHERE id = $1 AND deleted_at IS NULL 
	`

	row := r.db.QueryRow(query, req.Id)
	if err := row.Scan(
		&result.Id,
		&result.ClientId,
		&result.Summa,
		&result.PaymentType,
		&result.CashboxId,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &patient.PaymentHistoryResp{}, err
	}

	return &result, nil
}

func (r *patientRepo) FindPaymentHistory(req *patient.PaymentHistoryFilter) (*patient.PaymentHistoriesResp, error) {
	var (
		result patient.PaymentHistoriesResp
	)
	offset := (req.Page - 1) * req.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)

	filter := " WHERE deleted_at IS NULL "
	if req.ClientId != 0 {
		filter += fmt.Sprintf("AND client_id=%d ", req.ClientId)
	}

	if req.FromDate != "" || req.ToDate != "" {
		if req.FromDate == "" {
			req.FromDate = "1000-01-01"
		}
		if req.ToDate == "" {
			req.ToDate = "3000-01-01"
		}
		filter += fmt.Sprintf(" AND (%s <= created_at AND created_at<= %s) ", req.FromDate, req.ToDate)
	}

	query := `
		SELECT
			id,
			client_id,
			summa,
			payment_type,
			cashbox_id,
			created_at,
			updated_at
		FROM payment_history
	` + filter + ` ORDER BY created_at asc ` + limit

	rows, err := r.db.Query(query)
	if err != nil {
		return &patient.PaymentHistoriesResp{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var temp patient.PaymentHistoryResp
		if err := rows.Scan(
			&temp.Id,
			&temp.ClientId,
			&temp.Summa,
			&temp.PaymentType,
			&temp.CashboxId,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		); err != nil {
			return &patient.PaymentHistoriesResp{}, err
		}
		result.PaymentHistory = append(result.PaymentHistory, &temp)
	}
	queryCount := `SELECT count(1) FROM payment_history ` + filter
	err = r.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		fmt.Println(743)
		return &patient.PaymentHistoriesResp{}, err
	}
	return &result, nil
}

func (r *patientRepo) DeletePaymentHistory(req *patient.PaymentHistoryId) error {
	query := `
		UPDATE 
			payment_history
		SET
			deleted_at = NOW()
		WHERE id = $1
	`
	effect, err := r.db.Exec(query, req.Id)
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
