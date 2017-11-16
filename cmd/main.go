package main

import (
	"try/bidder"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	c, e := bidder.ReadConfig("config.json")
	if e != nil {
		fmt.Println(e)
	}

	pool := bidder.NewRedisPool(c)
	handlers := bidder.NewHandlers(pool, c)

	r := gin.Default()
	r.Use(handlers.AddStatMiddleware)
	r.POST("/", handlers.IncrementIfaStat)
	r.GET("/stats", handlers.GetStat)
	r.GET("/stats/sort", handlers.GetSortedStat)

	log.Fatal(r.Run(":8080"))

}
