package activity

// Treasurehunting
type Treasurehunting struct {
	TreasureID       string `json:"treasureID" gorm:"primaryKey"`           //寻宝活动的ID
	Poster           string `gorm:"foreignKey:Student.StuID" json:"poster"` //发布人的ID
	HeadImage        string `json:"headImage"`                              //发布人的头像
	Title            string `json:"title"`                                  //活动标题
	Deadline         string `json:"deadline"`                               //截止日期
	Clue             string `json:"clue"`                                   //线索
	Treasurelocation int    `json:"treasurelocation"`                       //寻宝活动的地点
	Thing            string `json:"thing"`                                  //要寻找的物品
	Time             string `json:"time"`                                   //发布时间
	Image            string `json:"image"`                                  //物品图片
}
