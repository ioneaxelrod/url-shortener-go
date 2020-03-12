package server

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener-go/shorten"
)

func (s Server) redirectToURL(c *gin.Context) {
	// Parse key from param
	key := c.Param("key")
	// make call to DynamoDB
		// TODO: Make DynamoDB call
	// redirect to the URL
	c.Redirect(http.StatusPermanentRedirect, longURL)
}


func(s Server) createShortURL(c *gin.Context) {
	// create "key" for url and database
	key, err := shorten.CreateKey()
	if err != nil {
		c.Errors(err)
	}
	// unpack request
	data, err := c.GetRawData()
	//store key and url in database
		// TODO: Make DynamoDB call
	// Return URL in JSON
		// TODO: Format retrieved data
	c.JSON(http.StatusCreated, json)
}