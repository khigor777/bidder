package main

import (
	"try/bidder"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	c, e := bidder.ReadConfig("config.json")

	if e != nil {
		panic(e)
	}

	pool := bidder.NewRedisPool(c)
	handlers := bidder.NewHandlers(pool, c)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(handlers.AddStatMiddleware)
	r.POST("/", handlers.IncrementIfaStat)
	r.GET("/stats", handlers.GetStat)
	r.GET("/stats/sort", handlers.GetSortedStat)

	log.Fatal(r.Run(":8080"))

}
