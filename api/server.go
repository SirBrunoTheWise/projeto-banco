package api

import (
	db "github.com/SirBrunoTheWise/hunt/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.PUT("/users/:id", server.updateUser)
	router.DELETE("/users/:id", server.deleteUser)

	router.POST("/cards", server.createCard)
	router.GET("/cards/:id", server.getCard)
	router.PUT("/cards/:id", server.updateCard)
	router.DELETE("/cards/:id", server.deleteCard)

	router.POST("/diary", server.createDiary)
	router.GET("/diary/:date_of/:user_id", server.getDiary)
	router.PUT("/diary/:date_of/:user_id", server.updateDiary)
	router.DELETE("/diary/:date_of/:user_id", server.deleteDiary)

	server.router = router
	return server
}

func (server *Server) Start(adress string) error {
	return server.router.Run(adress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
