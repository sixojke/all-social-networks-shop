package repository

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/sixojke/internal/domain"
)

type Users interface {
	Create(user *domain.User, code string) (int, error)
	GetByCredentials(username, password string) (*domain.User, error)
	GetByRefreshToken(refreshToken string) (*domain.Session, error)
	Verify(userId int, code string) error
	GetUserByUsernameOrEmail(usernameOrEmail string) (*domain.User, error)
	SetSession(session *domain.Session) error
	GetById(id int) (*domain.User, error)
	Ban(id int, banStatus bool) error
	ChangePassword(inp *domain.UserChangePasswordInp) error
	CreatePasswordRecovery(inp *domain.UserCreatePasswordRecoveryInp) error
	PasswordRecovery(secretCode string, newPassword string) error
}

type Telegram interface {
	CreateAuthLink(code string, userId int) (string, error)
	Bind(telegramId int, code string) (userId int, err error)
	Unbind(userId int) error
}

type Category interface {
	CreateCategory(cat *domain.Category) (id int, err error)
	UpdateCategory(cat *domain.Category) error
	DeleteCategory(id int) error
	GetCategories() (*[]domain.Category, error)
	CreateSubcategory(subcat *domain.Subcategory) (id int, err error)
	UpdateSubcategory(subcat *domain.Subcategory) error
	DeleteSubcategory(id int) error
	GetSubcategories(categoryId int) (*[]domain.Subcategory, error)
}

type ReferralSystem interface {
	CreateCode(ref domain.ReferralSystem) error
	AddVisitor(referralCode string) error
	GetStats(limit, offset int) (*domain.Pagination, error)
	DeleteCode(referralCode string) error
}

type Log interface {
	WriteAdminLog(log *domain.Log) error
	GetAdminLogs(limit int, offset int) (*domain.Pagination, error)
}

type Products interface {
	// Create(product *domain.Product) (int, error)
	GetAll(filters *domain.ProductFilters) (*domain.Pagination, error)
	// GetById(id int) (*domain.Product, error)
	// GetBySubcategory(id int) (*[]domain.Product, error)
	// Update(product *domain.Product) (*domain.Product, error)
}

type Cart interface {
	GetById(userId int) (*[]domain.CartGetByIdOut, error)
	SetQuantity(inp *domain.CartSetQuantityInp) error
}

type TwoFa interface {
	CreatePairingLink(userId int, secret_code string) error
	GetSecretCode(userId int) (string, error)
}

type Deps struct {
	Postgres *sqlx.DB
	Redis    *redis.Client
}

type Repository struct {
	Users          Users
	Telegram       Telegram
	Category       Category
	Products       Products
	Cart           Cart
	ReferralSystem ReferralSystem
	Log            Log
	TwoFa          TwoFa
}

func NewRepository(deps *Deps) *Repository {
	return &Repository{
		Users:          NewUsersPostgres(deps.Postgres),
		Telegram:       NewBindPostgres(deps.Postgres),
		Category:       NewCategoryPostgres(deps.Postgres),
		Products:       NewProductsPostgres(deps.Postgres),
		Cart:           NewCartPostgres(deps.Postgres),
		ReferralSystem: NewReferralLinksPostgres(deps.Postgres),
		Log:            NewLogPostgres(deps.Postgres),
		TwoFa:          NewTwoFaPostgres(deps.Postgres),
	}
}
