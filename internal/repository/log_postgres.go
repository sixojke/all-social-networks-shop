package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sixojke/internal/domain"
)

type LogPostgres struct {
	db *sqlx.DB
}

func NewLogPostgres(db *sqlx.DB) *LogPostgres {
	return &LogPostgres{
		db: db,
	}
}

func (r *LogPostgres) WriteAdminLog(log *domain.Log) error {
	query := fmt.Sprintf(`
	INSERT INTO %s
		(user_id, message)
	VALUES
		($1, $2)`, adminLogs)

	if _, err := r.db.Exec(query, log.UserId, log.Message); err != nil {
		return fmt.Errorf("insert log: %v", err)
	}

	return nil
}

func (r *LogPostgres) GetAdminLogs(limit int, offset int) (*domain.Pagination, error) {
	query := fmt.Sprintf(`
	SELECT u.username, a.message, a.created_at
	FROM %s a
	INNER JOIN %s u ON a.user_id = u.id
	ORDER BY created_at DESC
	LIMIT $1 OFFSET $2`, adminLogs, users)

	var logs []domain.GetAdminLogsOut
	if err := r.db.Select(&logs, query, limit, offset); err != nil {
		return nil, fmt.Errorf("select logs: %v", err)
	}

	query = fmt.Sprintf(`
	SELECT COUNT(*)
	FROM %s`, adminLogs)

	var totalItems int
	if err := r.db.QueryRow(query).Scan(&totalItems); err != nil {
		return nil, fmt.Errorf("select count items: %v", err)
	}

	return &domain.Pagination{
		Data:       logs,
		Limit:      limit,
		TotalItems: totalItems,
		TotalPages: getPages(totalItems, limit),
	}, nil
}

func getPages(totalItems, limit int) int {
	totalPages := 0
	if totalItems%limit == 0 {
		totalPages = totalItems / limit
	} else {
		totalPages = totalItems/limit + 1
	}

	return totalPages
}
