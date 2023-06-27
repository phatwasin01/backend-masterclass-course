package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/phatwasin01/ticketx-line-oa/db/sqlc"
	"github.com/phatwasin01/ticketx-line-oa/token"
)

type createEventRequest struct {
	Name   string `json:"name" binding:"required"`
	Price  int32  `json:"price" binding:"required"`
	Amount int32  `json:"amount" binding:"required"`
	// Description sql.NullString `json:"description" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
}

func (server *Server) createEvent(ctx *gin.Context) {
	var req createEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	organizer, err := server.store.GetOrganizer(ctx, authPayload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateEventParams{
		Name:        req.Name,
		OrganizerID: organizer.ID,
		Price:       req.Price,
		Amount:      req.Amount,
		// Description: req.Description,
		StartTime: req.StartTime,
	}

	event, err := server.store.CreateEvent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, event)

}

func (server *Server) listEvents(ctx *gin.Context) {
	event, err := server.store.ListEventsOpen(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, event)
}

type getEventRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getEvent(ctx *gin.Context) {
	var req getEventRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	event, err := server.store.GetEvent(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, event)

}
