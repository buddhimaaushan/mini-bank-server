package api

import (
	"net/http"
	"time"

	"github.com/buddhimaaushan/mini_bank/db/sqlc"
	app_error "github.com/buddhimaaushan/mini_bank/errors"
	"github.com/buddhimaaushan/mini_bank/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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
		ctx.JSON(http.StatusBadRequest, errorResponse(app_error.ApiError.ErrInvalidRequest.Wrap(err)))
		return
	}

	// Hash the password
	hashedPassword, err := utils.GenerateHashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(app_error.HashError.ErrHashPassword))
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(app_error.DbError.ErrCreateUser))
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
		ctx.JSON(http.StatusBadRequest, errorResponse(app_error.ApiError.ErrInvalidRequest.Wrap(err)))
		return
	}

	// Get user
	user, err := server.Store.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(app_error.ApiError.ErrErrInvalidUsernameOrPassword))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(app_error.ApiError.ErrDataFetching))
		return
	}

	// Check if the password is correct
	if !utils.CheckPasswordHash(req.Password, user.HashedPassword) {
		ctx.JSON(http.StatusUnauthorized, errorResponse(app_error.ApiError.ErrErrInvalidUsernameOrPassword))
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
func (server *Server) createTokens(res *Response, sessionID uuid.UUID) error {
	// Create access token
	accessToken, accessPayload, err := server.TokenMaker.CreateToken(
		uuid.New(),
		res.User.ID,
		res.User.Username,
		res.User.Role,
		res.User.AccStatus,
		server.Config.AccessTokenDuration,
	)
	if err != nil {
		return app_error.TokenError.ErrCreateToken
	}

	// Create refresh token
	refreshToken, refreshPayload, err := server.TokenMaker.CreateToken(
		sessionID,
		res.User.ID,
		res.User.Username,
		res.User.Role,
		res.User.AccStatus,
		server.Config.RefreshTokenDuration,
	)
	if err != nil {
		return app_error.TokenError.ErrCreateToken
	}

	// Add token data to the response
	res.AccessToken = accessToken
	res.AccessTokenExpiresAt = accessPayload.ExpiresAt
	res.RefreshToken = refreshToken
	res.RefreshTokenExpiresAt = refreshPayload.ExpiresAt

	return nil
}

// Create session
func (server *Server) createSession(ctx *gin.Context, res *Response) (*sqlc.Session, error) {
	session, err := server.Store.CreateSession(ctx, sqlc.CreateSessionParams{
		ID:           utils.NewUUID(),
		Username:     res.User.Username,
		RefreshToken: res.RefreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    utils.TimeToPgTime(res.RefreshTokenExpiresAt),
	})
	if err != nil {
		return nil, app_error.DbError.ErrCreateSession.Wrap(err)
	}

	return &session, nil
}

type userResponse struct {
	ID        int64       `json:"id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Username  string      `json:"username" `
	Nic       string      `json:"nic"`
	Email     string      `json:"email"`
	Phone     string      `json:"phone" `
	AccStatus sqlc.Status `json:"acc_status" `
	Role      string      `json:"role" `
}

type Response struct {
	AccessToken           string        `json:"access_token"`
	AccessTokenExpiresAt  time.Time     `json:"access_token_expires_at"`
	RefreshToken          string        `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time     `json:"refresh_token_expires_at"`
	User                  *userResponse `json:"user"`
}

// createResponse creates a new response
func createResponse(ctx *gin.Context, server *Server, user sqlc.User) (res *Response, err error) {
	// Init response
	res = &Response{}

	// Create user response
	res.User = &userResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Nic:       user.Nic,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role.String,
		AccStatus: user.AccStatus,
	}

	// Create session
	session, err := server.createSession(ctx, res)
	if err != nil {
		return nil, err
	}

	// Creates  authentication and authorization tokens
	err = server.createTokens(res, session.ID.Bytes)
	if err != nil {
		return nil, err
	}

	// Return the user and access token
	return res, nil
}
