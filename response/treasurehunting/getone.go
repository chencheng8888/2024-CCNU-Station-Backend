package response

import (
	"github.com/gin-gonic/gin"
	"guizizhan/common"
	"guizizhan/model/activity"
)

type Treasurehunting_data struct {
	YNLogin         bool                     `json:"YNLogin"`
	Treasurehunting activity.Treasurehunting `json:"treasurehunting"`
}
type Treasurehunting_resp struct {
	Code int                  `json:"code"`
	Msg  string               `json:"msg"`
	Data Treasurehunting_data `json:"data"`
}

func GetTheTreasurehunting_ok(c *gin.Context, treasurehunting activity.Treasurehunting, msg string) {
	var data = Treasurehunting_data{
		YNLogin:         true,
		Treasurehunting: treasurehunting,
	}
	c.JSON(200, gin.H{
		"code": common.SUCCESS,
		"msg":  msg,
		"data": data,
	})
}
func GetTheTreasurehunting_fail(c *gin.Context) {
	var data = Treasurehunting_data{
		YNLogin: false,
	}
	c.JSON(200, gin.H{
		"code": common.FAIL,
		"msg":  "未登录",
		"data": data,
	})
}
