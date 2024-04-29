package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (r *ProductsPostgres) GetAll(filters *domain.ProductFilters) (*domain.Pagination, error) {
	query := fmt.Sprintf(`
	SELECT 
		products.id, products.name, products.description, products.price, products.quantity, 
		products.quantity_sales, products.category_id, products.uploaded_at, category.img_path
	FROM %s
	INNER JOIN %s ON products.category_id = category.id `, products, category)

	where := make([]string, 0)
	whereArgs := make([]interface{}, 0)
	args := make([]interface{}, 0)
	argsId := 1

	if filters.CategoryId != 0 {
		where = append(where, fmt.Sprintf("category_id=$%v ", argsId))
		whereArgs = append(whereArgs, filters.CategoryId)
		args = append(args, filters.CategoryId)
		argsId++
	}

	if filters.SubcategoryId != 0 {
		where = append(where, fmt.Sprintf("subcategory_id=$%v ", argsId))
		whereArgs = append(whereArgs, filters.SubcategoryId)
		args = append(args, filters.SubcategoryId)
		argsId++
	}

	if filters.IsAvailable == 1 {
		where = append(where, fmt.Sprintf("quantity > $%v ", argsId))
		whereArgs = append(whereArgs, 0)
		args = append(args, 0)
		argsId++
	}

	if len(where) != 0 {
		query += "WHERE " + strings.Join(where, "AND ")
	}

	if filters.SortPrice == "asc" {
		query += "ORDER BY products.price ASC "
	} else if filters.SortPrice == "desc" {
		query += "ORDER BY products.price DESC "
	} else if filters.SortDefect == "asc" {
		query += "ORDER BY products.price ASC "
	} else if filters.SortDefect == "desc" {
		query += "ORDER BY products.price DESC "
	} else {
		query += "ORDER BY products.id "
	}

	pagination := make([]string, 0)

	pagination = append(pagination, fmt.Sprintf("LIMIT $%v", argsId))
	args = append(args, filters.Limit)
	argsId++

	pagination = append(pagination, fmt.Sprintf("OFFSET $%v", argsId))
	args = append(args, filters.Offset)
	argsId++

	query += strings.Join(pagination, " ")

	var p []domain.Product
	if err := r.db.Select(&p, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("get all products: %v", err)
	}
	logrus.Info(query, args)

	query = fmt.Sprintf(`
	SELECT COUNT(*)
	FROM %s `, products)

	if len(where) != 0 {
		query += "WHERE " + strings.Join(where, "AND ")
	}

	var totalItems int
	if err := r.db.QueryRow(query, whereArgs...).Scan(&totalItems); err != nil {
		return nil, fmt.Errorf("error select count rows: %v", err)
	}

	return &domain.Pagination{
		Data:       p,
		TotalItems: totalItems,
		Limit:      filters.Limit,
		TotalPages: getPages(totalItems, filters.Limit),
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
