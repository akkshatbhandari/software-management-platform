package auth

import (
	"database/sql"
	"errors"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) CreateUser(email, passwordHash string) error {
	query:= `
	INSERT INTO users (email, password_hash) VALUES ($1, $2)		
	`
	_, err := r.DB.Exec(query, email, passwordHash)
	return err
}

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	query := `
		select id,email,password_hash, role from users 
		where email = $1
	`
	var user User

	err := r.DB.QueryRow(query,email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Role)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	return &user, err
}