package main

import (
    "github.com/gin-gonic/gin"
    "elecnose_gps/routers"
    // "fmt"
)
    

func main() {
	r := gin.Default()
	routers.SetupRouter(r)
	// Listen and serve on http://127.0.0.1:8080
	r.Run(":8080")
}