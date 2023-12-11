package api

import (
	"net/http"
	"time"

	"github.com/buddhimaaushan/mini_bank/db/sqlc"
	app_error "github.com/buddhimaaushan/mini_bank/errors"
	"github.com/buddhimaaushan/mini_bank/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type createUserRequest struct {
	FirstName string `json:"first_name" binding:"required,min=2,max=16"`
	LastName  string `json:"last_name" binding:"required,min=2,max=16"`
	Username  string `json:"username" binding:"required,min=4,max=16,alphanum"`
	Nic       string `json:"nic" binding:"required,min=10,max=16"`
	Password  string `json:"password" binding:"required,min=8,max=16"`
	Email     string `json:"email" binding:"required,email,min=16,max=48,emailType"`
	Phone     string `json:"phone" binding:"required,min=11,max=13"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	// Check if the request body is valid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(app_error.ApiError.InvalidRequestError.Wrap(err)))
		return
	}

	// Hash the password
	hashedPassword, err := utils.GenerateHashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(app_error.HashError.HashPasswordError))
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(app_error.DbError.CreateUserError))
		return
	}

	// Create response
	res, err := createResponse(ctx, server, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return the user and access token
	ctx.JSON(http.StatusOK, res)

}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,min=4,max=16,alphanum"`
	Password string `json:"password" binding:"required,min=8,max=16"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest

	// Check if the request body is valid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(app_error.ApiError.InvalidRequestError.Wrap(err)))
		return
	}

	// Get user
	user, err := server.Store.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(app_error.ApiError.InvalidUsernameOrPasswordError))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(app_error.ApiError.FetchingDataError))
		return
	}

	// Check if the password is correct
	if !utils.CheckPasswordHash(req.Password, user.HashedPassword) {
		ctx.JSON(http.StatusUnauthorized, errorResponse(app_error.ApiError.InvalidUsernameOrPasswordError))
		return
	}

	// Create response
	res, err := createResponse(ctx, server, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Return the user and access token
	ctx.JSON(http.StatusOK, res)

}

// createTokens creates authentication and authorization tokens
func (server *Server) createTokens(res *Response) error {
	// Create access token
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		res.User.Username,
		server.Config.AccessTokenDuration,
	)
	if err != nil {
		return app_error.TokenError.CreateTokenError
	}

	// Create refresh token
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		res.User.Username,
		server.Config.RefreshTokenDuration,
	)
	if err != nil {
		return app_error.TokenError.CreateTokenError
	}

	// Add token data to the response
	res.AccessToken = accessToken
	res.AccessTokenExpiresAt = accessPayload.ExpiresAt
	res.RefreshToken = refreshToken
	res.RefreshTokenExpiresAt = refreshPayload.ExpiresAt

	return nil
}

// Create session
func (server *Server) createSession(ctx *gin.Context, res *Response) error {
	_, err := server.Store.CreateSession(ctx, sqlc.CreateSessionParams{
		Username:     res.User.Username,
		RefreshToken: res.RefreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    pgtype.Timestamptz{Time: res.RefreshTokenExpiresAt},
	})
	if err != nil {
		return app_error.DbError.CreateSessionError
	}

	return nil
}

type userResponse struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username" `
	Nic       string `json:"nic"`
	Email     string `json:"email"`
	Phone     string `json:"phone" `
}

type Response struct {
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  userResponse `json:"user"`
}

// CreateUser creates a new response
func createResponse(ctx *gin.Context, server *Server, user sqlc.User) (res *Response, err error) {
	// Create user response
	res.User = userResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Nic:       user.Nic,
		Email:     user.Email,
		Phone:     user.Phone,
	}

	// Creates  authentication and authorization tokens
	err = server.createTokens(res)
	if err != nil {
		return nil, err
	}

	// Create session
	err = server.createSession(ctx, res)
	if err != nil {
		return nil, err
	}

	// Return the user and access token
	return res, nil
}
