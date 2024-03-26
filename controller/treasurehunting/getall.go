package treasurehunting

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"guizizhan/model/activity"
	response "guizizhan/response/treasurehunting"
	"strconv"
)

// GetAllTreasureHuntings 获取特定用户发布的所有寻宝活动的接口。
// @Summary 获取所有寻宝活动
// @Description poster是指发布人的ID，treasureID是指寻宝活动的ID，image是物品的图片，thing是要寻找的物品，text是发布寻宝活动的文本内容，time是发布的时间，treasurelocation是寻宝活动的具体地点。只有当YNLogin=false,code才会是FAIL即1001，其他时候code为SUCCESS即1000。注意返回的是包含寻宝活动信息的数组。
// @ID get-all-treasure-huntings
// @Produce json
// @Param where query string true "发布的地点"
// @Security Bearer
// @Api(tags="获取")
// @Success 200 {object} response.GetTreasurehuntings_resp
// @Failure 200 {object} response.GetTreasurehuntings_resp
// @Router /api/getactivity/alltreasurehunting [get]
func GetAllTreasureHuntings(c *gin.Context, db *gorm.DB) {
	var msg string
	wherestring, _ := c.GetQuery("where")
	whereint, _ := strconv.Atoi(wherestring)
	var TreasureHuntings []activity.Treasurehunting
	res := db.Model(&activity.Treasurehunting{}).Where(&activity.Treasurehunting{Treasurelocation: whereint}).Order("time desc").Find(&TreasureHuntings)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		msg = "这个人没有发布寻宝活动"
	} else {
		msg = "找到了"
	}

	response.GetTreasurehuntings_ok(c, TreasureHuntings, msg)

}
