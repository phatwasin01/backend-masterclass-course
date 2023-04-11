package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/phatwasin01/ticketx-line-oa/db/sqlc"
	"github.com/phatwasin01/ticketx-line-oa/util"
)

type createOrganizerRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}
type newOrganizerResponse struct {
	Name  string         `json:"name"`
	Email string         `json:"email" `
	Phone sql.NullString `json:"phone"`
}

func (server *Server) createOrganizer(ctx *gin.Context) {
	var req createOrganizerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateOrganizerParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Phone: sql.NullString{
			String: req.Phone,
			Valid:  true,
		},
	}

	user, err := server.store.CreateOrganizer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newOrganizerResponse{
		Name:  user.Name,
		Email: user.Name,
		Phone: user.Phone,
	}
	ctx.JSON(http.StatusOK, rsp)

}

type loginOrganizerRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type loginOrganizerResponse struct {
	Name                 string         `json:"name"`
	Email                string         `json:"email" `
	Phone                sql.NullString `json:"phone" `
	AccessToken          string         `json:"access_token"`
	AccessTokenExpiresAt time.Time      `json:"access_token_expires_at"`
}

func (server *Server) loginOrganizer(ctx *gin.Context) {
	var req loginOrganizerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	organizer, err := server.store.GetOrganizer(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = util.CheckPassword(req.Password, organizer.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		organizer.Email,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginOrganizerResponse{
		Name:                 organizer.Name,
		Email:                organizer.Name,
		Phone:                organizer.Phone,
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, rsp)

}
