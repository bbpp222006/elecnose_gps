package routers

import (
    "github.com/gin-gonic/gin"
)



func SetupRouter(e *gin.Engine)  {
    e.POST("/test/get_data", TestGetDataHandler)
    e.POST("/test/push_data", TestPushDataHandler)
}

