package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/phatwasin01/ticketx-line-oa/db/sqlc"
	"github.com/phatwasin01/ticketx-line-oa/token"
	"github.com/phatwasin01/ticketx-line-oa/util"
)

type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()
	authRoutesLine := router.Group("/").Use(lineAuthMiddleware(config))
	authRoutesOrganizer := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutesOrganizer.POST("/event", server.createEvent)
	authRoutesLine.POST("/user", server.createUser)
	router.POST("/organizer", server.createOrganizer)
	router.POST("/login", server.loginOrganizer)
	router.GET("/event", server.listEvents)
	router.GET("/event/:id", server.getEvent)
	authRoutesLine.POST("/order", server.createOrder)
	authRoutesLine.GET("/order", server.getOrders)
	authRoutesLine.GET("/ticket/:order_id", server.getTicketOrder)
	server.router = router
	return server, nil
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
