package api

import (
	"net/http"
	"strings"

	app_error "github.com/buddhimaaushan/mini_bank/errors"
	"github.com/buddhimaaushan/mini_bank/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get authorization header
		authorizationHeader := c.GetHeader(authorizationHeaderKey)

		// Check if authorization header is present
		if len(authorizationHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(app_error.ApiError.ErrMissingAuthHeader))
			return
		}

		// Get token and token type from authorization header
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(app_error.ApiError.ErrInvalidAuthHeader))
			return
		}

		// Verify authorization type
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(app_error.ApiError.ErrInvalidAuthType))
			return
		}

		// Verify token
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(app_error.TokenError.ErrInvalidToken))
			return
		}

		// Check account status
		if payload.AccStatus != "active" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(app_error.ApiError.ErrAccountNotActive))
			return
		}

		// Set payload
		c.Set(authorizationPayloadKey, payload)

		// Next
		c.Next()
	}
}
