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

		authenticated := users.Group("/", h.userIdentity)
		{
			authenticated.POST("/verify/:code", h.userVerify)
		}
	}
}

type userSignUpInp struct {
	Username string `json:"username" binding:"required,min=8,max=32"`
	Password string `json:"password" binding:"required,min=8,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
}

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
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, idResponse{ID: id})
}

type userSignInInp struct {
	Username string `json:"username" binding:"required,min=8,max=32"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

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

func (h *Handler) userVerify(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		newResponse(c, http.StatusBadRequest, "code is empty")

		return
	}

	id, err := getUserId(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	if err := h.services.Users.Verify(id, code); err != nil {
		if errors.Is(err, domain.ErrVerificationCodeInvalid) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, response{"success"})
}
