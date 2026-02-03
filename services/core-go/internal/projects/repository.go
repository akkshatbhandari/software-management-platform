package projects

import "database/sql"

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetAll() ([]Project, error) {
	rows, err := r.DB.Query(
		"SELECT id,name,description from projects",
	)

	if err != nil {
		return nil,err
	}

	defer rows.Close()

	var projects []Project

	for rows.Next() {
		var p Project

		if err:= rows.Scan(&p.ID, &p.Name, &p.Description); err != nil {
			return nil,err
		}

		projects = append(projects, p)
	}
	return projects, nil
}