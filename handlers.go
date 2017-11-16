package bidder

import (
	"github.com/gin-gonic/gin"
"github.com/gin-gonic/gin/binding"
	"net/http"

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
		if c.Request.Method == http.MethodPost {
			var stat Statistics
			if c.BindJSON(&stat) == nil{
				h.Redis.AddStat(&stat)
			}
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
	r, _ := h.Redis.IncrIfa(&pos, h.Config.LongSeries)
	c.JSON(http.StatusOK, gin.H{"pos":r})
}


func BindJSON(c *gin.Context, obj interface{}) error {
	if err := binding.JSON.Bind(c.Request, obj); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return err
	}
	return nil
}