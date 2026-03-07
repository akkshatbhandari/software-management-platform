package audit

import "database/sql"

type Repository struct{
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateAuditLog(userID int, action string, resource string, resourceID int) error {
	query := `
	insert into audit_logs (user_id, action, resource, resource_id)
	values ($1, $2, $3, $4)
	`
	_, err := r.DB.Exec(query, userID, action, resource, resourceID)
	return err
}