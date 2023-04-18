package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/phatwasin01/ticketx-line-oa/db/sqlc"
)

type createOrderRequest struct {
	EventID int64 `json:"event_id" binding:"required"`
	Amount  int32 `json:"amount" binding:"required,min=1"`
	// Payment  sql.NullString `json:"payment" binding:"required"`
}

func (server *Server) createOrder(ctx *gin.Context) {
	var req createOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet("user_info").(LineAuthResponse)

	event, err := server.store.GetEvent(ctx, req.EventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateOrderParams{
		UserID:   authPayload.Sub,
		EventID:  req.EventID,
		Amount:   req.Amount,
		SumPrice: event.Price * req.Amount,
		// Payment:  req.Payment,
	}

	user, err := server.store.CreateOrderTickets(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)

}

func (server *Server) getOrders(ctx *gin.Context) {
	authPayload := ctx.MustGet("user_info").(LineAuthResponse)

	event, err := server.store.ListOrdersUser(ctx, authPayload.Sub)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, event)

}
