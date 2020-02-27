package wechat

import (
	"math"
	"encoding/json"
)

const CLIENT_MESSAGE  = "clientMessage" //客服发送
const WECHAY_MESSAGE  = "wechatMessage" //微信用户发送
const abortIndex int8 = math.MaxInt8 / 2
//发送消息信息接收体
type Send_Message struct {
	Action string
	Data User_Message
}
//消息发送用户信息结构
type User_Message struct {
	Avatar,Content,Username,Type,Fromusername,Appid,To_Username,To_Avatar string
	Unread int
	Mine,NewUser,Jion bool
	Id,To_id,Fd int64
}

type Handlers_Chain []Handler_Func

type Handler_Func func(*Conn_Service) []byte

//客服结构体
type Conn_Service struct {
	Fd int64 //连接fd
	RequestByte []byte
	Mine Send_Message //消息回调
	Handlers Handlers_Chain //消息回调
	index   int8
	ReturnByte []byte
}

//客服处理对象
func New_Service() *Conn_Service {
	return &Conn_Service{}
}

//Next::中间件
func (self *Conn_Service) Next() []byte {
	//self.index
	for self.index < int8(len(self.Handlers)) {
		self.ReturnByte = self.Handlers[self.index](self)
		self.index++
	}
	return self.ReturnByte
}

//提前结束请求处理
func (self *Conn_Service) Abort() {
	self.index = abortIndex
}

//加载信息
func (self *Conn_Service) Reset(data []byte)  {
	self.RequestByte = data
	json.Unmarshal(data,&self.Mine)
}

//赋值动作信息
func (self *Conn_Service) ResetData(action string,data []byte) {
	json.Unmarshal(data,&self.Mine.Data)
	self.Mine.Action = action
}

