package chathandler

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"guizizhan/model/chat"
	"time"
)

var Ctx context.Context

func StoreSingleMsg(db *gorm.DB, id string, msg string) {
	var sendid, recipientid string
	fmt.Sscanf(id, "%s->%s", &sendid, &recipientid)
	var chatmsg = chat.ChatMessage{
		SendID:      sendid,
		RecipientID: recipientid,
		Content:     msg,
		SendTime:    time.Now().Format("2006-01-02 15:04:05"),
	}
	db.Create(&chatmsg)
}
func StoreGroupMsg(db *gorm.DB, id string, groupid string, msg string) {
	var chatmsg = chat.ChatMessage{
		SendID:   id,
		GroupID:  groupid,
		Content:  msg,
		SendTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	db.Create(&chatmsg)
}

func FindSingleMsg(db *gorm.DB, id, recipientid, time string) []chat.ChatMessage {
	var msgs []chat.ChatMessage
	db.Model(&chat.ChatMessage{}).Where("((id = ? AND recipient = ?) OR (id = ? AND recipient = ?)) AND sendtime < ? ", id, recipientid, recipientid, id, time).Order("sendtime asc").Limit(20).Find(&msgs)
	return msgs
}
func FindPublicMsg(db *gorm.DB, groupid string, time string) []chat.ChatMessage {
	var msgs []chat.ChatMessage
	db.Model(&chat.ChatMessage{}).Where("groupid = ? AND sendtime < ? ", groupid, time).Order("sendtime asc").Limit(20).Find(&msgs)
	return msgs
}

//func Cache(id string, msg string) error {
//	var sendid, recipientid string
//	fmt.Sscanf(id, "%s->%s", &sendid, &recipientid)
//	var chatmsg = chat.ChatMessage{
//		SendID:      sendid,
//		RecipientID: recipientid,
//		Content:     msg,
//		SendTime:    time.Now(),
//	}
//	ID1 := fmt.Sprintf("%s:%s", sendid, recipientid)
//	ID2 := fmt.Sprintf("%s:%s", recipientid, sendid)
//	msgbytes, _ := json.Marshal(chatmsg)
//	myredis.Rdb.SetEX(Ctx)
//
//}
