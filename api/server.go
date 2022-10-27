package api

import (
	db "backend/db/sqlc"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server serves our HTTP requests.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// This is removed since it is not needed in production, and it poses a security risk.
	// router.POST("/booth", server.AddBooth)
	router.GET("/booth/:id", server.GetBooth)
	router.GET("/booth", server.ListBooth)

	router.POST("/rating", server.AddRating)
	router.GET("/rating/:booth_id", server.ListRatingsByBoothId)

	server.router = router
	return server
}

// Start runs the HTTP server on the specified address.
func (server *Server) Start(port int) error {
	address := fmt.Sprintf("0.0.0.0:%d", port)

	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
