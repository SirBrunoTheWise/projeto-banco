package api

import (
	"database/sql"
	"net/http"

	db "github.com/SirBrunoTheWise/hunt/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createCardRequest struct {
	CardType        int16       `json:"card_type"`
	CardNumber      int64       `json:"card_number"`
	CardProgression interface{} `json:"card_progression"`
	CardImage       []byte      `json:"card_image"`
}

func (server *Server) createCard(ctx *gin.Context) {
	var req createCardRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCardParams{
		CardType:        req.CardType,
		CardNumber:      req.CardNumber,
		CardProgression: req.CardProgression,
		CardImage:       req.CardImage,
	}

	card, err := server.store.CreateCard(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, card)
}

type getCardRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getCard(ctx *gin.Context) {
	var req getCardRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	card, err := server.store.GetCard(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, card)
}

type listCardsRequest struct {
	Limit  int32 `form:"limit,default=10" binding:"required,min=1,max=100"`
	Offset int32 `form:"offset,default=0" binding:"required,min=0"`
}

func (server *Server) listCards(ctx *gin.Context) {
	var req listCardsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCardsParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	cards, err := server.store.ListCards(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, cards)
}

type updateCardRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateCardRequestBody struct {
	CardType        int16       `json:"card_type"`
	CardNumber      int64       `json:"card_number"`
	CardProgression interface{} `json:"card_progression"`
	CardImage       []byte      `json:"card_image"`
}

func (server *Server) updateCard(ctx *gin.Context) {
	var req updateCardRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqBody updateCardRequestBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCardParams{
		CardID:          req.ID,
		CardType:        reqBody.CardType,
		CardNumber:      reqBody.CardNumber,
		CardProgression: reqBody.CardProgression,
		CardImage:       reqBody.CardImage,
	}

	card, err := server.store.UpdateCard(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, card)
}

func (server *Server) deleteCard(ctx *gin.Context) {
	var req getCardRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteCard(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Card deleted"})
}
