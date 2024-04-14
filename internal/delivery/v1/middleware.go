package v1

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"

	userCtx = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	id, err := h.parseAuthHeader(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, id)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (id string, err error) {
	header := c.GetHeader(authHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts) == 0 {
		return "", errors.New("token is empty")
	}

	return h.tokenManager.Parse(headerParts[1])
}

func getUserId(c *gin.Context) (int, error) {
	return getIdByContext(c, userCtx)
}

func getIdByContext(c *gin.Context, context string) (int, error) {
	idFromCtx, ok := c.Get(context)
	if !ok {
		return 0, errors.New("context not found")
	}

	idStr, ok := idFromCtx.(string)
	if !ok {
		return 0, errors.New("invalid type")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("parse id")
	}

	return id, nil
}
