package api

import (
	"net/http"

	"github.com/buddhimaaushan/mini_bank/db"
	"github.com/buddhimaaushan/mini_bank/db/sqlc"
	app_error "github.com/buddhimaaushan/mini_bank/errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type AccountResponse struct {
	ID             int64                `json:"id"`
	Type           string               `json:"type"`
	Balance        int64                `json:"balance"`
	AccountHolders []sqlc.AccountHolder `json:"account_holders"`
	Status         sqlc.Status          `json:"status"`
	CreatedAt      pgtype.Timestamptz   `json:"created_at"`
}

type createAccountRequest struct {
	Type    string  `json:"type" binding:"required"`
	UserIDs []int64 `json:"user_ids" binding:"required,gt=0"`
}

// CreateAccount creates a new account
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	// Check if the request body is valid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(app_error.ApiError.InvalidRequestError.Wrap(err)))
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

	// Create account response
	accountResponse := AccountResponse{
		ID:             result.Account.ID,
		Type:           result.Account.Type,
		Balance:        result.Account.Balance,
		AccountHolders: result.AccountHolders,
	}

	// Return the account
	ctx.JSON(http.StatusOK, accountResponse)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// GetAccount gets an account
func (server *Server) GetAccount(ctx *gin.Context) {
	var req getAccountRequest

	// Check if the request body is valid
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(app_error.ApiError.InvalidRequestError.Wrap(err)))
		return
	}

	// Get the account
	account, err := server.Store.GetAccount(ctx, req.ID)
	if err != nil {
		// Check if the error is a pgx.ErrNoRows
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(app_error.DbError.AccountNotFoundError))
			return
		}

		ctx.JSON(http.StatusBadRequest, errorResponse(app_error.ApiError.FetchingDataError))
		return
	}

	// Get the account holders arguments
	arg := sqlc.GetAccountHoldersByAccountIDParams{
		AccID:  account.ID,
		Limit:  10,
		Offset: 0,
	}

	// Get the account holders
	accountHolders, err := server.Store.GetAccountHoldersByAccountID(ctx, arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(app_error.DbError.AccountHoldersNotFoundError))
			return
		}

		ctx.JSON(http.StatusBadRequest, errorResponse(app_error.ApiError.FetchingDataError))
		return
	}

	// Create account response
	accountResponse := AccountResponse{
		ID:             account.ID,
		Type:           account.Type,
		Balance:        account.Balance,
		AccountHolders: accountHolders,
		Status:         account.AccStatus,
		CreatedAt:      account.CreatedAt,
	}

	// Return the account response
	ctx.JSON(http.StatusOK, accountResponse)

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
		ctx.JSON(http.StatusBadRequest, errorResponse(app_error.ApiError.InvalidRequestError.Wrap(err)))
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(app_error.ApiError.FetchingDataError))
		return
	}

	// Return the accounts
	ctx.JSON(http.StatusOK, accounts)
}
