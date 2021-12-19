package main

import (
    "fmt"
    "github.com/imroc/req"
	"github.com/gin-gonic/gin"
)

func main() {
    e,_:=req.Get("http://www.baidu.com")
    fmt.Println(e.String())
}