package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sixojke/internal/config"
	"github.com/sixojke/internal/domain"
	"github.com/sixojke/internal/repository"
	"github.com/sixojke/pkg/auth"
	"github.com/sixojke/pkg/hash"
	"github.com/sixojke/pkg/otp"
)

type UsersService struct {
	repo         repository.Users
	config       config.UsersService
	tokenManager auth.TokenManager
	hasher       hash.PasswordHasher
	otpGenerator otp.Generator
	email        *emailSerive
}

func NewUsersService(repo repository.Users, config config.UsersService, tokenManager auth.TokenManager, hasher hash.PasswordHasher,
	otpGenerator otp.Generator, email *emailSerive) *UsersService {
	return &UsersService{
		repo:         repo,
		config:       config,
		tokenManager: tokenManager,
		hasher:       hasher,
		otpGenerator: otpGenerator,
		email:        email,
	}
}

func (s *UsersService) SignUp(inp UserSignUnInp) (id int, err error) {
	passwordHash, err := s.hasher.Hash(inp.Password)
	if err != nil {
		return 0, err
	}

	verificationCode, err := s.otpGenerator.RandomSecret()
	if err != nil {
		return 0, err
	}

	id, err = s.repo.Create(&domain.User{
		Username:    inp.Username,
		Password:    passwordHash,
		Email:       inp.Email,
		Balance:     0,
		LastVisitAt: time.Now(),
	}, verificationCode)
	if err != nil {
		return 0, err
	}

	if err := s.email.SendUserVerificationEmail(&VerificationEmailInp{
		Email:            inp.Email,
		VerificationCode: verificationCode,
	}); err != nil {
		return 0, err
	}

	return id, nil
}

func (s *UsersService) SignIn(inp UserSignInInp) (Tokens, error) {
	passwordHash, err := s.hasher.Hash(inp.Password)
	if err != nil {
		return Tokens{}, err
	}

	user, err := s.repo.GetByCredentials(inp.Username, passwordHash)
	if err != nil {
		return Tokens{}, err
	}

	return s.createSession(user.Id, user.Role)
}

func (s *UsersService) RefreshTokens(refreshToken string) (Tokens, error) {
	session, err := s.repo.GetByRefreshToken(refreshToken)
	if err != nil {
		return Tokens{}, err
	}

	return s.createSession(session.UserId, session.UserRole)
}

func (s *UsersService) Verify(userId int, code string) error {
	err := s.repo.Verify(userId, code)
	if err != nil {
		return err
	}

	return nil
}

func (s *UsersService) createSession(userId int, userRole string) (Tokens, error) {
	var (
		tokens Tokens
		err    error
	)

	tokens.AccessToken, err = s.tokenManager.NewJWT(strconv.Itoa(userId), userRole, s.config.Auth.JWT.AccessTokenTTL)
	if err != nil {
		return tokens, err
	}

	tokens.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return tokens, err
	}

	session := domain.Session{
		RefreshToken: tokens.RefreshToken,
		ExpiresAt:    time.Now().Add(s.config.Auth.JWT.RefreshTokenTTL),
		UserId:       userId,
	}

	if err := s.repo.SetSession(&session); err != nil {
		return Tokens{}, err
	}

	return tokens, nil
}

func (s *UsersService) GetById(id int) (*domain.User, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("error users service UserByRefreshToken: %v", err)
	}

	return user, nil
}
