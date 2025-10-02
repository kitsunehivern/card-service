package handler

import (
	cardpb "card-service/gen/proto"
	"card-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CardHandler struct {
	cardSvc *service.CardService
}

func NewCardHandler(cardSvc *service.CardService) *CardHandler {
	return &CardHandler{cardSvc: cardSvc}
}

type RequestCardBody struct {
	UserID string `json:"user_id" binding:"required"`
}

func (cardHdl *CardHandler) RequestCard(c *gin.Context) {
	var body RequestCardBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := &cardpb.RequestCardRequest{UserId: body.UserID}

	resp, err := cardHdl.cardSvc.RequestCard(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.GetCard())
}

func (cardHdl *CardHandler) ActivateCard(c *gin.Context) {
	id := c.Param("id")
	req := &cardpb.ActivateCardRequest{Id: id}

	resp, err := cardHdl.cardSvc.ActivateCard(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.GetCard())
}

func (cardHdl *CardHandler) BlockCard(c *gin.Context) {
	id := c.Param("id")
	req := &cardpb.BlockCardRequest{Id: id}

	resp, err := cardHdl.cardSvc.BlockCard(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.GetCard())
}

func (cardHdl *CardHandler) UnblockCard(c *gin.Context) {
	id := c.Param("id")
	req := &cardpb.UnblockCardRequest{Id: id}

	resp, err := cardHdl.cardSvc.UnblockCard(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.GetCard())
}

func (cardHdl *CardHandler) CloseCard(c *gin.Context) {
	id := c.Param("id")
	req := &cardpb.CloseCardRequest{Id: id}

	resp, err := cardHdl.cardSvc.CloseCard(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.GetCard())
}

func (cardHdl *CardHandler) GetCard(c *gin.Context) {
	id := c.Param("id")
	req := &cardpb.GetCardRequest{Id: id}

	resp, err := cardHdl.cardSvc.GetCard(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.GetCard())
}
