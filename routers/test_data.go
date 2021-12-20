package routers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"math/rand"
)

const board_num int =3
const sensor_num int =6
//测试模拟地区,分别是光电大楼，北京某高尔夫球场，深圳的老三和人才市场
var place_test = [board_num][2]float64{{114.440429,30.513163},{116.448569,40.012752},{114.029224,22.666757}} 

type Msg struct {
	Begin     int32 `json:"begin"   binding:"required"`
	End int32 `json:"end" binding:"required"`
	Interval int32 `json:"interval" binding:"required"`
}

type Timedata struct {
	Time 	int32
	Sensor  [sensor_num]int32
	Place [2]float64
}

type boardData struct {
	Id 	int
	Name  string
	Timedata []Timedata
}



func TestDataHandler(c *gin.Context) {
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
				target_data.Sensor[sensor_id] = rand.Int31n(100)
			}
			target_data.Place=place_test[i]
			board_data.Timedata = append(board_data.Timedata,target_data)
		}
		return_data[i] = board_data
	}

	c.IndentedJSON(200,return_data)
	
}
