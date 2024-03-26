package note

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"guizizhan/model/activity"
	response "guizizhan/response/note"
	"strconv"
)

// GetAllPostNote 获取特定用户发布的所有帖子的接口。
// @Summary 获取所有帖子
// @Description postid是指帖子的ID，poster是发布人的ID，postLocation是发布地点，text是发布帖子的内容，time是发布的时间。只有当YNLogin=false,code才会是FAIL即1001，其他时候code为SUCCESS即1000。注意返回的是包含帖子信息的数组。说明一下，并没有返回图片，因为帖子的图片数可能较多。
// @ID get-all-posts
// @Produce json
// @Param where query string true "发帖地点"
// @Security Bearer
// @Api(tags="获取")
// @Success 200 {object} response.GetNotes_resp
// @Failure 200 {object} response.GetNotes_resp
// @Router /api/getactivity/allpostnote [get]
func GetAllPostNote(c *gin.Context, db *gorm.DB) {
	var msg string
	//_, yn := tool.GetStudentID(c)

	wherestring, _ := c.GetQuery("where")
	whereint, _ := strconv.Atoi(wherestring)
	var posts []activity.Post
	res := db.Model(&activity.Post{}).Where(&activity.Post{PostLocation: whereint}).Order("time desc").Find(&posts)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		msg = "这个人没有发布帖子"
	} else {
		msg = "成功获取到所有帖子"
	}

	response.Getallnotes_ok(c, posts, msg)

}
