package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/phatwasin01/ticketx-line-oa/token"
	"github.com/phatwasin01/ticketx-line-oa/util"
)

type LineAuthResponse struct {
	Iss     string   `json:"iss"`
	Sub     string   `json:"sub"`
	Aud     string   `json:"aud"`
	Exp     int      `json:"exp"`
	Iat     int      `json:"iat"`
	Nonce   string   `json:"nonce"`
	Amr     []string `json:"amr"`
	Name    string   `json:"name"`
	Picture string   `json:"picture"`
	Email   string   `json:"email"`
}

func lineAuthMiddleware(config util.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the ID token from the request header
		idToken := c.GetHeader("id_token")
		if idToken == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID token not found in request header"})
			return
		}

		// Get the Channel ID from the request header
		clientID := config.ClientId
		if clientID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Channel ID not found in request header"})
			return
		}

		// Create a new HTTP client
		client := &http.Client{}

		// Create a new POST request
		req, err := http.NewRequest("POST", "https://api.line.me/oauth2/v2.1/verify", nil)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Set the Content-Type header to application/x-www-form-urlencoded
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// Set the id_token and client_id parameters in the request body
		values := url.Values{}
		values.Set("id_token", idToken)
		values.Set("client_id", clientID)
		req.Body = io.NopCloser(strings.NewReader(values.Encode()))

		// Send the request
		resp, err := client.Do(req)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer resp.Body.Close()

		// Read the response body
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		var user LineAuthResponse
		err = json.Unmarshal(responseBody, &user)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Set the response body and content type
		c.Set("user_info", user)
		c.Next()
	}
}

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware creates a gin middleware for authorization
func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
