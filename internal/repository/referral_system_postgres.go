package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sixojke/internal/domain"
)

type ReferralSystemPostgres struct {
	db *sqlx.DB
}

func NewReferralLinksPostgres(db *sqlx.DB) *ReferralSystemPostgres {
	return &ReferralSystemPostgres{db: db}
}

func (r *ReferralSystemPostgres) CreateCode(ref domain.ReferralSystem) error {
	query := fmt.Sprintf(`
	INSERT INTO %s
		(referral_code, description)
	VALUES 
		($1, $2)`, referralSystem)

	if _, err := r.db.Exec(query, ref.ReferralCode, ref.Description); err != nil {
		return fmt.Errorf("insert referal system: %v", err)
	}

	return nil
}

func (r *ReferralSystemPostgres) AddVisitor(referralCode string) error {
	query := fmt.Sprintf(`
	UPDATE %s
	SET total_visitors = total_visitors + 1
	WHERE referral_code = $1`, referralSystem)

	if _, err := r.db.Exec(query, referralCode); err != nil {
		return fmt.Errorf("update total_visitors: %v", err)
	}

	return nil
}

func (r *ReferralSystemPostgres) GetStats(limit, offset int) (*domain.Pagination, error) {
	query := fmt.Sprintf(`
	SELECT *
	FROM %s
	LIMIT $1 OFFSET $2`, referralSystem)

	var stats []domain.ReferralSystem
	if err := r.db.Select(&stats, query, limit, offset); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &domain.Pagination{
				Data:       nil,
				Limit:      limit,
				TotalItems: 0,
				TotalPages: 0,
			}, nil
		}

		return nil, fmt.Errorf("select referral stats: %v", err)
	}

	query = fmt.Sprintf(`
	SELECT COUNT(*)
	FROM %s`, referralSystem)

	var totalItems int
	if err := r.db.QueryRow(query).Scan(&totalItems); err != nil {
		return nil, fmt.Errorf("select count items: %v", err)
	}

	return &domain.Pagination{
		Data:       &stats,
		Limit:      limit,
		TotalItems: totalItems,
		TotalPages: getPages(totalItems, limit),
	}, nil
}

func (r *ReferralSystemPostgres) DeleteCode(referralCode string) error {
	query := fmt.Sprintf(`
	DELETE
	FROM %s
	WHERE referral_code = $1`, referralSystem)

	if _, err := r.db.Exec(query, referralCode); err != nil {
		return fmt.Errorf("delete referral code: %v", err)
	}

	return nil
}
