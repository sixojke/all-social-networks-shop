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

	userIdCtx   = "userId"
	userRoleCtx = "userRole"
)

func (h Handler) adminIdentity(c *gin.Context) {
	h.userIdentity(c)

	role, ok := c.Get(userRoleCtx)
	if !ok {
		newResponse(c, http.StatusUnauthorized, "user unauthorized")

		return
	}

	if role != "admin" {
		newResponse(c, http.StatusForbidden, "forbidden")

		return
	}
}

func (h *Handler) userIdentity(c *gin.Context) {
	sub, err := h.parseAuthHeader(c)
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	user := strings.Split(sub, "/")

	c.Set(userIdCtx, user[0])
	c.Set(userRoleCtx, user[1])
}

func (h *Handler) parseAuthHeader(c *gin.Context) (sub string, err error) {
	header := c.GetHeader(authHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, "=")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts) == 0 {
		return "", errors.New("token is empty")
	}

	return h.tokenManager.Parse(headerParts[1])
}

func getUserId(c *gin.Context) (int, error) {
	return getIdByContext(c, userIdCtx)
}

func getUserRole(c *gin.Context) (string, error) {
	return getRoleByConetxt(c, userRoleCtx)
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

func getRoleByConetxt(c *gin.Context, context string) (string, error) {
	roleFromCtx, ok := c.Get(context)
	if !ok {
		return "", errors.New("context not found")
	}

	roleStr, ok := roleFromCtx.(string)
	if !ok {
		return "", errors.New("invalid type")
	}

	return roleStr, nil
}
