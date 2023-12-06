package api

import (
	"fmt"
	"net/http"

	"github.com/buddhimaaushan/mini_bank/db/sqlc"
	"github.com/buddhimaaushan/mini_bank/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type userResponse struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username" `
	Nic       string `json:"nic"`
	Email     string `json:"email"`
	Phone     string `json:"phone" `
}

type createAndLoginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

type createUserRequest struct {
	FirstName string `json:"first_name" binding:"required,min=2,max=16"`
	LastName  string `json:"last_name" binding:"required,min=2,max=16"`
	Username  string `json:"username" binding:"required,min=4,max=16,alphanum"`
	Nic       string `json:"nic" binding:"required,min=10,max=16"`
	Password  string `json:"password" binding:"required,min=8,max=16"`
	Email     string `json:"email" binding:"required,email,min=16,max=48,emailType"`
	Phone     string `json:"phone" binding:"required,min=11,max=13"`
}

func newUserResponse(user sqlc.User) userResponse {
	return userResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Nic:       user.Nic,
		Email:     user.Email,
		Phone:     user.Phone,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	// Check if the request body is valid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Hash the password
	hashedPassword, err := utils.GenerateHashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Create user arguments
	arg := sqlc.CreateUserParams{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Username:       req.Username,
		Nic:            req.Nic,
		HashedPassword: hashedPassword,
		Email:          req.Email,
		Phone:          req.Phone,
	}

	// Create user
	user, err := server.Store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Create access token
	accessToken, err := server.tokenMaker.CreateToken(
		user.Username,
		server.Config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Create user response
	userResponse := newUserResponse(user)

	// Return the user and access token
	ctx.JSON(http.StatusOK, createAndLoginUserResponse{
		AccessToken: accessToken,
		User:        userResponse,
	})

}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,min=4,max=16,alphanum"`
	Password string `json:"password" binding:"required,min=8,max=16"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest

	// Check if the request body is valid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get user
	user, err := server.Store.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Check if the password is correct
	if !utils.CheckPasswordHash(req.Password, user.HashedPassword) {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("incorrect password")))
		return
	}

	// Create access token
	accessToken, err := server.tokenMaker.CreateToken(
		user.Username,
		server.Config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Create user response
	userResponse := newUserResponse(user)

	// Return the user and access token
	ctx.JSON(http.StatusOK, createAndLoginUserResponse{
		AccessToken: accessToken,
		User:        userResponse,
	})

}
