package generateID

import (
	"errors"
	"gorm.io/gorm"
	"guizizhan/model/activity"
	"math/rand"
	"time"
)

func GenerateNumericID(length int) string {
	// 种子值
	rand.Seed(time.Now().UnixNano())

	// 数字字符集
	numbers := "0123456789"

	// 生成ID
	id := make([]byte, length)
	for i := range id {
		id[i] = numbers[rand.Intn(len(numbers))]
	}

	return string(id)
}

func GeneratePostID(db *gorm.DB) string {
	var PostID string
	for {
		PostID = GenerateNumericID(10)
		res := db.Model(&activity.Post{}).First(&activity.Post{PostID: PostID})
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			break
		}
	}
	return PostID
}

func GenerateRecruitID(db *gorm.DB) string {
	var PostID string
	for {
		PostID = GenerateNumericID(10)
		res := db.Model(&activity.Recruit{}).First(&activity.Recruit{RecruitID: PostID})
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			break
		}
	}
	return PostID
}

func GenerateTreasureID(db *gorm.DB) string {
	var PostID string
	for {
		PostID = GenerateNumericID(10)
		res := db.Model(&activity.Treasurehunting{}).First(&activity.Treasurehunting{TreasureID: PostID})
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			break
		}
	}
	return PostID
}
