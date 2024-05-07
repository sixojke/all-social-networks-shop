package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type BindPostgres struct {
	db *sqlx.DB
}

func NewBindPostgres(db *sqlx.DB) *BindPostgres {
	return &BindPostgres{
		db: db,
	}
}

func (r *BindPostgres) CreateAuthLink(code string, userId int) (string, error) {
	query := fmt.Sprintf(`
	SELECT 
		code
	FROM %s
	WHERE user_id = $1`, bindTelegram)

	var oldCode string
	if err := r.db.Get(&oldCode, query, userId); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("error select old link: %v", err)
		}
	}

	if oldCode == "" {
		query = fmt.Sprintf(`
		INSERT INTO %s
			(code, user_id)
		VALUES
			($1, $2)
		RETURNING
			code`, bindTelegram)

		if _, err := r.db.Exec(query, code, userId); err != nil {
			return "", fmt.Errorf("insert code: %v", err)
		}

		return code, nil
	}

	return oldCode, nil
}

func (r *BindPostgres) Bind(telegramId int, code string) (userId int, err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("create tx: %v", err)
	}

	query := fmt.Sprintf(`
	DELETE 
	FROM %s
	WHERE code = $1
	RETURNING
		user_id`, bindTelegram)

	if err := tx.QueryRow(query, code).Scan(&userId); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("delete auth code: %v", err)
	}

	query = fmt.Sprintf(`
	UPDATE %s
	SET telegram_id = $1
	WHERE id = $2`, users)

	if _, err := tx.Exec(query, telegramId, userId); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("add telegram id: %v", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("commit tx: %v", err)
	}

	return userId, nil
}

func (r *BindPostgres) Unbind(userId int) error {
	query := fmt.Sprintf(`
	UPDATE %s
	SET telegram_id = NULL
	WHERE id = $1`, users)

	if _, err := r.db.Exec(query, userId); err != nil {
		return fmt.Errorf("update telegram id: %v", err)
	}

	return nil
}
