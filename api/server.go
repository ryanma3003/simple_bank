package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/ryanma3003/simplebank/db/sqlc"
)

// server serves HTTP requests for services
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/accounts", server.listAccount)
	router.POST("/accounts", server.createAccount)
	router.GET("/account/:id", server.getAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// func to return error
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
