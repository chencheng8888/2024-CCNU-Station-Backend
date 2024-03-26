package chathandler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"guizizhan/common"
	"guizizhan/model/chat"
	"time"
)

func Read(client *chat.Client, db *gorm.DB) {
	defer func() { // 避免忘记关闭，所以要加上close
		chat.Manager.Unregister <- client
		_ = client.Socket.Close()
	}()
	for {
		client.Socket.PongHandler()
		sendmsg := new(chat.SendMsg)
		err := client.Socket.ReadJSON(&sendmsg)
		if err != nil {
			fmt.Println("数据格式不正确")
			replymsg := chat.CreateReplymsg(common.ErrDataformat, "SYSTEM", common.MsgFlags[common.ErrDataformat], time.Now().Format("2006-01-02 15:04:05"))
			chat.SendReplyMsg(client, replymsg)
			continue
		}
		time := sendmsg.Time
		switch sendmsg.Type {
		case 1:
			private_chat(client, sendmsg.Content)
		case 2:
			public_chat(client, sendmsg.Content)
		case 3:
			send_singlehistory(db, client, time)
		case 4:
			send_publichistory(db, client, time)
		}
	}
}

func Write(client *chat.Client) {
	defer func() {
		_ = client.Socket.Close()
	}()
	for {
		select {
		//读取管道里面的消息
		case message, ok := <-client.Send:
			//连接不到就返回消息
			if !ok {
				_ = client.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			replymsg := chat.CreateReplymsg(common.ReadMsgSUCCESS, client.ID, fmt.Sprintf("%s", string(message)), time.Now().Format("2006-01-02 15:04:05"))
			chat.SendReplyMsg(client, replymsg)
		}
	}
}
func private_chat(client *chat.Client, content string) {
	chat.Manager.Broadcast <- &chat.Broadcast{
		Client:  client,
		Message: []byte(content),
		Type:    1,
	}
	replymsg := chat.CreateReplymsg(common.WebsocketSuccessMessage, "SYSTEM", common.MsgFlags[common.WebsocketSuccessMessage], time.Now().Format("2006-01-02 15:04:05"))
	chat.SendReplyMsg(client, replymsg)
}
func public_chat(client *chat.Client, content string) {
	chat.Manager.GroupBroadcast <- &chat.Broadcast{
		Client:  client,
		Message: []byte(content),
		Type:    2,
	}
	replymsg := chat.CreateReplymsg(common.WebsocketSuccessMessage, "SYSTEM", common.MsgFlags[common.WebsocketSuccessMessage], time.Now().Format("2006-01-02 15:04:05"))
	chat.SendReplyMsg(client, replymsg)
}
func send_singlehistory(db *gorm.DB, client *chat.Client, time string) {
	var sendid, recipientid string
	fmt.Sscanf(client.ID, "%s->%s", &sendid, &recipientid)
	msgs := FindSingleMsg(db, sendid, recipientid, time)
	for _, msg := range msgs {
		replymsg := chat.CreateReplymsg(common.WebsocketSuccessMessage, msg.SendID, msg.Content, msg.SendTime)
		chat.SendReplyMsg(client, replymsg)
	}
}
func send_publichistory(db *gorm.DB, client *chat.Client, time string) {
	groupid := client.GroupID
	msgs := FindPublicMsg(db, groupid, time)
	for _, msg := range msgs {
		replymsg := chat.CreateReplymsg(common.WebsocketSuccessMessage, msg.SendID, msg.Content, msg.SendTime)
		chat.SendReplyMsg(client, replymsg)
	}
}
