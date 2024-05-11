package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sixojke/internal/domain"
	"github.com/sixojke/internal/service"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-up", h.userSignUp)
		users.POST("/sign-in", h.userSignIn)
		users.POST("/auth/refresh", h.userRefresh)
		users.POST("/verify", h.userVerify)
		users.POST("/forgot-password", h.userForgotPassword)
		users.POST("/password-recovery", h.userPasswordRecovery)
	}

	user := api.Group("/user", h.userIdentity)
	{
		user.GET("", h.userById)

		h.initCartRoutes(user)

		telegram := user.Group("/telegram")
		{
			telegram.POST("/bind", h.userBindTelegram)
			telegram.POST("/unbind", h.userUnbindTelegram)
		}

		supplier := user.Group("/supplier")
		{
			_ = supplier
		}

		buyer := user.Group("/buyer")
		{
			_ = buyer
		}

		twoFa := user.Group("/2fa")
		{
			twoFa.GET("/authenticator", h.twoFaCheckPin)
			twoFa.POST("/authenticator", h.twoFaCreatePairingLink)
		}

		security := user.Group("/security")
		{
			security.PUT("/password", h.securityChangePassword)
		}
	}
}

type userSignUpInp struct {
	Username string `json:"username" binding:"required,min=6,max=32"`
	Password string `json:"password" binding:"required,min=8,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
}

// @Summary User SignUp
// @Tags users-auth
// @Description create user account
// @ModuleID userSignUp
// @Accept  json
// @Produce  json
// @Param input body userSignUpInp true "sign up info"
// @Success 201 {string} string "ok"
// @Failure 400,404,422 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/sign-up [post]
func (h *Handler) userSignUp(c *gin.Context) {
	var inp userSignUpInp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	id, err := h.services.Users.SignUp(service.UserSignUnInp{
		Username: inp.Username,
		Password: inp.Password,
		Email:    inp.Email,
	})
	if err != nil {
		if err == domain.ErrDuplicateKey {
			newResponse(c, http.StatusUnprocessableEntity, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusCreated, idResponse{ID: id})
}

type userSignInInp struct {
	Username string `json:"username" binding:"required,min=6,max=32"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

// @Summary User SignIn
// @Tags users-auth
// @Description user sign in
// @ModuleID userSignIn
// @Accept  json
// @Produce  json
// @Param input body userSignInInp true "sign up info"
// @Success 200 {object} tokenResponse
// @Failure 400,403,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/sign-in [post]
func (h *Handler) userSignIn(c *gin.Context) {
	var inp userSignInInp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	tokens, err := h.services.Users.SignIn(service.UserSignInInp{
		Username: inp.Username,
		Password: inp.Password,
	})
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		if errors.Is(err, domain.ErrUserNotVerified) {
			newResponse(c, http.StatusForbidden, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

type refreshTokenInp struct {
	Token string `json:"refresh_token" binding:"required"`
}

// @Summary User Refresh Tokens
// @Tags users-auth
// @Description user refresh tokens
// @Accept  json
// @Produce  json
// @Param input body refreshTokenInp true "sign up info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/auth/refresh [post]
func (h *Handler) userRefresh(c *gin.Context) {
	var inp refreshTokenInp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	tokens, err := h.services.Users.RefreshTokens(inp.Token)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			newResponse(c, http.StatusUnauthorized, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

type userVerifyInp struct {
	Id   int    `binding:"required"`
	Code string `binding:"required"`
}

// @Summary User Verify Registration
// @Tags users-auth
// @Description user verify registration
// @ModuleID userVerify
// @Accept  json
// @Produce  json
// @Param input body userVerifyInp true "user verify"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/verify/ [post]
func (h *Handler) userVerify(c *gin.Context) {
	var inp userVerifyInp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.services.Users.Verify(inp.Id, inp.Code); err != nil {
		if errors.Is(err, domain.ErrVerificationCodeInvalid) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{"success"})
}

type userForgotPasswordInp struct {
	UsernameOrEmail string `json:"username_or_email" binding:"required"`
}

// @Summary User Forgot Password
// @Tags users-auth
// @Description sends an email with a recovery link if the user is found
// @ModuleID userForgotPassword
// @Accept  json
// @Produce  json
// @Param input body userForgotPasswordInp true "username or email"
// @Success 200 {object} idResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/forgot-password [post]
func (h *Handler) userForgotPassword(c *gin.Context) {
	var inp userForgotPasswordInp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	userId, err := h.services.Users.ForgotPassword(inp.UsernameOrEmail)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, idResponse{ID: userId})
}

type userPasswordRecoveryInp struct {
	SecretCode  string `json:"secret_code" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

// @Summary User Password Recovery
// @Tags users-auth
// @Description password recovery
// @ModuleID userPasswordRecovery
// @Accept  json
// @Produce  json
// @Param input body userPasswordRecoveryInp true "password recovery"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/password-recovery [post]
func (h *Handler) userPasswordRecovery(c *gin.Context) {
	var inp userPasswordRecoveryInp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.services.Users.PasswordRecovery(inp.SecretCode, inp.NewPassword); err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{Message: "success"})
}

// @Summary User get by refresh token
// @Security UsersAuth
// @Tags user
// @Description user get by refresh token
// @ModuleID userGetByRefresh
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.User
// @Failure 400,401,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /user [get]
func (h *Handler) userById(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	user, err := h.services.Users.GetById(id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			newResponse(c, http.StatusUnauthorized, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary User Bind Telegram
// @Security UsersAuth
// @Tags user
// @Description get a link for bind telegram account
// @ModuleID userBindTelegram
// @Accept  json
// @Produce  json
// @Success 200 {object} linkResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /user/telegram/bind [post]
func (h *Handler) userBindTelegram(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	link, err := h.services.Telegram.CreateAuthLink(userId)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, linkResponse{Link: link})
}

// @Summary User Unbind Telegram
// @Security UsersAuth
// @Tags user
// @Description unbind telegram account
// @ModuleID userUnbindTelegram
// @Accept  json
// @Produce  json
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /user/telegram/unbind [post]
func (h *Handler) userUnbindTelegram(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	if err := h.services.Telegram.Unbind(userId); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{Message: "success"})
}

// @Summary 2fa Create Pairing Link
// @Security UsersAuth
// @Tags user
// @Description creates a link for pairing, the link contains a qr-code image
// @ModuleID twoFaCreatePairingLink
// @Accept  json
// @Produce  json
// @Success 200 {object} linkResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /user/2fa/authenticator [post]
func (h *Handler) twoFaCreatePairingLink(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	link, err := h.services.TwoFa.CreatePairingLink(userId)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, linkResponse{Link: link})
}

// @Summary 2fa Check Pin
// @Security UsersAuth
// @Tags user
// @Description checks the PIN code from two-step authentication
// @ModuleID twoFaCheckPin
// @Accept  json
// @Produce  json
// @Param pin query string false "6-digit pin code"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /user/2fa/authenticator [get]
func (h *Handler) twoFaCheckPin(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	pin, err := processIntParam(c.Query("pin"))
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	check, err := h.services.TwoFa.СheckTwoFactorPin(userId, pin)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidPin) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{Message: fmt.Sprintf("%v", check)})
}

type securityChangePasswordInp struct {
	OldPassword string `json:"old_password" binding:"required,min=8"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

// @Summary User Change Password
// @Security UsersAuth
// @Tags user
// @Description user change password
// @ModuleID securityChangePassword
// @Accept  json
// @Produce  json
// @Param input body securityChangePasswordInp true "change password"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /user/security/password [put]
func (h *Handler) securityChangePassword(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	var inp securityChangePasswordInp
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.services.Users.ChangePassword(&domain.UserChangePasswordInp{
		UserId:      userId,
		OldPassword: inp.OldPassword,
		NewPassword: inp.NewPassword,
	}); err != nil {
		if errors.Is(err, domain.ErrInvalidPassword) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{Message: "success"})
}
