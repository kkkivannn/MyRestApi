package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) Identity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewErrorMessage(c, http.StatusUnauthorized, "Вы не авторизованы")
		return
	}
	splitsHeader := strings.Split(header, " ")
	if len(splitsHeader) != 2 {
		NewErrorMessage(c, http.StatusUnauthorized, "Не валидный токен")
		return
	}
	userId, err := h.service.ParseToken(splitsHeader[1])
	if err != nil {
		NewErrorMessage(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
