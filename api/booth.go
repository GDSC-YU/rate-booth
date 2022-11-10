package api

import (
	db "backend/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type addBoothRequest struct {
	Name         string `json:"name" binding:"required"`
	LogoUrl      string `json:"logo_url" binding:"required"`
	Description  string `json:"description" binding:"required"`
	TwitterUrl   string `json:"twitter_url" binding:"required"`
	InstagramUrl string `json:"instagram_url" binding:"required"`
}

func (s *Server) AddBooth(ctx *gin.Context) {
	var req addBoothRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddBoothParams{
		Name:         req.Name,
		LogoUrl:      req.LogoUrl,
		Description:  req.Description,
		TwitterUrl:   req.TwitterUrl,
		InstagramUrl: req.InstagramUrl,
	}

	booth, err := s.store.AddBooth(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, booth)
}

type getBoothRequest struct {
	ID uuid.UUID `uri:"id" binding:"required,min=1"`
}

func (s *Server) GetBooth(ctx *gin.Context) {
	var req getBoothRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	booth, err := s.store.GetBoothById(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, booth)
}

func (s *Server) ListBooth(ctx *gin.Context) {

	booth, err := s.store.ListBooths(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, booth)
}
