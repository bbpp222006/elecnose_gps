package routers

// import (
//     "github.com/gin-gonic/gin"
//     "net/http"
// )

// type Login struct {
// 	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
// 	Password string `form:"password" json:"password" xml:"password" binding:"required"`
// }

// func TestDataHandler(c *gin.Context) {
// 	var form Login
// 	// This will infer what binder to use depending on the content-type header.
// 	if err := c.ShouldBind(&form); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if form.User != "manu" || form.Password != "123" {
// 		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
// }



// func SetupRouter(e *gin.Engine)  {
//     e.POST("/test", TestDataHandler)
// }

