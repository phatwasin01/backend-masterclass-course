package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/phatwasin01/ticketx-line-oa/db/sqlc"
)

// type createUserRequest struct {
// 	UserID      string `json:"user_id" binding:"required"`
// 	Email       string `json:"email" binding:"required"`
// 	DisplayName string `json:"display_name" binding:"required"`
// }

// func (server *Server) createUser(ctx *gin.Context) {
// 	var req createUserRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	arg := db.CreateUserParams{
// 		UserID:      req.UserID,
// 		Email:       req.Email,
// 		DisplayName: req.DisplayName,
// 	}

// 	user, err := server.store.CreateUser(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, user)

// }

func (server *Server) createUser(ctx *gin.Context) {

	authPayload := ctx.MustGet("user_info").(LineAuthResponse)
	arg := db.CreateUserParams{
		UserID:      authPayload.Sub,
		Email:       authPayload.Email,
		DisplayName: authPayload.Name,
	}
	if authPayload.Sub == "" {
		ctx.JSON(http.StatusInternalServerError, "Empty UserID")
		return
	}
	fmt.Println("authPayload ID:", authPayload.Sub)
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			//Be careful if the user is the same person or not??!!!!
			case "unique_violation":
				ctx.JSON(http.StatusOK, authPayload)
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	fmt.Println("Response:", user)
	ctx.JSON(http.StatusOK, user)

}
