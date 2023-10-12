package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) profile(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}
	client, err := h.service.GetProfileClient(id)
	if err != nil {
		NewErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, client)
}
