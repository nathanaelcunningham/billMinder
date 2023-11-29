package db

import (
	"github.com/VauntDev/tqla"
	"github.com/nathanaelcunningham/billReminder/models"
)

type UserRepository struct {
	DB *DB
}

func NewUserRepository(db *DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByID(ID int64) (*models.User, error) {
	t, err := tqla.New()
	if err != nil {
		return nil, err
	}

	stmt, args, err := t.Compile(
		`SELECT * FROM users WHERE id = {{ . }}`, ID,
	)
	if err != nil {
		return nil, err
	}

	row := r.DB.QueryRow(stmt, args...)
	var user models.User
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *models.User) (int64, error) {
	t, err := tqla.New()
	if err != nil {
		return 0, err
	}

	stmt, args, err := t.Compile(
		`INSERT INTO users (name, email, password) VALUES ({{ .Name }}, {{ .Email }}, {{ .Password }})`,
		user,
	)
	if err != nil {
		return 0, err
	}

	res, err := r.DB.Exec(stmt, args...)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *UserRepository) Update(user *models.User) error {
	t, err := tqla.New()
	if err != nil {
		return err
	}

	stmt, args, err := t.Compile(
		`UPDATE users SET name = {{ .Name }}, email = {{ .Email }}, password = {{ .Password }} WHERE id = {{ .ID }}`,
		user,
	)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(stmt, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(ID int64) error {
	t, err := tqla.New()
	if err != nil {
		return err
	}

	stmt, args, err := t.Compile(
		`DELETE FROM users WHERE id = {{ . }}`, ID,
	)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(stmt, args...)
	if err != nil {
		return err
	}
	return nil
}
