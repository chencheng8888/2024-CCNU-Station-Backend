package activity

// Recruit
type Recruit struct { //招募活动
	RecruitID string `json:"recruitID" gorm:"primaryKey"` //招募活动的ID
	Poster    string `json:"posterid"`                    //发布人的ID
	HeadImage string `json:"headimage"`                   //发布人的头像
	Time      string `json:"time"`                        //发布时间
	Where     int    `json:"where"`                       //发布地点
	Request   string `json:"request"`                     //招募要求
	Text      string `json:"text"`                        //招募活动的内容
}
