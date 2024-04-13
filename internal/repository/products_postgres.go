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

func (r *ProductsPostgres) GetById(id int) (*domain.Product, error) {
	query := fmt.Sprintf(`
  	SELECT *
  	FROM %s
  	WHERE id = $1`, products)

	var product domain.Product
	if err := r.db.Get(&product, query, id); err != nil {
		return nil, fmt.Errorf("select product: %v", err)
	}

	return &product, nil
}

func (r *ProductsPostgres) GetBySubcategory(id int) (*[]domain.Product, error) {
	query := fmt.Sprintf(`
	SELECT *
	FROM %s
	WHERE subcategory_id = $1`, products)

	var products []domain.Product
	if err := r.db.Select(&products, query, id); err != nil {
		return nil, fmt.Errorf("error select product: %v", err)
	}

	return &products, nil
}

// func (r *ProductsPostgres) Update(product *domain.Product) (*domain.Product, error) {

// }
