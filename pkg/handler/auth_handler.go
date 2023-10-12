package handler

import (
	api "RestaurantRestApi"
	"RestaurantRestApi/pkg/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthInput struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthEmployee struct {
}

func (h *Handler) signIn(c *gin.Context) {
	var inputData AuthInput
	if err := c.BindJSON(&inputData); err != nil {
		NewErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.GenerateTokens(inputData.Phone, inputData.Password)
	if err != nil {
		NewErrorMessage(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) singUp(c *gin.Context) {
	var inputData api.Client
	if err := c.BindJSON(&inputData); err != nil {
		NewErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.CreateUser(inputData)
	if err != nil {
		NewErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) refreshToken(c *gin.Context) {
	var inputData auth.Token
	if err := c.BindJSON(&inputData); err != nil {
		NewErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.RefreshToken(inputData.Token)
	if err != nil {
		NewErrorMessage(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}

func (h *Handler) signInEmployee(c *gin.Context) {
	var inputData api.Employee
	if err := c.BindJSON(&inputData); err != nil {
		NewErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.CreateEmployee(inputData)
	if err != nil {
		NewErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signUpEmployee(c *gin.Context) {
	var input api.Employee
	if err := c.BindJSON(&input); err != nil {
		NewErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.CreateEmployee(input)
	if err != nil {
		NewErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
