package response

import (
	"github.com/gin-gonic/gin"
	"guizizhan/common"
	"guizizhan/model/activity"
)

type Recruit_data struct {
	YNLogin bool             `json:"YNLogin"`
	Recruit activity.Recruit `json:"recruit"`
}
type GetRecruit_resp struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data Recruit_data `json:"data"`
}

func GetTheRecruit_ok(c *gin.Context, recruit activity.Recruit, msg string) {
	var data = Recruit_data{
		YNLogin: true,
		Recruit: recruit,
	}
	c.JSON(200, gin.H{
		"code": common.SUCCESS,
		"msg":  msg,
		"data": data,
	})
}
func GetTheRecruit_fail(c *gin.Context) {
	var data = Recruit_data{
		YNLogin: false,
	}
	c.JSON(200, gin.H{
		"code": common.FAIL,
		"msg":  "未登录",
		"data": data,
	})
}
