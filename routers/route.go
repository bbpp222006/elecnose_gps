package routers

import (
    "github.com/gin-gonic/gin"
)



func SetupRouter(e *gin.Engine)  {
    e.POST("/test", TestDataHandler)
}

