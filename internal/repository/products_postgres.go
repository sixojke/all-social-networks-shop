package repository

import (
	"database/sql"
	"errors"
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

// func (r *ProductsPostgres) Create(product *domain.Product) (int, error) {
// 	query := fmt.Sprintf(`
//     INSERT INTO %s
//         (name, price, quantity, description, category_id, subcategory_id)
//     VALUES
//         ($1, $2, $3, $4, $5, $6)
//     `, products,
// 	)

// 	var id int
// 	if err := r.db.QueryRow(query, product.Name, product.Price, product.Quantity, product.Description,
// 		product.CategoryId, product.SubcategoryId).Scan(&id); err != nil {
// 		return 0, fmt.Errorf("create product: %v", err)
// 	}

// 	return id, nil
// }

func (r *ProductsPostgres) GetAll(limit, offset int) (*domain.Pagination, error) {
	query := fmt.Sprintf(`
	SELECT 
		products.id, products.name, products.description, products.price, products.quantity, 
		products.quantity_sales, products.category_id, products.uploaded_at, category.img_path
	FROM %s
	INNER JOIN %s ON products.category_id = category.id
	ORDER BY products.quantity_sales
	LIMIT $1 OFFSET $2`, products, category)

	var p []domain.Product
	if err := r.db.Select(&p, query, limit, offset); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrProductsNotFound
		}

		return nil, fmt.Errorf("get all products: %v", err)
	}

	query = fmt.Sprintf(`
	SELECT COUNT(*)
	FROM %s`, products)

	var rows int
	if err := r.db.QueryRow(query).Scan(&rows); err != nil {
		return nil, fmt.Errorf("error select count rows: %v", err)
	}

	return &domain.Pagination{
		Data:   p,
		Total:  rows,
		Limit:  limit,
		Offset: offset,
	}, nil
}

// func (r *ProductsPostgres) GetById(id int) (*domain.Product, error) {
// 	query := fmt.Sprintf(`
//   	SELECT *
//   	FROM %s
//   	WHERE id = $1`, products)

// 	var product domain.Product
// 	if err := r.db.Get(&product, query, id); err != nil {
// 		return nil, fmt.Errorf("select product: %v", err)
// 	}

// 	return &product, nil
// }

// func (r *ProductsPostgres) GetBySubcategory(id int) (*[]domain.Product, error) {
// 	query := fmt.Sprintf(`
// 	SELECT *
// 	FROM %s
// 	WHERE subcategory_id = $1`, products)

// 	var products []domain.Product
// 	if err := r.db.Select(&products, query, id); err != nil {
// 		return nil, fmt.Errorf("error select product: %v", err)
// 	}

// 	return &products, nil
// }

// func (r *ProductsPostgres) Update(product *domain.Product) (*domain.Product, error) {

// }
