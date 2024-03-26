package activity

// Post
type Post struct {
	PostID       string `json:"postID" gorm:"primaryKey"` //帖子的ID
	Poster       string `json:"poster"`                   //发布人的ID
	HeadImage    string `json:"headimage"`                //发布人的头像
	PostLocation int    `json:"postLocation"`             //发布地点
	Title        string `json:"title"`                    //帖子的标题
	Text         string `json:"text"`                     //帖子的内容
	Time         string `json:"time"`                     //发布时间
	Image1       string `json:"image1"`                   //帖子的图片
}
