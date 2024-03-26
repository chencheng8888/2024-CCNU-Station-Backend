package common

const (
	SUCCESS               = 1000 //成功
	FAIL                  = 1001 //失败
	UpdatePasswordSuccess = 201
	NotExistInentifier    = 202
	ERROR                 = 500
	InvalidParams         = 400
	ErrorDatabase         = 40001

	WebsocketSuccessMessage = 50001
	WebsocketSuccess        = 50002 //websocket成功注册
	WebsocketEnd            = 50003
	WebsocketOnlineReply    = 50004 //对方在线应答
	WebsocketOfflineReply   = 50005 //对方不在线应答
	WebsocketLimit          = 50006
	ErrDataformat           = 50007 //数据格式不正确
	ReadMsgSUCCESS          = 50008 //读取消息成功
	GroupmsgSUCCESS         = 50009 //群聊消息发送成功
)

var MsgFlags = map[int]string{
	SUCCESS:                 "ok",
	NotExistInentifier:      "该第三方账号未绑定",
	ERROR:                   "fail",
	InvalidParams:           "请求参数错误",
	ErrorDatabase:           "数据库操作出错,请重试",
	WebsocketSuccessMessage: "解析content内容信息",
	WebsocketSuccess:        "发送信息，请求历史纪录操作成功",
	WebsocketEnd:            "请求历史纪录，但没有更多记录了",
	WebsocketOnlineReply:    "针对回复信息在线应答成功",
	WebsocketOfflineReply:   "针对回复信息离线回答成功",
	WebsocketLimit:          "请求收到限制",
	ErrDataformat:           "数据格式不正确",
	ReadMsgSUCCESS:          "读取消息成功",
	GroupmsgSUCCESS:         "群聊消息发送成功",
}
