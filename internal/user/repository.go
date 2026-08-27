package user

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func RegisterUser(tx *sqlx.Tx, user CreateUserRequest) (sql.Result, error) {
	query := `INSERT INTO user(name,email,username,password,role) VALUES(name,email,username,password,role)`

	result, err := tx.NamedExec(query, user)
	if err != nil {
		return result, err
	}

	return result, err
}
