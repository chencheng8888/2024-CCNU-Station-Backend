package response

import (
	"github.com/gin-gonic/gin"
	"guizizhan/common"
	"guizizhan/model/activity"
)

type Recruits_data struct {
	YNLogin  bool               `json:"YNLogin"`
	Recruits []activity.Recruit `json:"recruits"`
}
type GetRecruits_resp struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data Recruits_data `json:"data"`
}

func GetRecruits_ok(c *gin.Context, Recruits []activity.Recruit, msg string) {
	var data = Recruits_data{
		YNLogin:  true,
		Recruits: Recruits,
	}
	c.JSON(200, gin.H{
		"code": common.SUCCESS,
		"msg":  msg,
		"data": data,
	})
}
func GetRecruits_fail(c *gin.Context) {
	var data = Recruits_data{
		YNLogin:  false,
		Recruits: nil,
	}
	c.JSON(200, gin.H{
		"code": common.FAIL,
		"msg":  "未登录",
		"data": data,
	})
}
