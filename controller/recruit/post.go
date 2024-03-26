package recruit

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"guizizhan/model"
	"guizizhan/model/activity"
	response "guizizhan/response/recruit"
	"guizizhan/service/generateID"
	"guizizhan/service/tool"
	"strconv"
	"time"
)

// PostRecruit 处理发布招募活动的请求。
// @Summary 发布招募活动
// @Description 返回的信息比较简单，code还是1000表示成功，1000表示失败（未登录），YNLogin代表是否登录，不过code信息已经说明了。
// @ID post-recruit
// @Accept json
// @Produce json
// @Param where query string true "招募地点"
// @Param request formData string true "招募要求"
// @Param text formData string true "招募详情"
// @Security Bearer
// @Api(tags="发布")
// @Success 200 {object} response.Postrecruit_resp
// @Failure 200 {object} response.Postrecruit_resp
// @Router /api/post/post_recruit_activity [post]
func PostRecruit(c *gin.Context, db *gorm.DB) {
	posterid, yn := tool.GetStudentID(c)
	student, _ := model.FindStudfromID(posterid, db)

	wherestring, _ := c.GetQuery("where")
	whereint, _ := strconv.Atoi(wherestring)

	request := c.PostForm("request")

	text := c.PostForm("text")

	recruitid := generateID.GenerateRecruitID(db)

	var recruit = activity.Recruit{
		RecruitID: recruitid,
		Poster:    posterid,
		HeadImage: student.HeadImage,
		Where:     whereint,
		Request:   request,
		Text:      text,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
	}

	db.Create(&recruit)

	if yn {
		response.Postrecruit_ok(c)
	} else {
		response.Postrecruit_fail(c)
	}
}
