package api

import (
	"fmt"

	"github.com/gin-contrib/cors"
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
	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	router.Use(cors.New(configCors))
	authRoutesLine := router.Group("/").Use(lineAuthMiddleware(config))
	authRoutesOrganizer := router.Group("/").Use(authMiddleware(server.tokenMaker))
	//Create Event
	authRoutesOrganizer.POST("/event", server.createEvent)
	//Create User
	// authRoutesLine.POST("/user", server.createUser)
	router.POST("/user", server.createUser)
	//Create Organizer
	router.POST("/organizer", server.createOrganizer)
	//Login Organizer
	router.POST("/login", server.loginOrganizer)
	//Get User
	router.GET("/event", server.listEvents)
	//Get Event by ID
	router.GET("/event/:id", server.getEvent)
	//Create Order -> Create Ticket
	authRoutesLine.POST("/order", server.createOrder)
	//Get Orders belong to User
	authRoutesLine.GET("/order", server.getOrders)
	//Get Order by ID
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
