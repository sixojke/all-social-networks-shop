package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
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

func (r *UsersPostgres) Create(user *domain.User, code string) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("error create tx: %v", err)
	}

	query := fmt.Sprintf(`
	INSERT INTO %s 
		(username, password, email, balance, last_visit_at)
	VALUES
		($1, $2, $3, $4, $5)
	RETURNING id`, users)

	var id int
	if err := tx.QueryRow(query, user.Username, user.Password, user.Email,
		user.Balance, user.LastVisitAt).Scan(&id); err != nil {
		tx.Rollback()

		if strings.Contains(err.Error(), "duplicate key value") {
			return 0, domain.ErrDuplicateKey
		}

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
		(user_id, code)
	VALUES
		($1, $2)`, verification)
	if _, err := tx.Exec(query, id, code); err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error insert base verification: %v", err)
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
		id, username, email, balance, role
	FROM %s 
	WHERE username = $1 AND password = $2`, users)

	var user domain.User
	if err := r.db.Get(&user, query, username, password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}

		return nil, fmt.Errorf("error get user by credentials: %v", err)
	}

	query = fmt.Sprintf(`
	SELECT verified 
	FROM %s
	WHERE user_id = $1`, verification)

	var verified bool
	if err := r.db.Get(&verified, query, user.Id); err != nil {
		return nil, fmt.Errorf("error get user verified: %v", err)
	}

	if !verified {
		return nil, domain.ErrUserNotVerified
	}

	return &user, nil
}

func (r *UsersPostgres) GetByRefreshToken(refreshToken string) (*domain.Session, error) {
	query := fmt.Sprintf(`
	SELECT
		sessions.user_id, users.role
	FROM %s
	INNER JOIN %s ON sessions.user_id = users.id
	WHERE refresh_token = $1 AND expires_at > NOW()`, sessions, users)

	var session domain.Session
	if err := r.db.Get(&session, query, refreshToken); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}

		return nil, fmt.Errorf("error get user by refresh token: %v", err)
	}

	return &session, nil
}

func (r *UsersPostgres) Verify(userId int, code string) error {
	query := fmt.Sprintf(`
	UPDATE %s
	SET code = '', verified = true
	WHERE user_id = $1 AND code = $2`, verification)

	result, err := r.db.Exec(query, userId, code)
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

func (r *UsersPostgres) GetById(id int) (*domain.User, error) {
	query := fmt.Sprintf(`
	SELECT 
		id, username, role, balance, email
	FROM %s
	WHERE id = $1`, users)

	var user domain.User
	if err := r.db.Get(&user, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}

		return nil, fmt.Errorf("get user by id: %v", err)
	}

	return &user, nil
}

func (r *UsersPostgres) Ban(id int, banStatus bool) error {
	query := fmt.Sprintf(`
	INSERT INTO %s
		(id, status)
	VALUES 
		($1, $2)
	WHERE NOT EXISTS (
		SELECT 1
		FROM %s
		WHERE id = $1
	);`, bannedUsers, bannedUsers)

	result, err := r.db.Exec(query, id, banStatus)
	if err != nil {
		return fmt.Errorf("insert banned user: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %v", err)
	}

	if rows > 0 {
		query = fmt.Sprintf(`
		UPDATE %s
		SET status = $1
		WHERE id = $2`, bannedUsers)

		if _, err := r.db.Exec(query, banStatus, id); err != nil {
			return fmt.Errorf("update banned user status: %v", err)
		}
	}

	return nil
}
