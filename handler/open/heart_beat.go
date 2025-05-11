package open

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/siriusol/plantscale_db/dal"
	"github.com/siriusol/plantscale_db/model"
)

func LatestHeartBeat(c *gin.Context) {
	beats, err := dal.GetHeartBeat(c, 0, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	if len(beats) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"beat": "empty",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"beat": map[string]interface{}{
			"id":           beats[0].Id,
			"desc":         beats[0].Desc,
			"extra":        beats[0].Extra,
			"created_time": beats[0].CreatedTime,
			"updated_time": beats[0].UpdatedTime,
		},
	})
}

func EmitHeartBeat(c *gin.Context) {
	err := NewHeartBeat(c, "emit")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func NewHeartBeat(ctx context.Context, desc string) error {
	now := time.Now()
	beat := model.HeartBeat{
		Desc:        desc,
		CreatedTime: now,
		UpdatedTime: now,
	}
	return dal.CreateHeartBeat(ctx, &beat)
}
