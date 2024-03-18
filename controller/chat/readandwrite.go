package chathandler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"guizizhan/common"
	"guizizhan/model/chat"
)

func Read(client *chat.Client) {
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
			replymsg := chat.CreateReplymsg(common.ErrDataformat, "SYSTEM", common.MsgFlags[common.ErrDataformat])
			chat.SendReplyMsg(client, replymsg)
			continue
		}
		switch sendmsg.Type {
			case 1:
				private_chat(client, sendmsg.Content)

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
			replymsg := chat.CreateReplymsg(common.ReadMsgSUCCESS, client.SendID, fmt.Sprintf("%s", string(message)))
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
	replymsg := chat.CreateReplymsg(common.WebsocketSuccessMessage, "SYSTEM", common.MsgFlags[common.WebsocketSuccessMessage])
	chat.SendReplyMsg(client, replymsg)
}
