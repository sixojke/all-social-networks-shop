package service

import (
	"github.com/sixojke/internal/config"
	"github.com/sixojke/internal/domain"
	"github.com/sixojke/internal/repository"
	"github.com/sixojke/pkg/auth"
	email "github.com/sixojke/pkg/email/smpt"
	"github.com/sixojke/pkg/hash"
	"github.com/sixojke/pkg/otp"
	"github.com/sixojke/pkg/payments/payok"
)

type UserSignUnInp struct {
	Username string
	Password string
	Email    string
}

type UserSignInInp struct {
	Username string
	Password string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Users interface {
	SignUp(inp UserSignUnInp) (id int, err error)
	SignIn(inp UserSignInInp) (Tokens, error)
	RefreshTokens(refreshToken string) (Tokens, error)
	Verify(userId int, code string) error
	GetById(id int) (*domain.User, error)
	Ban(id int, banStatus bool) error
}

type Telegram interface {
	CreateAuthLink(userId int) (string, error)
	Bind(telegramId int, code string) (userId int, err error)
	Unbind(userId int) error
}

type Category interface {
	CreateCategory(category *domain.Category) (id int, err error)
	UpdateCategory(category *domain.Category) error
	DeleteCategory(id int) error
	GetCategories() (*[]domain.Category, error)
	CreateSubcategory(subcategory *domain.Subcategory) (id int, err error)
	UpdateSubcategory(subcategory *domain.Subcategory) error
	DeleteSubcategory(id int) error
	GetSubcategories(categoryId int) (*[]domain.Subcategory, error)
}

type ReferralSystem interface {
	CreateCode(description string) (link string, err error)
	AddVisitor(referralCode string) error
	GetStats(limit, offset int) (*domain.Pagination, error)
	DeleteCode(referralCode string) error
}

type Log interface {
	WriteAdminLog(log *domain.Log) error
	GetAdminLogs(limit, offset int) (*domain.Pagination, error)
}

type Products interface {
	// Create(product *domain.Product) (int, error)
	GetAll(filters *domain.ProductFilters) (*domain.Pagination, error)
	// GetById(id int) (*domain.Product, error)
	// GetBySubcategoryId(id int) (*[]domain.Product, error)
	// Update(product *domain.Product) (*domain.Product, error)
}

type Cart interface {
	GetById(userId int) (*[]domain.CartGetByIdOut, error)
	SetQuantity(inp *domain.CartSetQuantityInp) error
}

type Deps struct {
	Repo         *repository.Repository
	Config       *config.Service
	Hasher       hash.PasswordHasher
	OtpGenerator otp.Generator
	EmailSender  *email.SMTPSender
	TokenManager auth.TokenManager
	PayokClient  payok.Service
}

type Service struct {
	Users          Users
	Telegram       Telegram
	Category       Category
	Products       Products
	Cart           Cart
	ReferralSystem ReferralSystem
	Log            Log
}

func NewService(deps *Deps) *Service {
	emailService := NewEmailService(deps.EmailSender)

	return &Service{
		Users:          NewUsersService(deps.Repo.Users, deps.Config.Users, deps.TokenManager, deps.Hasher, deps.OtpGenerator, emailService),
		Telegram:       NewBindSerivce(deps.Repo.Telegram, deps.Config.Telegram, deps.OtpGenerator),
		Category:       NewCategoryService(deps.Repo.Category),
		Products:       NewProductsService(deps.Repo.Products),
		Cart:           NewCartService(deps.Repo.Cart),
		ReferralSystem: NewReferralSystemService(deps.Repo.ReferralSystem, deps.Config.ReferralSystem, deps.OtpGenerator),
		Log:            NewLogService(deps.Repo.Log),
	}
}
