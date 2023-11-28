package api

import (
	"net/http"

	"github.com/buddhimaaushan/mini_bank/db"
	"github.com/buddhimaaushan/mini_bank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type createAccountRequest struct {
	Type    string  `json:"type" binding:"required"`
	UserIDs []int64 `json:"user_ids" binding:"required,gt=0"`
}

// CreateAccount creates a new account
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	// Check if the request body is valid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Arguments for a new account
	arg := db.AccountTxParams{
		Type:      req.Type,
		Balance:   0,
		AccStatus: sqlc.StatusInactive,
		UserIDs:   req.UserIDs,
	}

	// Create account and account holders
	result, err := server.Store.AccountTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Return the account
	ctx.JSON(http.StatusOK, result)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// GetAccount gets an account
func (server *Server) GetAccount(ctx *gin.Context) {
	var req getAccountRequest

	// Check if the request body is valid
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get the account
	account, err := server.Store.GetAccount(ctx, req.ID)
	if err != nil {
		// Check if the error is a pgx.ErrNoRows
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Return the account
	ctx.JSON(http.StatusOK, account)
}

type getAccountsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// GetAccounts gets accounts
func (server *Server) GetAccounts(ctx *gin.Context) {
	var req getAccountsRequest

	// Check if the request body is valid
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Get accounts Pagination
	arg := sqlc.GetAccountsParams{
		LimitNo:  req.PageSize,
		OffsetNo: (req.PageID - 1) * req.PageSize,
	}

	// Get the accounts
	accounts, err := server.Store.GetAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Return the accounts
	ctx.JSON(http.StatusOK, accounts)
}
