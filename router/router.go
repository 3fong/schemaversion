package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/list", listSchema)
	r.POST("/get", getSchema)

	return r
}

func listSchema(c *gin.Context) {
	log.Printf("aaaaa")
}

func getSchema(c *gin.Context) {
	log.Printf("get")
}
