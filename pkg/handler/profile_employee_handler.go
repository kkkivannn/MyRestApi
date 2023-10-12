package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileEmployeeInput struct {
	Id int `json:"id"`
}

func (h *Handler) employeeProfile(c *gin.Context) {
	var input ProfileEmployeeInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	employee, err := h.service.GetProfileEmployee(input.Id)
	if err != nil {
		NewErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, employee)
}
