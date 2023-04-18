package api

import (
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
	authPayload := ctx.MustGet("user_info").(LineAuthResponse)

	arg := db.GetTicketOrderParams{
		OrderID: req.OrderID,
		UserID:  authPayload.Sub,
	}
	tickets, err := server.store.GetTicketOrder(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tickets)

}
