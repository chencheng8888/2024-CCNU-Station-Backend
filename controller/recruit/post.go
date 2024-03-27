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
// @Param where formData string true "招募地点"
// @Param request formData string true "招募要求"
// @Param title formData string true "活动名称"
// @Param time formData string true "活动时间"
// @Security Bearer
// @Api(tags="发布")
// @Success 200 {object} response.Postrecruit_resp
// @Failure 200 {object} response.Postrecruit_resp
// @Router /api/post/post_recruit_activity [post]
func PostRecruit(c *gin.Context, db *gorm.DB) {
	posterid, yn := tool.GetStudentID(c)
	student, _ := model.FindStudfromID(posterid, db)

	wherestring := c.PostForm("where")
	whereint, _ := strconv.Atoi(wherestring)
	request := c.PostForm("request")
	title := c.PostForm("title")
	acitivity_time := c.PostForm("time")

	recruitid := generateID.GenerateRecruitID(db)

	var recruit = activity.Recruit{
		RecruitID:    recruitid,
		Poster:       posterid,
		HeadImage:    student.HeadImage,
		Where:        whereint,
		Request:      request,
		Title:        title,
		PostTime:     time.Now().Format("2006-01-02 15:04:05"),
		ActivityTime: acitivity_time,
	}

	db.Create(&recruit)

	if yn {
		response.Postrecruit_ok(c)
	} else {
		response.Postrecruit_fail(c)
	}
}
