package api

import (
	"net/http"

	"github.com/buddhimaaushan/mini_bank/db"
	"github.com/buddhimaaushan/mini_bank/db/sqlc"
	app_error "github.com/buddhimaaushan/mini_bank/errors"
	"github.com/buddhimaaushan/mini_bank/token"
	"github.com/gin-gonic/gin"
)

type createTransferRequest struct {
	FromAccountID int64 `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64 `json:"to_account_id" binding:"required,min=1"`
	Amount        int64 `json:"amount" binding:"required,gt=0"`
}

// CreateTransfer creates a new transfer
func (server *Server) createTransfer(ctx *gin.Context) {
	var req createTransferRequest

	// Check if the request body is valid
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(app_error.ApiError.ErrInvalidRequest.Wrap(err)))
		return
	}

	// Create a new transfer
	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	// Verify access roles
	payload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if payload.Role == "admin" || payload.Role == "bankTeller" {
		arg.TransferedByID = payload.UserID
	} else if payload.Role == "customer" {
		accountHolders, err := server.Store.GetAccountHoldersByAccountID(ctx, sqlc.GetAccountHoldersByAccountIDParams{
			AccID:  req.FromAccountID,
			Limit:  10,
			Offset: 0,
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse(app_error.ApiError.ErrUnauthorized))
		}
		for i := range accountHolders {
			if accountHolders[i].UserID == payload.UserID {
				arg.TransferedByID = payload.UserID
			}
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, errorResponse(app_error.ApiError.ErrUnauthorized))
	}

	// Create the transfer
	transfer, err := server.Store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Return the transfer
	ctx.JSON(http.StatusOK, transfer)
}
