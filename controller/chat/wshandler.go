package chathandler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"guizizhan/model/chat"
	"net/http"
)

var Buildings = [...]string{"1", "2", "3", "4", "5", "6", "7", "8"}

// WsHandler 建立websocket连接
// @Summary 建立websocket连接
// @Description 这个只是建立websocket连接，后续的发消息是根据websocket连接来发送的
// @Produce json
// @Param uid query string true "自己的ID"
// @Param touid query string true "对方的ID或者群的ID"
// @Security Bearer
// @Router /api/talk [get]
func WsHandler(c *gin.Context, db *gorm.DB) {
	uid := c.Query("uid")
	touid := c.Query("touid")
	//chat_type := c.Query("type")

	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	//创建一个用户实例
	client := new(chat.Client)
	if IsGroup(touid) {
		client = &chat.Client{
			ID:      uid,
			GroupID: touid,
			SendID:  "",
			Socket:  conn,
			Send:    make(chan []byte),
		}
	} else {
		client = &chat.Client{
			ID:      Createid(uid, touid),
			SendID:  Createid(touid, uid),
			GroupID: "",
			Socket:  conn,
			Send:    make(chan []byte),
		}
	}

	chat.Manager.Register <- client

	//开两个协程用于读写消息
	go Read(client, db)
	go Write(client)
}

func Createid(uid, touid string) string {
	return uid + "->" + touid
}

func IsGroup(touid string) bool {
	for _, v := range Buildings {
		if v == touid {
			return true
		}
	}
	return false
}
