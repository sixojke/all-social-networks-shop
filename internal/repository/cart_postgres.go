package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sixojke/internal/domain"
)

type CartPostgres struct {
	db *sqlx.DB
}

func NewCartPostgres(db *sqlx.DB) *CartPostgres {
	return &CartPostgres{
		db: db,
	}
}

func (r *CartPostgres) GetById(userId int) (*[]domain.CartGetByIdOut, error) {
	query := fmt.Sprintf(`
	SELECT 
		cart.product_id, cart.quantity,
		products.name, products.price,
		category.img_path,
		users.username
	FROM %s
	INNER JOIN %s ON products.id = cart.product_id
	INNER JOIN %s ON users.id = products.supplier_id
	INNER JOIN %s ON category.id = products.category_id
	WHERE cart.user_id = $1`,
		cart, products, users, category)

	var cart []domain.CartGetByIdOut
	if err := r.db.Select(&cart, query, userId); err != nil {
		return nil, fmt.Errorf("select cart: %v", err)
	}

	return &cart, nil
}

func (r *CartPostgres) SetQuantity(inp *domain.CartSetQuantityInp) error {
	query := fmt.Sprintf(`
	UPDATE %s
	SET quantity = $1
	WHERE user_id = $2 AND product_id = $3`, cart)

	if _, err := r.db.Exec(query, inp.Quantity, inp.UserId, inp.ProductId); err != nil {
		return fmt.Errorf("update quantity: %v", err)
	}

	return nil
}
