package user

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func RegistorUser(tx *sqlx.Tx, user CreateUserRequest) (error, sql.Result) {
	query := `INSERT INTO user(name,email,username,password,role) VALUES(name,email,username,password,role)`

	result, err := tx.NamedExec(query, user)
	if err != nil {
		return err, result
	}

	return nil, result
}
