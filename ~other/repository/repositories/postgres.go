package repositories

import (
	"database/sql"
	"rep/interfaces"
)

type PostgresUserRepo struct {
	db *sql.DB
}

func (r *PostgresUserRepo) GetByID(id int) (*interfaces.User, error) {
	row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
	var user interfaces.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
