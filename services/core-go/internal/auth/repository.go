package auth

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) CreateUser(email, passwordHash string) error {
	query := `
	INSERT INTO users (email, password_hash) VALUES ($1, $2)		
	`
	_, err := r.DB.Exec(query, email, passwordHash)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return ErrUserAlreadyExists
		}
	}
	return err
}

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	query := `
		select id,email,password_hash, role from users 
		where email = $1
	`
	var user User

	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Role)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}

	return &user, err
}

func (r *Repository) StoreRefreshToken(userID int, token string) error {
	query := `
		INSERT INTO refresh_tokens (user_id, token, expires_at)
		VALUES ($1, $2, $3)
	`
	expiresAt := time.Now().Add(30 * 24 * time.Hour)
	_, err := r.DB.Exec(query, userID, token, expiresAt)
	return err
}

func (r *Repository) IsRefreshTokenValid(userID int, token string) (bool, error) {
	query := `
		SELECT 1 FROM refresh_tokens
		WHERE user_id = $1 AND token = $2 AND expires_at > NOW()
		LIMIT 1
	`
	var dummy int
	err := r.DB.QueryRow(query, userID, token).Scan(&dummy)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *Repository) RevokeRefreshToken(userID int, token string) error {
	query := `
		DELETE FROM refresh_tokens
		WHERE user_id = $1 AND token = $2
	`
	_, err := r.DB.Exec(query, userID, token)
	return err
}
