package mysql

import (
	"database/sql"
	"github.com/anggardagasta/mit/models"
	"github.com/anggardagasta/mit/service"
	"github.com/go-sql-driver/mysql"
	"time"
)

type serviceUsersRepository struct {
	DB *sql.DB
}

func NewServiceUsersRepository(db *sql.DB) service.IServiceUsersRepository {
	return serviceUsersRepository{DB: db}
}

func (repo serviceUsersRepository) GetUserByPhoneNumber(input models.FormPhone) (result models.FormUserResponse, err error) {
	q := "SELECT id, first_name, last_name, phone_number, birth_date, gender, email FROM users WHERE phone_number = ?"
	rows, err := repo.DB.Query(q, input.PhoneNumber)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		if err := rows.Scan(&result.Id, &result.FirstName, &result.LastName, &result.PhoneNumber, &result.BirthDate,
			&result.Gender, &result.Email); err != nil {
			return result, err
		}
	}
	_ = rows.Close()
	return result, nil
}

func (repo serviceUsersRepository) GetUserByEmail(input models.FormEmail) (result models.FormUserResponse, err error) {
	q := `SELECT id, first_name, last_name, phone_number, birth_date, gender, email FROM users WHERE email = ?`
	rows, err := repo.DB.Query(q, input.Email)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		if err := rows.Scan(&result.Id, &result.FirstName, &result.LastName, &result.PhoneNumber, &result.BirthDate,
			&result.Gender, &result.Email); err != nil {
			return result, err
		}
	}
	_ = rows.Close()
	return result, nil
}

func (repo serviceUsersRepository) InsertUser(input models.FormUserRequest) (err error) {
	q := `INSERT INTO users (first_name, last_name, phone_number, birth_date, gender, email) VALUES (?,?,?,?,?,?)`
	if input.BirthDate != "" {
		birthDate, err := time.Parse("2006-01-02", input.BirthDate)
		if err != nil {
			return err
		}
		rows, err := repo.DB.Query(q, input.FirstName, input.LastName, input.PhoneNumber, birthDate, input.Gender, input.Email)
		if err != nil {
			return err
		}
		_ = rows.Close()
	} else {
		birthDateNull := mysql.NullTime{}
		rows, err := repo.DB.Query(q, input.FirstName, input.LastName, input.PhoneNumber, birthDateNull, input.Gender, input.Email)
		if err != nil {
			return err
		}
		_ = rows.Close()
	}

	return nil
}
