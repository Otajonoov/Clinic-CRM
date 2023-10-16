package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/clinic-crm/labs/genproto/lab"
	"gitlab.com/clinic-crm/labs/storage/repo"
	// "github.com/google/uuid"
)

type labRepo struct {
	db *sqlx.DB
}

func NewLab(db *sqlx.DB) repo.LabStorageI {
	return &labRepo{
		db: db,
	}
}

// Labs
func (lr *labRepo) LabCreate(req *lab.LabCreateReq) (*lab.LabCreateRes, error) {
	var result lab.LabCreateRes
	query := `
		INSERT INTO labs(
			id,
			name,
			price,
			type,
			sub_category_id
		) VALUES($1, $2, $3, $4, $5)
		RETURNING 
			id,
			COALESCE(name,'') as name,
			price,
			COALESCE(type,'') as type,
			sub_category_id,
			created_at,
			updated_at
		`
	if err := lr.db.DB.QueryRow(query,
		req.Id, req.Name, req.Price, req.Type, req.SubCategoryId,
	).Scan(
		&result.Id,
		&result.Name,
		&result.Price,
		&result.Type,
		&result.SubCategoryId,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &lab.LabCreateRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) LabGet(req *lab.LabGetReq) (*lab.LabCreateRes, error) {
	var result lab.LabCreateRes
	query := `
		SELECT 
			id,
			COALESCE(name,'') as name,
			price,
			COALESCE(type,'') as type,
			sub_category_id,
			created_at,
			updated_at
		FROM labs
		WHERE ` + req.Field + `='` + req.Value + `' and deleted_at IS NULL
	`
	err := lr.db.DB.QueryRow(query).Scan(
		&result.Id,
		&result.Name,
		&result.Price,
		&result.Type,
		&result.SubCategoryId,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &lab.LabCreateRes{}, err
	} else if err != nil {
		log.Fatalln(err.Error())
		return &lab.LabCreateRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) LabsFind(req *lab.LabsFindReq) (*lab.LabsRes, error) {
	result := lab.LabsRes{
		Labs: make([]*lab.LabCreateRes, 0),
	}

	offset := (req.Page - 1) * req.Limit
	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)
	filter := "WHERE deleted_at IS NULL"

	if req.Search != "" {
		str := "%" + req.Search + "%"
		filter += fmt.Sprintf(`
			AND (type ILIKE '%s' OR name ILIKE '%s' OR price ILIKE '%s')
				`,
			str, str, str,
		)
	}

	query := `
		SELECT
			id,
			COALESCE(name,'') as name,
			price,
			COALESCE(type,'') as type,
			sub_category_id,
			created_at,
			updated_at
		FROM labs
		` + filter + `
		ORDER BY created_at DESC
		` + limit
	rows, err := lr.db.Query(query)
	if err != nil {
		return &lab.LabsRes{}, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := lab.LabCreateRes{}
		err := rows.Scan(
			&temp.Id,
			&temp.Name,
			&temp.Price,
			&temp.Type,
			&temp.SubCategoryId,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return &lab.LabsRes{}, err
		}
		result.Labs = append(result.Labs, &temp)
	}
	queryCount := `SELECT COUNT(1) FROM labs ` + filter
	err = lr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return &lab.LabsRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) LabUpdate(req *lab.LabUpdateReq) (*lab.LabCreateRes, error) {
	var result lab.LabCreateRes
	query := `
		UPDATE
			labs
		SET
		    name=$1,
			price=$2,
			type=$3,
			updated_at=NOW()
		WHERE
			id=$4 AND deleted_at IS NULL
		RETURNING
			id,
			COALESCE(name,'') as name,
			price,
			COALESCE(type,'') as type,
			sub_category_id,
			created_at,
			updated_at
	`
	err := lr.db.DB.QueryRow(query,
		req.Name,
		req.Price,
		req.Type,
		req.Id,
	).Scan(
		&result.Id,
		&result.Name,
		&result.Price,
		&result.Type,
		&result.SubCategoryId,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		log.Println(err.Error())
		return &lab.LabCreateRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) LabDelete(req *lab.LabId) error {
	query := `
		UPDATE
			labs
		SET
			deleted_at = NOW()
		WHERE id = $1
	`
	effect, err := lr.db.Exec(query, req.Id)
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


// Aparats
func (lr *labRepo) AparatCreate(req *lab.AparatCreateReq) (*lab.AparatCreateRes, error) {
	var result lab.AparatCreateRes
	query := `
		INSERT INTO aparats(
			id,
			name,
			price,
			type,
			sub_category_id
		) VALUES($1, $2, $3, $4, $5)
		RETURNING 
			id,
			COALESCE(name,'') as name,
			price,
			COALESCE(type,'') as type,
			sub_category_id,
			created_at,
			updated_at
		`
	if err := lr.db.DB.QueryRow(query,
		req.Id, req.Name, req.Price, req.Type, req.SubCategoryId,
	).Scan(
		&result.Id,
		&result.Name,
		&result.Price,
		&result.Type,
		&result.SubCategoryId,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &lab.AparatCreateRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) AparatGet(req *lab.AparatGetReq) (*lab.AparatCreateRes, error) {
	var result lab.AparatCreateRes
	query := `
		SELECT 
			id,
			COALESCE(name,'') as name,
			price,
			COALESCE(type,'') as type,
			sub_category_id,
			created_at,
			updated_at
		FROM aparats
		WHERE ` + req.Field + `='` + req.Value + `' and deleted_at IS NULL
	`
	err := lr.db.DB.QueryRow(query).Scan(
		&result.Id,
		&result.Name,
		&result.Price,
		&result.Type,
		&result.SubCategoryId,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &lab.AparatCreateRes{}, err
	} else if err != nil {
		log.Fatalln(err.Error())
		return &lab.AparatCreateRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) AparatsFind(req *lab.AparatsFindReq) (*lab.AparatsRes, error) {
	result := lab.AparatsRes{
		Aparats	: make([]*lab.AparatCreateRes, 0),
	}

	offset := (req.Page - 1) * req.Limit
	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)
	filter := "WHERE deleted_at IS NULL"

	if req.Search != "" {
		str := "%" + req.Search + "%"
		filter += fmt.Sprintf(`
			AND (type ILIKE '%s' OR name ILIKE '%s' OR price ILIKE '%s')
				`,
			str, str, str,
		)
	}

	query := `
		SELECT
			id,
			COALESCE(name,'') as name,
			price,
			COALESCE(type,'') as type,
			sub_category_id,
			created_at,
			updated_at
		FROM aparats
		` + filter + `
		ORDER BY created_at DESC
		` + limit
	rows, err := lr.db.Query(query)
	if err != nil {
		return &lab.AparatsRes{}, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := lab.AparatCreateRes{}
		err := rows.Scan(
			&temp.Id,
			&temp.Name,
			&temp.Price,
			&temp.Type,
			&temp.SubCategoryId,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return &lab.AparatsRes{}, err
		}
		result.Aparats = append(result.Aparats, &temp)
	}
	queryCount := `SELECT COUNT(1) FROM labs ` + filter
	err = lr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return &lab.AparatsRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) AparatsUpdate(req *lab.AparatUpdateReq) (*lab.AparatCreateRes, error) {
	var result lab.AparatCreateRes
	query := `
		UPDATE
			aparats
		SET
		    name=$1,
			price=$2,
			type=$3,
			updated_at=NOW()
		WHERE
			id=$4 AND deleted_at IS NULL
		RETURNING
			id,
			COALESCE(name,'') as name,
			price,
			COALESCE(type,'') as type,
			sub_category_id,
			created_at,
			updated_at
	`
	err := lr.db.DB.QueryRow(query,
		req.Name,
		req.Price,
		req.Type,
		req.Id,
	).Scan(
		&result.Id,
		&result.Name,
		&result.Price,
		&result.Type,
		&result.SubCategoryId,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		log.Println(err.Error())
		return &lab.AparatCreateRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) AparatsDelete(req *lab.AparatId) error {
	fmt.Println(req)
	query := `
		UPDATE
			aparats
		SET
			deleted_at = NOW()
		WHERE id = $1
	`
	effect, err := lr.db.Exec(query, req.Id)
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

// Labs category
func (lr *labRepo) LabCategoryCreate(req *lab.Category) (*lab.CategoryRes, error) {
	var result lab.CategoryRes
	query := `
		INSERT INTO lab_category (
			id,
			name
		) VALUES($1, $2)
		RETURNING 
			id,
			COALESCE(name,'') as name,
			created_at,
			updated_at
		`
	if err := lr.db.DB.QueryRow(query,
		req.Id, req.Name,
	).Scan(
		&result.Id,
		&result.Name,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &lab.CategoryRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) LabCategoryGet(req *lab.CategoryGetReq) (*lab.CategoryRes, error) {
	var result lab.CategoryRes
	query := `
		SELECT 
			id,
			COALESCE(name,'') as name,
			created_at,
			updated_at
		FROM lab_category
		WHERE ` + req.Field + `='` + req.Value + `' and deleted_at IS NULL
	`
	err := lr.db.DB.QueryRow(query).Scan(
		&result.Id,
		&result.Name,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &lab.CategoryRes{}, err
	} else if err != nil {
		log.Fatalln(err.Error())
		return &lab.CategoryRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) LabCategoryFind(req *lab.CategoryFindReq) (*lab.CategoriesRes, error) {
	result := lab.CategoriesRes{
		Info: make([]*lab.CategoryRes, 0),
	}

	offset := (req.Page - 1) * req.Limit
	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)
	filter := "WHERE deleted_at IS NULL"

	if req.Search != "" {
		str := "%" + req.Search + "%"
		filter += fmt.Sprintf(`
			AND (name ILIKE '%s')
				`,
			str,
		)
	}

	query := `
		SELECT
			id,
			COALESCE(name,'') as name,
			created_at,
			updated_at
		FROM lab_category
		` + filter + `
		ORDER BY created_at DESC
		` + limit
	rows, err := lr.db.Query(query)
	if err != nil {
		return &lab.CategoriesRes{}, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := lab.CategoryRes{}
		err := rows.Scan(
			&temp.Id,
			&temp.Name,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return &lab.CategoriesRes{}, err
		}
		result.Info = append(result.Info, &temp)
	}
	queryCount := `SELECT COUNT(1) FROM labs ` + filter
	err = lr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return &lab.CategoriesRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) LabCategoryUpdate(req *lab.Category) (*lab.CategoryRes, error) {
	var result lab.CategoryRes
	query := `
		UPDATE
			lab_category
		SET
		    name=$1,
			updated_at=NOW()
		WHERE
			id=$2 AND deleted_at IS NULL
		RETURNING
			id,
			COALESCE(name,'') as name,
			created_at,
			updated_at
	`
	err := lr.db.DB.QueryRow(query,
		req.Name,
		req.Id,
	).Scan(
		&result.Id,
		&result.Name,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		log.Println(err.Error())
		return &lab.CategoryRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) LabCategoryDelete(req *lab.CategoryId) error {
	query := `
		UPDATE
			lab_category
		SET
			deleted_at = NOW()
		WHERE id = $1
	`
	effect, err := lr.db.Exec(query, req.Id)
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

// Aparat category
func (lr *labRepo) AparatCategoryCreate(req *lab.Category) (*lab.CategoryRes, error) {
	var result lab.CategoryRes
	query := `
		INSERT INTO aparat_category(
			id,
			name
		) VALUES($1, $2)
		RETURNING 
			id,
			COALESCE(name,'') as name,
			created_at,
			updated_at
		`
	if err := lr.db.DB.QueryRow(query,
		req.Id, req.Name,
	).Scan(
		&result.Id,
		&result.Name,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &lab.CategoryRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) AparatCategoryGet(req *lab.CategoryGetReq) (*lab.CategoryRes, error) {
	var result lab.CategoryRes
	query := `
		SELECT 
			id,
			COALESCE(name,'') as name,
			created_at,
			updated_at
		FROM aparat_category
		WHERE ` + req.Field + `='` + req.Value + `' and deleted_at IS NULL
	`
	err := lr.db.DB.QueryRow(query).Scan(
		&result.Id,
		&result.Name,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &lab.CategoryRes{}, err
	} else if err != nil {
		log.Fatalln(err.Error())
		return &lab.CategoryRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) AparatCategoryFind(req *lab.CategoryFindReq) (*lab.CategoriesRes, error) {
	result := lab.CategoriesRes{
		Info: make([]*lab.CategoryRes, 0),
	}

	offset := (req.Page - 1) * req.Limit
	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)
	filter := "WHERE deleted_at IS NULL"

	if req.Search != "" {
		str := "%" + req.Search + "%"
		filter += fmt.Sprintf(`
			AND (name ILIKE '%s')
				`,
			str,
		)
	}

	query := `
		SELECT
			id,
			COALESCE(name,'') as name,
			created_at,
			updated_at
		FROM aparat_category
		` + filter + `
		ORDER BY created_at DESC
		` + limit
	rows, err := lr.db.Query(query)
	if err != nil {
		return &lab.CategoriesRes{}, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := lab.CategoryRes{}
		err := rows.Scan(
			&temp.Id,
			&temp.Name,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return &lab.CategoriesRes{}, err
		}
		result.Info = append(result.Info, &temp)
	}
	queryCount := `SELECT COUNT(1) FROM aparat_category ` + filter
	err = lr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return &lab.CategoriesRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) AparatCategoryUpdate(req *lab.Category) (*lab.CategoryRes, error) {
	var result lab.CategoryRes
	query := `
		UPDATE
			aparat_category
		SET
		    name=$1,
			updated_at=NOW()
		WHERE
			id=$2 AND deleted_at IS NULL
		RETURNING
			id,
			COALESCE(name,'') as name,
			created_at,
			updated_at
	`
	err := lr.db.DB.QueryRow(query,
		req.Name,
		req.Id,
	).Scan(
		&result.Id,
		&result.Name,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		log.Println(err.Error())
		return &lab.CategoryRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) AparatCategoryDelete(req *lab.CategoryId) error {
	query := `
		UPDATE
			aparat_category
		SET
			deleted_at = NOW()
		WHERE id = $1
	`
	effect, err := lr.db.Exec(query, req.Id)
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

// Lab sub category
func (lr *labRepo) LabSubCategoryCreate(req *lab.SubCategory) (*lab.SubCategoryRes, error) {
	fmt.Println(req)
	var result lab.SubCategoryRes
	query := `
		INSERT INTO lab_sub_category(
			id,
			name,
			category_id
		) VALUES($1, $2, $3)
		RETURNING 
			id,
			name,
			category_id,
			created_at,
			updated_at
		`
	if err := lr.db.DB.QueryRow(query,
		req.Id, req.Name, req.CategoryId,
	).Scan(
		&result.Id,
		&result.Name,
		&result.CategoryId,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &lab.SubCategoryRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) LabSubCategoryGet(req *lab.CategoryGetReq) (*lab.SubCategoryRes, error) {
	var result lab.SubCategoryRes
	query := `
		SELECT 
			id,
			COALESCE(name,'') as name,
			category_id,
			created_at,
			updated_at
		FROM lab_sub_category
		WHERE ` + req.Field + `='` + req.Value + `' and deleted_at IS NULL
	`
	err := lr.db.DB.QueryRow(query).Scan(
		&result.Id,
		&result.Name,
		&result.CategoryId,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &lab.SubCategoryRes{}, err
	} else if err != nil {
		log.Fatalln(err.Error())
		return &lab.SubCategoryRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) LabSubCategoryFind(req *lab.SubCategoryFindReq) (*lab.SubCategoriesRes, error) {
	result := lab.SubCategoriesRes{
		Info: make([]*lab.SubCategoryRes, 0),
	}

	offset := (req.Page - 1) * req.Limit
	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)
	filter := "WHERE deleted_at IS NULL"

	if req.CategoryId != "" {
		filter += fmt.Sprintf(`
			AND (category_id = '%s')`,req.CategoryId,
		)
	}

	query := `
		SELECT
			id,
			COALESCE(name,'') as name,
			category_id,
			created_at,
			updated_at
		FROM lab_sub_category
		` + filter + `
		ORDER BY created_at DESC
		` + limit

	rows, err := lr.db.Query(query)
	if err != nil {
		return &lab.SubCategoriesRes{}, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := lab.SubCategoryRes{}
		err := rows.Scan(
			&temp.Id,
			&temp.Name,
			&temp.CategoryId,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return &lab.SubCategoriesRes{}, err
		}
		result.Info = append(result.Info, &temp)
	}
	queryCount := `SELECT COUNT(1) FROM lab_sub_category ` + filter
	err = lr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return &lab.SubCategoriesRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) LabSubCategoryUpdate(req *lab.SubCategory) (*lab.SubCategoryRes, error) {
	var result lab.SubCategoryRes
	query := `
		UPDATE
			lab_sub_category
		SET
		    name=$1,
			updated_at=NOW()
		WHERE
			id=$2 AND deleted_at IS NULL
		RETURNING
			id,
			COALESCE(name,'') as name,
			category_id,
			created_at,
			updated_at
	`
	err := lr.db.DB.QueryRow(query,
		req.Name,
		req.Id,
	).Scan(
		&result.Id,
		&result.Name,
		&result.CategoryId,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		log.Println(err.Error())
		return &lab.SubCategoryRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) LabSubCategoryDelete(req *lab.CategoryId) error {
	query := `
		UPDATE
			lab_sub_category
		SET
			deleted_at = NOW()
		WHERE id = $1
	`
	effect, err := lr.db.Exec(query, req.Id)
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

// Aparat sub category
func (lr *labRepo) AparatSubCategoryCreate(req *lab.SubCategory) (*lab.SubCategoryRes, error) {
	var result lab.SubCategoryRes
	query := `
		INSERT INTO aparat_sub_category(
			id,
			name,
			category_id
		) VALUES($1, $2, $3)
		RETURNING 
			id,
			COALESCE(name,'') as name,
			category_id,
			created_at,
			updated_at
		`
	if err := lr.db.DB.QueryRow(query,
		req.Id, req.Name, req.CategoryId,
	).Scan(
		&result.Id,
		&result.Name,
		&result.CategoryId,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &lab.SubCategoryRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) AparatSubCategoryGet(req *lab.CategoryGetReq) (*lab.SubCategoryRes, error) {
	var result lab.SubCategoryRes
	query := `
		SELECT 
			id,
			COALESCE(name,'') as name,
			category_id,
			created_at,
			updated_at
		FROM aparat_sub_category
		WHERE ` + req.Field + `='` + req.Value + `' and deleted_at IS NULL
	`
	err := lr.db.DB.QueryRow(query).Scan(
		&result.Id,
		&result.Name,
		&result.CategoryId,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &lab.SubCategoryRes{}, err
	} else if err != nil {
		log.Fatalln(err.Error())
		return &lab.SubCategoryRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) AparatSubCategoryFind(req *lab.SubCategoryFindReq) (*lab.SubCategoriesRes, error) {
	result := lab.SubCategoriesRes{
		Info: make([]*lab.SubCategoryRes, 0),
	}

	offset := (req.Page - 1) * req.Limit
	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)
	filter := "WHERE deleted_at IS NULL"

	if req.CategoryId != "" {
		filter += fmt.Sprintf(` AND category_id = '%s' `, req.CategoryId)	
	}

	query := `
		SELECT
			id,
			COALESCE(name,'') as name,
			category_id,
			created_at,
			updated_at
		FROM aparat_sub_category
		` + filter + `
		ORDER BY created_at DESC
		` + limit
	rows, err := lr.db.Query(query)
	if err != nil {
		return &lab.SubCategoriesRes{}, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := lab.SubCategoryRes{}
		err := rows.Scan(
			&temp.Id,
			&temp.Name,
			&temp.CategoryId,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return &lab.SubCategoriesRes{}, err
		}
		result.Info = append(result.Info, &temp)
	}
	queryCount := `SELECT COUNT(1) FROM aparat_sub_category ` + filter
	err = lr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return &lab.SubCategoriesRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) AparatSubCategoryUpdate(req *lab.SubCategory) (*lab.SubCategoryRes, error) {
	var result lab.SubCategoryRes
	query := `
		UPDATE
			aparat_sub_category
		SET
		    name=$1,
			updated_at=NOW()
		WHERE
			id=$2 AND deleted_at IS NULL
		RETURNING
			id,
			COALESCE(name,'') as name,
			category_id,
			created_at,
			updated_at
	`
	err := lr.db.DB.QueryRow(query,
		req.Name,
		req.Id,
	).Scan(
		&result.Id,
		&result.Name,
		&result.CategoryId,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		log.Println(err.Error())
		return &lab.SubCategoryRes{}, err
	}

	return &result, nil
}

func (lr *labRepo) AparatSubCategoryDelete(req *lab.CategoryId) error {
	query := `
		UPDATE
			aparat_sub_category
		SET
			deleted_at = NOW()
		WHERE id = $1
	`
	effect, err := lr.db.Exec(query, req.Id)
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


// Aparat analysis
func (lr *labRepo) AparatAnalysisCreate(req *lab.AnalysisReq) (*lab.AnalysisResp, error) {
	var result lab.AnalysisResp
	query := `
		INSERT INTO aparat_analysis(
			id,
			client_id,
			aparat_id,
			analysis_url
		) VALUES($1, $2, $3, $4)
		RETURNING 
			id,
			client_id,
			aparat_id,
			analysis_url,
			created_at,
			updated_at
		`
	if err := lr.db.DB.QueryRow(query,
		req.Id, req.ClientId, req.AparatId, req.AnalysisUrl,
	).Scan(
		&result.Id,
		&result.ClientId,
		&result.AparatId,
		&result.AnalysisUrl,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &lab.AnalysisResp{}, err
	}

	return &result, nil
}

func (lr *labRepo) AparatAnalysisGet(req *lab.AnalysisGetReq) (*lab.AnalysisResp, error) {
	var result lab.AnalysisResp
	query := `
		SELECT 
			id,
			client_id,
			aparat_id,
			analysis_url,
			created_at,
			updated_at
		FROM aparat_analysis
		WHERE ` + req.Field + `='` + req.Value + `' and deleted_at IS NULL
	`
	err := lr.db.DB.QueryRow(query).Scan(
		&result.Id,
		&result.ClientId,
		&result.AparatId,
		&result.AnalysisUrl,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &lab.AnalysisResp{}, err
	} else if err != nil {
		log.Fatalln(err.Error())
		return &lab.AnalysisResp{}, err
	}

	return &result, nil
}

func (lr *labRepo) AparatAnalysisDelete(req *lab.AparatId) error {
	query := `
		UPDATE
			aparat_analysis
		SET
			deleted_at = NOW()
		WHERE id = $1
	`
	effect, err := lr.db.Exec(query, req.Id)
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

// Lab analysis
func (lr *labRepo) LabAnalysisCreate(req *lab.AnalysisReq) (*lab.AnalysisResp, error) {
	var result lab.AnalysisResp
	query := `
		INSERT INTO lab_analysis(
			id,
			client_id,
			aparat_id,
			analysis_url
		) VALUES($1, $2, $3, $4)
		RETURNING 
			id,
			client_id,
			aparat_id,
			analysis_url,
			created_at,
			updated_at
		`
	if err := lr.db.DB.QueryRow(query,
		req.Id, req.ClientId, req.AparatId, req.AnalysisUrl,
	).Scan(
		&result.Id,
		&result.ClientId,
		&result.AparatId,
		&result.AnalysisUrl,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		return &lab.AnalysisResp{}, err
	}

	return &result, nil
}

func (lr *labRepo) LabAnalysisGet(req *lab.AnalysisGetReq) (*lab.AnalysisResp, error) {
	var result lab.AnalysisResp
	query := `
		SELECT 
			id,
			client_id,
			aparat_id,
			analysis_url,
			created_at,
			updated_at
		FROM lab_analysis
		WHERE ` + req.Field + `='` + req.Value + `' and deleted_at IS NULL
	`
	err := lr.db.DB.QueryRow(query).Scan(
		&result.Id,
		&result.ClientId,
		&result.AparatId,
		&result.AnalysisUrl,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &lab.AnalysisResp{}, err
	} else if err != nil {
		log.Fatalln(err.Error())
		return &lab.AnalysisResp{}, err
	}

	return &result, nil
}

func (lr *labRepo) LabAnalysisDelete(req *lab.AparatId) error {
	query := `
		UPDATE
			lab_analysis
		SET
			deleted_at = NOW()
		WHERE id = $1
	`
	effect, err := lr.db.Exec(query, req.Id)
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

