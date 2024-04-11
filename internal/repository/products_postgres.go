package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sixojke/internal/domain"
)

type ProductsPostgres struct {
	db *sqlx.DB
}

func NewProductsPostgres(db *sqlx.DB) *ProductsPostgres {
	return &ProductsPostgres{db: db}
}

func (r *ProductsPostgres) Create(product *domain.Product) (int, error) {
	query := fmt.Sprintf(`
    INSERT INTO %s
        (name, price, quantity, description, category_id, subcategory_id)
    VALUES
        (:name, :price, :quantity, :description, :category_id, :subcategory_id)
    `, products,
	)

	var id int
	if err := r.db.QueryRow(query, product).Scan(&id); err != nil {
		return 0, fmt.Errorf("create product: %v", err)
	}

	return id, nil
}
