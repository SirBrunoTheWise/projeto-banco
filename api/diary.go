package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/SirBrunoTheWise/hunt/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createDiaryRequest struct {
	DateOf   time.Time `json:"date_of" binding:"required"`
	UserID   int64     `json:"user_id" binding:"required,min=1"`
	Exercise int64     `json:"exercise" binding:"required"`
	Meal     int64     `json:"meal" binding:"required"`
	Cards    int64     `json:"cards" binding:"required"`
}

func (server *Server) createDiary(ctx *gin.Context) {
	var req createDiaryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateDiaryEntryParams{
		DateOf:   req.DateOf,
		UserID:   req.UserID,
		Exercise: req.Exercise,
		Meal:     req.Meal,
		Cards:    req.Cards,
	}

	diary, err := server.store.CreateDiaryEntry(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, diary)
}

type getDiaryRequest struct {
	DateOf time.Time `uri:"date_of" binding:"required"`
	UserID int64     `uri:"user_id" binding:"required,min=1"`
}

func (server *Server) getDiary(ctx *gin.Context) {
	var req getDiaryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetDiaryEntryParams{
		DateOf: req.DateOf,
		UserID: req.UserID,
	}

	diary, err := server.store.GetDiaryEntry(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, diary)
}

type listDiariesRequest struct {
	UserID int64 `uri:"user_id" binding:"required,min=1"`
	Limit  int32 `form:"limit" binding:"required,min=1"`
	Offset int32 `form:"offset" binding:"required,min=0"`
}

func (server *Server) listDiaries(ctx *gin.Context) {
	var req listDiariesRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListDiaryEntriesParams{
		UserID: req.UserID,
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	diaries, err := server.store.ListDiaryEntries(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, diaries)
}

type listDiariesByDateRangeRequest struct {
	UserID    int64     `uri:"user_id" binding:"required,min=1"`
	StartDate time.Time `form:"start_date" binding:"required"`
	EndDate   time.Time `form:"end_date" binding:"required"`
}

func (server *Server) listDiariesByDateRange(ctx *gin.Context) {
	var req listDiariesByDateRangeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListDiaryEntriesByDateRangeParams{
		UserID:   req.UserID,
		DateOf:   req.StartDate,
		DateOf_2: req.EndDate,
	}

	diaries, err := server.store.ListDiaryEntriesByDateRange(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, diaries)
}

type updateDiaryUriParams struct {
	DateOf time.Time `uri:"date_of" binding:"required"`
	UserID int64     `uri:"user_id" binding:"required,min=1"`
}

type updateDiaryRequestBody struct {
	Exercise int64 `json:"exercise" binding:"required"`
	Meal     int64 `json:"meal" binding:"required"`
}

func (server *Server) updateDiary(ctx *gin.Context) {
	var uri updateDiaryUriParams
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqBody updateDiaryRequestBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateDiaryEntryParams{
		DateOf:   uri.DateOf,
		UserID:   uri.UserID,
		Exercise: reqBody.Exercise,
		Meal:     reqBody.Meal,
	}

	diary, err := server.store.UpdateDiaryEntry(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, diary)
}

type deleteDiaryRequest struct {
	DateOf time.Time `uri:"date_of" binding:"required"`
	UserID int64     `uri:"user_id" binding:"required,min=1"`
}

func (server *Server) deleteDiary(ctx *gin.Context) {
	var req deleteDiaryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteDiaryEntryParams{
		DateOf: req.DateOf,
		UserID: req.UserID,
	}

	err := server.store.DeleteDiaryEntry(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Diary entry deleted"})
}

func (server *Server) deleteUserDiaries(ctx *gin.Context) {
	var req getDiaryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteUserDiaryEntries(ctx, req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "All diary entries deleted"})
}
