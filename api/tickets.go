package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/phatwasin01/ticketx-line-oa/db/sqlc"
)

type getTicketRequest struct {
	OrderID int64 `uri:"order_id" binding:"required,min=1"`
}

func (server *Server) getTicketOrder(ctx *gin.Context) {
	var req getTicketRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet("user_info").(*LineAuthResponse)
	user, err := server.store.GetUserLine(ctx, authPayload.Sub)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.GetTicketOrderParams{
		OrderID: req.OrderID,
		UserID:  user.ID,
	}
	tickets, err := server.store.GetTicketOrder(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tickets)

}
