package recruit

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"guizizhan/model"
	response "guizizhan/response/recruit"
	"guizizhan/service/tool"
	"strconv"
)

// GetAllRecruits 获取特定用户发布的所有招募活动的接口。
// @Summary 获取所有招募活动
// @Description posterid是指发布人的ID，recruitID是指招募活动的ID，request是招募的要求，text是招募活动的文本内容，time是发布的时间，where是招募活动的具体地点。只有当YNLogin=false,code才会是FAIL即1001，其他时候code为SUCCESS即1000。注意返回的是包含招募活动信息的数组。
// @ID get-all-recruits
// @Produce json
// @Param where query string true "发布的地点"
// @Security Bearer
// @Api(tags="获取")
// @Success 200 {object} response.GetRecruits_resp
// @Failure 200 {object} response.GetRecruits_resp
// @Router /api/getactivity/allrecruit [get]
func GetAllRecruits(c *gin.Context, db *gorm.DB) {
	var msg string
	_, yn := tool.GetStudentID(c)

	wherestring, _ := c.GetQuery("where")
	whereint, _ := strconv.Atoi(wherestring)

	var recruits []model.Recruit
	res := db.Model(&model.Recruit{}).Where(&model.Recruit{Where: whereint}).Find(&recruits)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		msg = "这个人没有发布招募活动"
	} else {
		msg = "找到了"
	}
	if yn {
		response.GetRecruits_ok(c, recruits, msg)
	} else {
		response.GetRecruits_fail(c)
	}
}
