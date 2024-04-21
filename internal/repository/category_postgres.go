package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sixojke/internal/domain"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) CreateCategory(cat *domain.Category) (id int, err error) {
	query := fmt.Sprintf(`
	INSERT INTO %s
		(name, img_path)
	VALUES
		($1, $2)`, category)

	if err := r.db.QueryRow(query, cat.Name, cat.ImgPath).Scan(&id); err != nil {
		return 0, fmt.Errorf("insert category: %v", err)
	}

	return
}

func (r *CategoryPostgres) CategoryEdit(cat *domain.Category) error {
	query := fmt.Sprintf(`
	UPDATE %s
	SET name = $1`, category)

	argsId := 2
	args := make([]interface{}, 0)
	args = append(args, cat.Name)

	if cat.ImgPath != "" {
		query += fmt.Sprintf(", img_path = $%v", argsId)
		args = append(args, cat.ImgPath)
		argsId++
	}

	query += fmt.Sprintf(" WHERE id = $%v", argsId)
	args = append(args, cat.Id)

	if _, err := r.db.Exec(query, args...); err != nil {
		return fmt.Errorf("update category: %v", err)
	}

	return nil
}

func (r *CategoryPostgres) GetCategories() (*[]domain.Category, error) {
	query := fmt.Sprintf(`
	SELECT *
	FROM %s`, category)

	var categories []domain.Category
	if err := r.db.Select(&categories, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error select categoryies: %v", err)
	}

	return &categories, nil
}

func (r *CategoryPostgres) GetSubcategories(categoryId int) (*[]domain.Subcategory, error) {
	query := fmt.Sprintf(`
	SELECT *
	FROM %s
	WHERE category_id = $1`, subcategory)

	var subcategories []domain.Subcategory
	if err := r.db.Select(&subcategories, query, categoryId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("error select subcategoryies: %v", err)
	}

	return &subcategories, nil
}