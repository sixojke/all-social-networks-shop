package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TwoFaPostgres struct {
	db *sqlx.DB
}

func NewTwoFaPostgres(db *sqlx.DB) *TwoFaPostgres {
	return &TwoFaPostgres{db: db}
}

func (r *TwoFaPostgres) CreatePairingLink(userId int, secret_code string) error {
	query := fmt.Sprintf(`
	SELECT COUNT(*)
	FROM %s
	WHERE user_id = $1`, twoFa)

	var rows int
	if err := r.db.QueryRow(query, userId).Scan(&rows); err != nil {
		return fmt.Errorf("select count: %v", err)
	}

	if rows == 0 {
		query = fmt.Sprintf(`
		INSERT INTO %s
			(user_id, is_active, secret_code)
		VALUES
			($1, $2, $3)`, twoFa)

		if _, err := r.db.Exec(query, userId, false, secret_code); err != nil {
			return fmt.Errorf("insert secret_code: %v", err)
		}
	} else if rows == 1 {
		query := fmt.Sprintf(`
		UPDATE %s
		SET secret_code = $1
		WHERE user_id = $2`, twoFa)

		if _, err := r.db.Exec(query, secret_code, userId); err != nil {
			return fmt.Errorf("update secret_code: %v", err)
		}
	} else {
		return fmt.Errorf("unknown error")
	}

	return nil
}

func (r *TwoFaPostgres) GetSecretCode(userId int) (string, error) {
	query := fmt.Sprintf(`
	SELECT 
		secret_code 
	FROM %s
	WHERE user_id = $1`, twoFa)

	var secretCode string
	if err := r.db.Get(&secretCode, query, userId); err != nil {
		return "", fmt.Errorf("select secret code: %v", err)
	}

	return secretCode, nil
}
