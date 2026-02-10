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

func (r *Repository) Create(input CreateProjectInput, userID int) (Project, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return Project{}, err
	}

	defer tx.Rollback()

	var project Project

	err = tx.QueryRow(
		`INSERT INTO projects (name, description)
		VALUES ($1, $2)
		RETURNING id, name, description`,
		input.Name,
		input.Description,
	).Scan(&project.ID, &project.Name, &project.Description)

	if err != nil {
		return Project{}, err
	}

	if err := tx.Commit(); err != nil {
		return Project{}, err
	}

	return project, nil
}