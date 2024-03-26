package response

import "github.com/gin-gonic/gin"

type Achi_data struct {
	Finished string `json:"Finished"`
}
type Getachi_resp struct {
	Code string    `json:"Code"`
	Data Achi_data `json:"data"`
	Msg  string    `json:"msg"`
}

func Getachi_ok(c *gin.Context, finished string) {
	var data = Achi_data{
		Finished: finished,
	}
	c.JSON(200, gin.H{
		"code": SUCCESS,
		"data": data,
		"msg":  "成功获取所有成就信息",
	})
}
func Getachi_fail(c *gin.Context) {
	var data = Achi_data{
		Finished: "",
	}
	c.JSON(200, gin.H{
		"code": FAIL,
		"data": data,
		"msg":  "未找到该用户",
	})
}
func Updateachi_ok(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": SUCCESS,
		"msg":  "修改成功",
	})
}
