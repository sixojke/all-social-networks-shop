package domain

type CartGetByIdOut struct {
	ProductId        int     `json:"product_id" db:"product_id"`
	Name             string  `json:"name" db:"name"`
	Price            float64 `json:"price" db:"price"`
	Quantity         int     `json:"quantity" db:"quantity"`
	SupplierUsername string  `json:"supplier_username" db:"username"`
	ImgPath          string  `json:"img_path" db:"img_path"`
}

type CartSetQuantityInp struct {
	UserId    int `db:"user_id"`
	ProductId int `db:"product_id"`
	Quantity  int `db:"quantity"`
}
