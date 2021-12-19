package routers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    // "fmt"
)

type Login struct {
	Begin     int32 `json:"begin"   binding:"required"`
	End int32 `json:"end" binding:"required"`
	Interval int32 `json:"interval" binding:"required"`
}

type Msg struct{
	Name string
	Message string
	Number int
}

func TestDataHandler(c *gin.Context) {
	var form Login
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if form.Begin > form.End  {
		c.JSON(http.StatusBadRequest, gin.H{"error": "开始时间大于终止时间"})
		return
	}

	var msg Msg
	msg.Name = "root"
	c.JSON(200,msg)
	
}
