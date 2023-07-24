package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"transactions/models"
)

type Handler struct {
	Engine   *gin.Engine
	services ServiceInterface
}

type ServiceInterface interface {
	MakeTransaction(senderPhone string, receiverPhone string, amount int) error
}

func NewHandler(services ServiceInterface, engine *gin.Engine) *Handler {
	return &Handler{
		Engine:   engine,
		services: services,
	}
}

func (h Handler) AllRoutes() {
	h.Engine.POST("/make_transactions", h.MakeTransaction)
}

func (h Handler) MakeTransaction(c *gin.Context) {

	var transactions models.Transaction
	if err := c.ShouldBindJSON(&transactions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data type"})
	}
	if err := h.services.MakeTransaction(transactions.SenderPhone, transactions.ReceiverPhone, transactions.Amount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid in updating user data"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User data successfully given"})
}
