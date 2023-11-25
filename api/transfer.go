package api

import (
	"net/http"

	"github.com/buddhimaaushan/mini_bank/db"
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
		ctx.JSON(400, errorResponse(err))
		return
	}

	// Create a new transfer
	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	// Create the transfer
	transfer, err := server.Store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}

	// Return the transfer
	ctx.JSON(http.StatusOK, transfer)
}
