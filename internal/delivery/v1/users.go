package v1

import (
	"errors"
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

		authenticated := users.Group("/", h.userIdentity)
		{
			authenticated.GET("", h.userById)
		}
	}
}

type userSignUpInp struct {
	Username string `json:"username" binding:"required,min=8,max=32"`
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

type UserVerifyInp struct {
	Id   int    `binding:"required"`
	Code string `binding:"required"`
}

// @Summary User Verify Registration
// @Tags users-auth
// @Description user verify registration
// @ModuleID userVerify
// @Accept  json
// @Produce  json
// @Param input body UserVerifyInp true "user verify"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/verify/ [post]
func (h *Handler) userVerify(c *gin.Context) {
	var inp UserVerifyInp
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

// @Summary User get by refresh token
// @Security UsersAuth
// @Tags users
// @Description user get by refresh token
// @ModuleID userGetByRefresh
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.User
// @Failure 400,401,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/ [get]
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
