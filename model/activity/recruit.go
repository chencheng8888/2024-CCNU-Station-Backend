package activity

// Recruit
type Recruit struct { //招募活动
	RecruitID    string `json:"recruitID" gorm:"primaryKey"` //招募活动的ID
	Poster       string `json:"posterid"`                    //发布人的ID
	HeadImage    string `json:"headimage"`                   //发布人的头像
	PostTime     string `json:"post-time"`                   //发布时间
	Where        string `json:"where"`                       //活动地点
	Request      string `json:"request"`                     //活动要求
	Title        string `json:"title"`                       //活动名称
	ActivityTime string `json:"activity-time"`               //活动时间
}
