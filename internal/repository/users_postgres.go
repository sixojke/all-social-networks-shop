package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sixojke/internal/domain"
)

type UsersPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{
		db: db,
	}
}

func (r *UsersPostgres) Create(user *domain.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("error create tx: %v", err)
	}

	query := fmt.Sprintf(`
	INSERT INTO %s 
		(username, password, email, balance, last_visit_at)
	VALUES
		(:username, :password, :email, :balance, :last_visit_at)`, users)

	var id int
	if err := tx.QueryRow(query, user).Scan(&id); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error insert user: %v", err)
	}

	query = fmt.Sprintf(`
	INSERT INTO %s
		(refresh_token, expires_at, user_id)
	VALUES
		('nil', $1, $2)`, sessions)

	if _, err := tx.Exec(query, time.Now(), id); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error insert base user session: %v", err)
	}

	query = fmt.Sprintf(`
	INSERT INTO %s
		(user_id)
	VALUES
		($1)`, verification)
	if _, err := tx.Exec(query, id); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error insert base user session: %v", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error commit tx: %v", err)
	}

	return id, nil
}

func (r *UsersPostgres) GetByCredentials(username, password string) (*domain.User, error) {
	query := fmt.Sprintf(`
	SELECT 
		id, username, email, balance, verified
	FROM %s 
	WHERE username = $1 AND password = $2`, users)

	var user domain.User
	if err := r.db.Get(&user, query, username, password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}

		return nil, fmt.Errorf("error get user by credentials: %v", err)
	}

	return &user, nil
}

func (r *UsersPostgres) GetByRefreshToken(refreshToken string) (*domain.User, error) {
	query := fmt.Sprintf(`
	SELECT
		user_id
	FROM %s
	WHERE refresh_token = $1 AND expires_at > NOW()`, sessions)

	var user domain.User
	if err := r.db.Get(&user, query, refreshToken); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (r *UsersPostgres) Verify(userId int, code string) error {
	query := fmt.Sprintf(`
	UPDATE %s
	SET code = $1 AND SET verified = true
	WHERE user_id = $2`, sessions)

	result, err := r.db.Exec(query, code, userId)
	if err != nil {
		return fmt.Errorf("error update session: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error rows affected: %v", err)
	}

	if rows == 0 {
		return domain.ErrVerificationCodeInvalid
	}

	return nil
}

func (r *UsersPostgres) SetSession(session *domain.Session) error {
	query := fmt.Sprintf(`
	UPDATE %s
	SET refresh_token = :refresh_token, expires_at = :expires_at
	WHERE user_id = :user_id
	`, sessions)

	if _, err := r.db.NamedExec(query, session); err != nil {
		return fmt.Errorf("error update session: %v", err)
	}

	return nil
}
