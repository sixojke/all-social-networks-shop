package domain

import "time"

type Product struct {
	Id            int       `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	Price         float64   `json:"price" db:"price"`
	Quantity      int       `json:"quantity" db:"quantity"`
	QuantitySales int       `json:"quantity_sales" db:"quantity_sales"`
	Description   string    `json:"description" db:"description"`
	Img           string    `json:"img" db:"img_path"`
	UploadedAt    time.Time `json:"uploaded_at,omitempty" db:"uploaded_at"`
	CategoryId    int       `json:"category_id" db:"category_id"`
	SubcategoryId int       `json:"subcategory_id,omitempty" db:"subcategory_id"`
}

type Category struct {
	Id      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	ImgPath string `json:"img_path" db:"img_path"`
}

type Subcategory struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	HoldTime   int    `json:"hold_time" db:"hold_time"`
	CategoryId int    `json:"category_id" db:"category_id"`
}
