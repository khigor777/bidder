package bidder

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)


type Handlers struct {
	Redis *RedisPool
	Config *Config
}

func NewHandlers(r *RedisPool, c *Config) *Handlers {
	return &Handlers{Redis:r, Config:c}
}

//Add statistic for all json request in middleware key - country:app:platform:count
func (h *Handlers) AddStatMiddleware(c *gin.Context) {
	fmt.Println(c.Request.Method )
	if c.Request.Method == http.MethodPost {
		var stat Statistics
		c.Bind(&stat)
		h.Redis.AddStat(&stat)
	}

}

//Get statistics from all json request ref:AddStatMiddleware
func (h *Handlers) GetStat(c *gin.Context) {
	c.JSON(http.StatusOK, h.Redis.GetStat())
}

//Get statistics from all json request ref:AddStatMiddleware
func (h *Handlers) GetSortedStat(c *gin.Context) {
	c.JSON(http.StatusOK, h.Redis.GetSortedStat())
}

//Set and Get ifa statistic
func (h *Handlers) IncrementIfaStat(c *gin.Context) {
	var pos StatisticsForIfa
	c.Bind(&pos)
	r, e  := h.Redis.IncrIfa(&pos, h.Config.ShortSeries)
	if e != nil{
		c.JSON(http.StatusBadRequest, e)
	}
	c.JSON(http.StatusOK, gin.H{"pos":r})
}
