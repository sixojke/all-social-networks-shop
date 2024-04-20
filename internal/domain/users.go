package domain

import "time"

type User struct {
	Id           int       `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	Password     string    `db:"password"`
	Email        string    `json:"email" db:"email"`
	Balance      float64   `json:"balance" db:"balance"`
	Role         string    `json:"role" db:"role"`
	LastVisitAt  time.Time `json:"last_visit_at" db:"last_visit_at"`
	RegisteredAt time.Time `json:"registered_at" db:"registered_at"`
}

type Session struct {
	RefreshToken string    `json:"refresh_token" db:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at" db:"expires_at"`
	UserId       int       `json:"user_id" db:"user_id"`
	UserRole     string    `json:"user_role" db:"role"`
}

type Supplier struct {
	Id                      int     `json:"id" db:"id"`
	TotalProfit             float64 `json:"total_profit" db:"total_profit"`
	Rating                  float64 `json:"rating" db:"rating"`
	TotalReviews            int     `json:"total_reviews" db:"total_reviews"`
	TotalOrders             int     `json:"total_orders" db:"total_orders"`
	AvgOrderFulfillmentTime int     `json:"avg_order_fulfillment_time" db:"AvgOrderFulfillmentTime"`
	UserId                  int     `json:"user_id" db:"user_id"`
}

type Buyer struct {
	Id            int     `json:"id" db:"id"`
	TotalExpenses float64 `json:"total_expenses" db:"total_expenses"`
	TotalOrders   int     `json:"total_orders" db:"total_orders"`
	UserId        int     `json:"user_id" db:"user_id"`
}

type Reviews struct {
	Id         int       `json:"id" db:"id"`
	Text       string    `json:"text" db:"text"`
	Rating     float64   `json:"rating" db:"rating"`
	CreatedAt  time.Time `json:"created_at"`
	BuyerId    int       `json:"buyer_id" db:"buyer_id"`
	SupplierId int       `json:"supplier_id" db:"supplier_id"`
}
