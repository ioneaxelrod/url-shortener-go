package server

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"url-shortener-go/shorten"
)

type Server struct {
	Router *gin.Engine
	DB *dynamodb.DynamoDB
}

func (Server) GetURL(c *gin.Context){
	// Retrieve url from database using key
	// Return long url
	key := c.Param("key")
	// Redirect
}

func(Server) CreateShortURL(c *gin.Context) {
	key, err := shorten.CreateKey()
	if err != nil {
		c.Errors(err)
	}
	// unpack request
	data, err := c.GetRawData()
	//store key and url in database

}


func Init() {
	// Create AWS Session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	server := Server{
		Router: gin.Default(),
		DB: dynamodb.New(sess),
	}

	server.Router.GET(	"/:key", server.GetURL)
	server.Router.POST(	"/", server.CreateShortURL)

	server.Router.Run()
}


}



