package routers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

const board_num int =3
const sensor_num int =6
//测试模拟地区,分别是光电大楼，北京某高尔夫球场，深圳的老三和人才市场  EN
var place_test = [board_num][2]float64{{114.440429,30.513163},{116.448569,40.012752},{114.029224,22.666757}} 

type Msg struct {
	Begin     int64 `json:"begin"   binding:"required"`
	End int64 `json:"end" binding:"required"`
	Interval int64 `json:"interval" binding:"required"`
}

type DataMsg struct {
	BoardName     string `json:"board_name"   binding:"required"`
	E float64 `json:"E" binding:"required"`
	N float64 `json:"N" binding:"required"`
	SensorData []int64 `json:"data" binding:"required,lte=6" `
}

type Timedata struct {
	Time 	int64
	Sensor  [sensor_num]int64
	Place [2]float64
}

type boardData struct {
	Id 	int
	Name  string
	Timedata []Timedata
}



func TestGetDataHandler(c *gin.Context) {
	var form Msg
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBindJSON(&form); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if form.Begin > form.End  {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "开始时间大于终止时间"})
		return
	}

	var return_data [board_num]boardData

	for i := 0; i < board_num; i++ {
		var board_data boardData
		board_data.Id = i
		board_data.Name = fmt.Sprint(i)

		var target_data Timedata
		for target_time := form.Begin; target_time  <= form.End ; target_time+=form.Interval{
			target_data.Time = target_time
			for  sensor_id := 0; sensor_id < sensor_num; sensor_id++  {
				target_data.Sensor[sensor_id] = rand.Int63n(100)
			}
			target_data.Place=place_test[i]
			board_data.Timedata = append(board_data.Timedata,target_data)
		}
		return_data[i] = board_data
	}

	c.IndentedJSON(200,return_data)
	
}

func TestPushDataHandler(c *gin.Context) {
	var form DataMsg
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBindJSON(&form); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var target_data Timedata
	board_name := form.BoardName
	target_data.Time= time.Now().Unix()
	target_data.Place = [2]float64{form.E,form.N}
	for i := 0; i < len(form.SensorData); i++ {
		target_data.Sensor[i] = form.SensorData[i]
	}
	fmt.Printf("收到来自%v的传感器信号:%+v  ",board_name,target_data)
	fmt.Println()
	c.IndentedJSON(200, gin.H{"time": time.Now().Unix()})
	
}