package mysql

import (
	"database/sql"
	"github.com/anggardagasta/mit/models"
	"github.com/anggardagasta/mit/service"
)

type serviceUsersRepository struct {
	DB *sql.DB
}

func NewServiceUsersRepository(db *sql.DB) service.IServiceUsersRepository {
	return serviceUsersRepository{DB: db}
}

func (repo serviceUsersRepository) GetUserByPhoneNumber(input models.FormPhone) (result models.FormUser, err error) {
	q := "SELECT id, first_name, last_name, phone_number, birth_date, gender, email FROM users WHERE phone_number = ?"
	err = repo.DB.QueryRow(q, input.PhoneNumber).Scan(&result.Id,
		&result.FirstName, &result.LastName, &result.PhoneNumber, &result.BirthDate, &result.Gender, &result.Email)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (repo serviceUsersRepository) GetUserByEmail(input models.FormEmail) (result models.FormUser, err error) {
	q := `SELECT id, first_name, last_name, phone_number, birth_date, gender, email FROM users WHERE email = ?`
	rows, err := repo.DB.Query(q, input.Email)
	if err != nil {
		return result, err
	}
	for rows.Next() {
		if err := rows.Scan(&result); err != nil {
			return result, err
		}
	}
	_ = rows.Close()
	return result, nil
}

func (repo serviceUsersRepository) InsertUser(input models.FormUser) error {
	q := `INSERT INTO users (first_name, last_name, phone_number, birth_date, gender, email) VALUES (?,?,?,?,?,?)`
	rows, err := repo.DB.Query(q, input.FirstName, input.LastName, input.PhoneNumber, input.BirthDate, input.Gender, input.Email)
	if err != nil {
		return err
	}
	_ = rows.Close()
	return nil
}
