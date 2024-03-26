package response

import (
	"github.com/gin-gonic/gin"
	"guizizhan/common"
	"guizizhan/model/activity"
)

type Treasurehuntings_data struct {
	YNLogin          bool                       `json:"YNLogin"`
	Treasurehuntings []activity.Treasurehunting `json:"treasurehuntings"`
}
type GetTreasurehuntings_resp struct {
	Code int                   `json:"code"`
	Msg  string                `json:"msg"`
	Data Treasurehuntings_data `json:"data"`
}

func GetTreasurehuntings_ok(c *gin.Context, Treasurehuntings []activity.Treasurehunting, msg string) {
	var data = Treasurehuntings_data{
		YNLogin:          true,
		Treasurehuntings: Treasurehuntings,
	}
	c.JSON(200, gin.H{
		"code": common.SUCCESS,
		"msg":  msg,
		"data": data,
	})
}
func GetTreasurehuntings_fail(c *gin.Context) {
	var data = Treasurehuntings_data{
		YNLogin:          false,
		Treasurehuntings: nil,
	}
	c.JSON(200, gin.H{
		"code": common.FAIL,
		"msg":  "未登录",
		"data": data,
	})
}
