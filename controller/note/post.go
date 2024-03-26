package note

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"guizizhan/model"
	"guizizhan/model/activity"
	"guizizhan/pkg/qiniu"
	response "guizizhan/response/note"
	"guizizhan/service/generateID"
	"strconv"
	"time"
)

// PostNote 处理发布帖子的请求。
// @Summary 发布帖子
// @Description 返回的信息比较简单，code还是1000表示成功，1000表示失败（未登录），YNLogin代表是否登录，不过code信息已经说明了。
// @ID post-note
// @Accept json
// @Produce json
// @Param where query string true "发布地点"
// @Param key1 query string false "图片1的key"
// @Param text formData string true "帖子的内容"
// @Param title formData string true "帖子的标题"
// @Security Bearer
// @Api(tags="发布")
// @Success 200 {object} response.Postnote_resp
// @Failure 200 {object} response.Postnote_resp
// @Router /api/post/postnote [post]
func PostNote(c *gin.Context, db *gorm.DB) {

	stuid, yn := c.Get("stuid")
	posterid := stuid.(string)
	student, _ := model.FindStudfromID(posterid, db)
	text := c.PostForm("text")
	title := c.PostForm("title")
	postid := generateID.GeneratePostID(db)
	wherestring, _ := c.GetQuery("where")
	whereint, _ := strconv.Atoi(wherestring)
	key1, _ := c.GetQuery("key1")
	url1 := qiniu.GenerateURL(key1)

	var post = activity.Post{
		PostID:       postid,
		Poster:       posterid,
		Text:         text,
		PostLocation: whereint,
		Time:         time.Now().Format("2006-01-02 15:04:05"),
		HeadImage:    student.HeadImage,
		Image1:       url1,
		Title:        title,
	}

	db.Create(&post)
	db.Model(&model.Student{StuID: postid}).Updates(&model.Student{PostNumber: student.PostNumber + 1})

	if yn {
		response.Postnote_ok(c)
	} else {
		response.Postnote_fail(c)
	}
}
