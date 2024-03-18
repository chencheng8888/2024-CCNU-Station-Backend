package chathandler

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"guizizhan/model/chat"
	"time"
)

var Ctx context.Context

func StoreMsg(db *gorm.DB, id string, msg string) {
	var sendid, recipientid string
	fmt.Sscanf(id, "%s->%s", &sendid, &recipientid)
	var chatmsg = chat.ChatMessage{
		SendID:      sendid,
		RecipientID: recipientid,
		Content:     msg,
		SendTime:    time.Now(),
	}
	db.Create(&chatmsg)
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
