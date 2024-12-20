package domain

import "time"

type User struct {
	Id           int       `json:"id,omitempty" db:"id"`
	Username     string    `json:"username,omitempty" db:"username"`
	Password     string    `json:"password,omitempty" db:"password"`
	Email        string    `json:"email,omitempty" db:"email"`
	Balance      float64   `json:"balance,omitempty" db:"balance"`
	Role         string    `json:"role,omitempty" db:"role"`
	LastVisitAt  time.Time `json:"last_visit_at,omitempty" db:"last_visit_at"`
	RegisteredAt time.Time `json:"registered_at,omitempty" db:"registered_at"`
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

type UserChangePasswordInp struct {
	UserId      int `db:"user_id"`
	OldPassword string
	NewPassword string
}

type UserCreatePasswordRecoveryInp struct {
	UserId       int       `db:"user_id"`
	SecretCode   string    `db:"secret_code"`
	RecoveryTime time.Time `db:"recovery_time"`
}
