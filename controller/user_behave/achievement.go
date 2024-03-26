package user_behave

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"guizizhan/model"
	"guizizhan/response"
	"strconv"
)

// GetAchievements  获取所有成就的完成情况
// @Summary 获取所有成就的完成情况
// @Description 返回的finished是由100个由'0'(代表未完成)和'1'(代表已完成)组成的字符串，其中第i位的字符的意义是第i+1个成就的完成情况
// @Produce json
// @Param stuid query string true "学生的ID"
// @Success 200 {object} response.Getachi_resp
// @Failure 200 {object} response.Getachi_resp
// @Router /api/user/achievement/get [get]
func GetAchievements(c *gin.Context, db *gorm.DB) {
	stuid := c.Query("stuid")
	finished, err := model.FindAchievement(stuid, db)
	if err != nil {
		response.Getachi_fail(c)
	} else {
		response.Getachi_ok(c, finished)
	}
}

// UpdateAchievement  更改成就的状态
// @Summary 更改成就的状态
// @Description 这个更新是取反，即如果为1就改为0，为0就改为1
// @Param stuid query string true "学生的ID"
// @Param achid query string true "要改的第几个成就"
// @Router /api/user/achievement/update [get]
func UpdateAchievement(c *gin.Context, db *gorm.DB) {
	stuid := c.Query("stuid")
	AchIDStr := c.Query("achid")
	Achid, _ := strconv.Atoi(AchIDStr)
	model.UpdateAchievement(stuid, Achid, db)
	response.Updateachi_ok(c)
}
