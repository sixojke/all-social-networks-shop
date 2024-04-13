package service

import (
	"github.com/sixojke/pkg/hash"
	"github.com/sixojke/pkg/otp"
)

type UsersService struct {
	hasher       hash.PasswordHasher
	otpGenerator otp.Generator
}

func NewUsersService(hasher hash.PasswordHasher) *UsersService {
	return &UsersService{
		hasher: hasher,
	}
}

// func (s *UsersService) SignUp(ctx context.Context, input UserSignUpInput) error {
// 	passwordHash, err := s.hasher.Hash(input.Password)
// 	if err != nil {
// 		return err
// 	}

// 	verificationCode := s.otpGenerator.RandomSecret(s.verificationCodeLength)

// 	user := domain.User{
// 		Name:         input.Name,
// 		Password:     passwordHash,
// 		Phone:        input.Phone,
// 		Email:        input.Email,
// 		RegisteredAt: time.Now(),
// 		LastVisitAt:  time.Now(),
// 		Verification: domain.Verification{
// 			Code: verificationCode,
// 		},
// 	}

// 	if err := s.repo.Create(ctx, user); err != nil {
// 		if errors.Is(err, domain.ErrUserAlreadyExists) {
// 			return err
// 		}

// 		return err
// 	}

// 	return s.emailService.SendUserVerificationEmail(VerificationEmailInput{
// 		Email:            user.Email,
// 		Name:             user.Name,
// 		VerificationCode: verificationCode,
// 	})
// }

// func (s *UsersService) SignIn(ctx context.Context, input UserSignInInput) (Tokens, error) {
// 	passwordHash, err := s.hasher.Hash(input.Password)
// 	if err != nil {
// 		return Tokens{}, err
// 	}

// 	user, err := s.repo.GetByCredentials(ctx, input.Email, passwordHash)
// 	if err != nil {
// 		if errors.Is(err, domain.ErrUserNotFound) {
// 			return Tokens{}, err
// 		}

// 		return Tokens{}, err
// 	}

// 	return s.createSession(ctx, user.ID)
// }

// func (s *UsersService) RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error) {
// 	student, err := s.repo.GetByRefreshToken(ctx, refreshToken)
// 	if err != nil {
// 		return Tokens{}, err
// 	}

// 	return s.createSession(ctx, student.ID)
// }

// func (s *UsersService) Verify(ctx context.Context, userID primitive.ObjectID, hash string) error {
// 	err := s.repo.Verify(ctx, userID, hash)
// 	if err != nil {
// 		if errors.Is(err, domain.ErrVerificationCodeInvalid) {
// 			return err
// 		}

// 		return err
// 	}

// 	return nil
// }
