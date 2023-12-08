package api

import (
	"github.com/buddhimaaushan/mini_bank/db"
	app_error "github.com/buddhimaaushan/mini_bank/errors"
	"github.com/buddhimaaushan/mini_bank/token"
	"github.com/buddhimaaushan/mini_bank/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	Config     utils.Config
	Store      db.Store
	tokenMaker token.Maker
	Router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	// Create a new token maker
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, app_error.TokenError.CreateTokenMakerError.Wrap(err)
	}

	// Create a new server
	server := &Server{Config: config, Store: store, tokenMaker: tokenMaker}

	// Set validators for gin
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("accountStatus", validAccountStatus)
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("emailType", validEmailTypes)
	}

	// Setup routing
	server.setupRouter()

	return server, nil
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

// errorResponse returns error in JSON format
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupRouter() {
	// Setup routing
	router := gin.Default()

	// Routes for authentication
	router.POST("/register", server.createUser)
	router.POST("/login", server.loginUser)

	// Routes for accounts
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.GetAccounts)
	router.GET("/accounts/:id", server.GetAccount)

	// Routes for transfers
	router.POST("/transfers", server.createTransfer)

	server.Router = router
}
