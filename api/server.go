package api

import (
	"github.com/buddhimaaushan/mini_bank/db"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	Store  db.Store
	Router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{Store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("accountStatus", validAccountStatus)
	}

	// Routes for accounts
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.GetAccounts)
	router.GET("/accounts/:id", server.GetAccount)

	// Routes for transfers
	router.POST("/transfers", server.createTransfer)

	server.Router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

// errorResponse returns error in JSON format
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
