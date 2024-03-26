package chat

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type ChatMessage struct {
	gorm.Model
	SendID      string //发送者的ID
	RecipientID string //接收者的ID
	GroupID     string //群ID，该消息要发到哪个群里面
	Content     string //内容
	SendTime    string //发送的时间
}

type Group struct {
	ID        int    `gorm:"primaryKey"` //群ID
	Groupname string `json:"group_name"`
}
type SendMsg struct {
	Type        int    `json:"type"`         //1 代表单聊消息 , 2 代表群聊消息 ,3 代表获取单聊的历史消息 , 4 代表获取群聊的历史消息
	SendID      string `json:"send_id"`      //发送消息的ID
	RecipientID string `json:"recipient_id"` //获取消息的用户ID或者群ID
	Content     string `json:"content"`      //发送内容
	Time        string `json:"time"`         //这个只在获取历史消息的时候才去填写，代表你想获取哪个时间点前的历史消息
}
type ReplyMsg struct {
	Code    int    `json:"code"`
	Content string `json:"content"`
	From    string `json:"from"`
	Time    string `json:"time"`
}
type Client struct {
	ID      string
	SendID  string
	GroupID string
	Socket  *websocket.Conn
	Send    chan []byte
}
type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}
type ClientManager struct {
	GroupBasic     map[string]*Client
	Clients        map[string]*Client
	Broadcast      chan *Broadcast
	GroupBroadcast chan *Broadcast
	Reply          chan *Client
	Register       chan *Client
	Unregister     chan *Client
}

var Manager = ClientManager{
	GroupBasic:     make(map[string]*Client),
	Clients:        make(map[string]*Client),
	Broadcast:      make(chan *Broadcast),
	GroupBroadcast: make(chan *Broadcast),
	Reply:          make(chan *Client),
	Register:       make(chan *Client),
	Unregister:     make(chan *Client),
}

func SendReplyMsg(client *Client, replymsg ReplyMsg) {
	msg, _ := json.Marshal(replymsg)
	client.Socket.WriteMessage(websocket.TextMessage, msg)
}
func CreateReplymsg(code int, from string, content string, time string) ReplyMsg {
	return ReplyMsg{
		Code:    code,
		From:    from,
		Content: content,
		Time:    time, //time.Now().Format("2006-01-02 15:04:05"),
	}
}
