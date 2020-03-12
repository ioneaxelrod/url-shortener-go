package server

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	db *dynamodb.DynamoDB
}

func (s Server) GetDB() *dynamodb.DynamoDB {
	return s.db
}

func (s Server) GetRouter() *gin.Engine {
	return s.router
}

func (s Server) Run() {
	s.router.Run()
}

func New() *Server {
	// Create AWS Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Build Server and DB connection
	server := Server{
		router: gin.Default(),
		db: dynamodb.New(sess),
	}

	// Create Routes
	server.GetRouter().GET("/:key", server.redirectToURL)
	server.GetRouter().POST("/", server.createShortURL)

	return &server
}



