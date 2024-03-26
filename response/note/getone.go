package response

import (
	"github.com/gin-gonic/gin"
	"guizizhan/common"
	"guizizhan/model"
	"guizizhan/model/activity"
)

type Note_data struct {
	YNLogin    bool            `json:"YNLogin"`
	Note       activity.Post   `json:"note"`
	Note_image model.PostImage `json:"note_image"`
}
type GetNote_resp struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data Note_data `json:"data"`
}

func GetTheNote_ok(c *gin.Context, post activity.Post, postimage model.PostImage, msg string) {
	var data = Note_data{
		YNLogin:    true,
		Note:       post,
		Note_image: postimage,
	}
	c.JSON(200, gin.H{
		"code": common.SUCCESS,
		"msg":  msg,
		"data": data,
	})
}
func GetTheNote_fail(c *gin.Context) {
	var data = Note_data{
		YNLogin: false,
	}
	c.JSON(200, gin.H{
		"code": common.FAIL,
		"msg":  "未登录",
		"data": data,
	})
}
