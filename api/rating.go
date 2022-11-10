package api

import (
	db "backend/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type addRatingRequest struct {
	Price   int32     `json:"price" binding:"required,min=1,max=5"`
	Quality int32     `json:"quality" binding:"required,min=1,max=5"`
	BoothId uuid.UUID `json:"booth_id" binding:"required"`
}

func (s *Server) AddRating(ctx *gin.Context) {
	var req addRatingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddRatingParams{
		Price:   req.Price,
		Quality: req.Quality,
		BoothID: req.BoothId,
	}

	rating, err := s.store.AddRating(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, rating)
}

type listRatingsByBoothIdRequest struct {
	BoothId uuid.UUID `uri:"booth_id" binding:"required"`
}

func (s *Server) ListRatingsByBoothId(ctx *gin.Context) {
	var req listRatingsByBoothIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ratings, err := s.store.ListRatingsByBoothId(ctx, req.BoothId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ratings)
}
