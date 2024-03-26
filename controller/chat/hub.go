package chathandler

import (
	"fmt"
	"gorm.io/gorm"
	"guizizhan/common"
	"guizizhan/model/chat"
	"time"
)

func Hub(db *gorm.DB) {
	for {
		fmt.Println("<-----监听通信管道----->")
		select {
		case conn := <-chat.Manager.Register:
			if conn.SendID != "" {
				replymsg := chat.CreateReplymsg(common.WebsocketSuccess, "SYSTEM", common.MsgFlags[common.WebsocketSuccess], time.Now().Format("2006-01-02 15:04:05"))
				chat.SendReplyMsg(conn, replymsg)
				chat.Manager.Clients[conn.ID] = conn
			} else {
				replymsg := chat.CreateReplymsg(common.WebsocketSuccess, "SYSTEM", common.MsgFlags[common.WebsocketSuccess], time.Now().Format("2006-01-02 15:04:05"))
				chat.SendReplyMsg(conn, replymsg)
				chat.Manager.GroupBasic[conn.GroupID] = conn
			}

		case conn := <-chat.Manager.Unregister:
			if conn.SendID != "" {
				replymsg := chat.CreateReplymsg(common.WebsocketEnd, "SYSTEM", common.MsgFlags[common.WebsocketEnd], time.Now().Format("2006-01-02 15:04:05"))
				chat.SendReplyMsg(conn, replymsg)
				delete(chat.Manager.Clients, conn.ID)
			} else {
				replymsg := chat.CreateReplymsg(common.WebsocketEnd, "SYSTEM", common.MsgFlags[common.WebsocketEnd], time.Now().Format("2006-01-02 15:04:05"))
				chat.SendReplyMsg(conn, replymsg)
				delete(chat.Manager.GroupBasic, conn.GroupID)
			}
		case broadcast := <-chat.Manager.Broadcast:
			message := broadcast.Message
			id := broadcast.Client.ID
			recipientID := broadcast.Client.SendID
			flag := false
			for _, conn := range chat.Manager.Clients {
				if conn.ID == recipientID {
					select {
					case conn.Send <- message:
						flag = true
					default:
						close(conn.Send)
						delete(chat.Manager.Clients, conn.ID)
					}
				}
			}
			if flag {
				replymsg := chat.CreateReplymsg(common.WebsocketOnlineReply, "SYSTEM", common.MsgFlags[common.WebsocketOnlineReply], time.Now().Format("2006-01-02 15:04:05"))
				chat.SendReplyMsg(broadcast.Client, replymsg)
				StoreSingleMsg(db, id, fmt.Sprintf("%s", string(message)))
			} else {
				replymsg := chat.CreateReplymsg(common.WebsocketOfflineReply, "SYSTEM", common.MsgFlags[common.WebsocketOfflineReply], time.Now().Format("2006-01-02 15:04:05"))
				chat.SendReplyMsg(broadcast.Client, replymsg)
				StoreSingleMsg(db, id, fmt.Sprintf("%s", string(message)))
			}
		case broadcast := <-chat.Manager.GroupBroadcast:
			message := broadcast.Message
			id := broadcast.Client.ID
			recipientID := broadcast.Client.GroupID
			for _, conn := range chat.Manager.GroupBasic {
				if conn.GroupID == recipientID && conn.ID != id {
					select {
					case conn.Send <- message:
						replymsg := chat.CreateReplymsg(common.WebsocketOnlineReply, "SYSTEM", common.MsgFlags[common.WebsocketOnlineReply], time.Now().Format("2006-01-02 15:04:05"))
						chat.SendReplyMsg(broadcast.Client, replymsg)
						StoreGroupMsg(db, id, recipientID, fmt.Sprintf("%s", string(message)))
					}
				}
			}
		}
	}
}
