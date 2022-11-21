package postgres

import (
	"fmt"

	"github.com/MuhammadyusufAdhamov/student/api/models"
	"github.com/MuhammadyusufAdhamov/student/storage/repo"
	"github.com/jmoiron/sqlx"
)

type studentRepo struct {
	db *sqlx.DB
}

func NewStudent(db *sqlx.DB) repo.StudentStorageI {
	return &studentRepo{
		db: db,
	}
}

func (sr *studentRepo) Create(student *models.CreateStudent) error {

	query := `
		insert into student(first_name,last_name,username,email,phone_number)
		values($1,$2,$3,$4,$5)
	`

	for _, students := range student.Students {
		_, err := sr.db.Exec(
			query,
			students.FirstName,
			students.LastName,
			students.Username,
			students.Email,
			students.PhoneNumber,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (sr *studentRepo) GetAll(params *repo.GetAllStudentsParams) (*repo.GetAllStudentsResult, error) {
	result := repo.GetAllStudentsResult{
		Students: make([]*repo.Student, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" limit %d offset %d ", params.Limit, offset)

	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(`
		where first_name ilike '%s' or last_name ilike '%s' or username ilike '%s' or email ilike '%s' or phone_number ilike '%s'`,
			str, str, str, str, str,
		)
	}

	query := `
		SELECT
			id,
			first_name,
			last_name,
			username,
			email,
			phone_number,
			created_at
		FROM users
		` + filter + `
		ORDER BY created_at desc
		` + limit

	rows, err := sr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u repo.Student

		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.PhoneNumber,
			&u.Email,
			&u.Username,
			&u.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		result.Students = append(result.Students, &u)
	}

	queryCount := `SELECT count(1) FROM users ` + filter
	err = sr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
