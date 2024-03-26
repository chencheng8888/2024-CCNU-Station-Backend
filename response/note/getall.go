package response

import (
	"github.com/gin-gonic/gin"
	"guizizhan/common"
	"guizizhan/model/activity"
)

type Notes_data struct {
	YNLogin bool            `json:"YNLogin"`
	Notes   []activity.Post `json:"notes"`
}
type GetNotes_resp struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data Notes_data `json:"data"`
}

func Getallnotes_ok(c *gin.Context, posts []activity.Post, msg string) {

	var data = Notes_data{
		YNLogin: true,
		Notes:   posts,
	}
	c.JSON(200, gin.H{
		"code": common.SUCCESS,
		"msg":  msg,
		"data": data,
	})
}
func Getallnotes_fail(c *gin.Context) {
	var data = Notes_data{
		YNLogin: false,
		Notes:   nil,
	}
	c.JSON(200, gin.H{
		"code": common.FAIL,
		"msg":  "未登录",
		"data": data,
	})
}
