package handler

import (
	cardpb "card-service/gen/proto"
	"card-service/internal/errmsg"
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

type ActivateCardBody struct {
	ID string `json:"id" binding:"required"`
}

type BlockCardBody struct {
	ID string `json:"id" binding:"required"`
}

type UnblockCardBody struct {
	ID string `json:"id" binding:"required"`
}

type RetireCardBody struct {
	ID string `json:"id" binding:"required"`
}

type CloseCardBody struct {
	ID string `json:"id" binding:"required"`
}

func (ch *CardHandler) RequestCard(c *gin.Context) {
	var body RequestCardBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errmsg.CardMissingFieldInBody})
		return
	}

	req := &cardpb.RequestCardRequest{UserId: body.UserID}

	resp, err := ch.cardSvc.RequestCard(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.GetCard())
}

func (ch *CardHandler) ActivateCard(c *gin.Context) {
	var body ActivateCardBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errmsg.CardMissingFieldInBody})
		return
	}

	req := &cardpb.ActivateCardRequest{Id: body.ID}

	resp, err := ch.cardSvc.ActivateCard(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.GetCard())
}

func (ch *CardHandler) BlockCard(c *gin.Context) {
	var body BlockCardBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errmsg.CardMissingFieldInBody})
		return
	}

	req := &cardpb.BlockCardRequest{Id: body.ID}

	resp, err := ch.cardSvc.BlockCard(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.GetCard())
}

func (ch *CardHandler) UnblockCard(c *gin.Context) {
	var body UnblockCardBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errmsg.CardMissingFieldInBody})
		return
	}

	req := &cardpb.UnblockCardRequest{Id: body.ID}

	resp, err := ch.cardSvc.UnblockCard(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.GetCard())
}

func (ch *CardHandler) RetireCard(c *gin.Context) {
	var body RetireCardBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errmsg.CardMissingFieldInBody})
		return
	}

	req := &cardpb.RetireCardRequest{Id: body.ID}

	resp, err := ch.cardSvc.RetireCard(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errmsg.CardMissingFieldInBody})
		return
	}

	c.JSON(http.StatusOK, resp.GetCard())
}

func (ch *CardHandler) CloseCard(c *gin.Context) {
	var body CloseCardBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errmsg.CardMissingFieldInBody})
		return
	}

	req := &cardpb.CloseCardRequest{Id: body.ID}

	resp, err := ch.cardSvc.CloseCard(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.GetCard())
}

func (ch *CardHandler) GetCard(c *gin.Context) {
	id := c.Param("id")
	req := &cardpb.GetCardRequest{Id: id}

	resp, err := ch.cardSvc.GetCard(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp.GetCard())
}
