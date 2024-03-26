package mysql

import (
	"errors"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"guizizhan/model"
	"guizizhan/model/activity"
	"guizizhan/model/chat"
	"log"
	"os"
	"time"
)

func InitMySQL() (*gorm.DB, error) {
	//自定义日志，打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)

	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{Logger: newLogger})
	if err != nil {
		err = errors.New("连接数据库失败")
		return nil, err
	}

	//自动迁移
	err = db.AutoMigrate(&model.Student{},
		&model.Traveler{},
		&model.Building{},
		&model.Contact{}, &activity.Post{},
		&activity.Treasurehunting{},
		&activity.Recruit{},
		&model.Achievement{},
		&chat.ChatMessage{})

	if err != nil {
		err = errors.New("自动迁移失败")
		return nil, err
	}
	return db, nil
}
