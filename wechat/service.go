package wechat
//
//import (
//	"encoding/json"
//	"fmt"
//)
//
//const CLIENT_MESSAGE  = "clientMessage" //客服发送
//const WECHAY_MESSAGE  = "wechatMessage" //微信用户发送
//type HandlerFuncs func(string,[]byte)  []byte //其他结构
//type Message func(SendMessage) []byte //触发消息信息
//type Service struct {
//	Fd int64 //连接fd
//	Message Message //消息回调
//	MessageByte HandlerFuncs//消息回调
//
//}
//
////请求动作
//type clientAction struct {
//	Action string //请求动作
//}
////客服处理对象
//func NewService() Service {
//	return Service{}
//}
////发送消息信息接收体
//type SendMessage struct {
//	Action string
//	Data UserMessage
//}
////消息发送用户信息结构
//type UserMessage struct {
//	Avatar,Content,Username,Type,Fromusername,Appid string
//	Unread int
//	Mine,IsNew bool
//	Id,To int64
//}
//
////处理消息返回
//func (self *Service) Connect(data []byte) []byte {
//	var action clientAction
//	self.LoadMessage(data,&action)
//	if action.Action == CLIENT_MESSAGE || action.Action == WECHAY_MESSAGE {
//		var messaage SendMessage
//		self.LoadMessage(data,&messaage)
//		return self.Message(messaage)
//	}
//	return self.MessageByte(action.Action,data)
//}
//
////监听连接事件
//func (self *Service) Open(addr int64 ,openFunc func(i int64))  {
//	self.Fd = addr
//	fmt.Println("fd是:",addr)
//	openFunc(addr)
//}
////监听消息接收
//func (self *Service) OnMessage(Function Message) {
//	self.Message = Function
//}
////监听消息接收
//func (self *Service) OnMessageByte(Function HandlerFuncs) {
//	self.MessageByte = Function
//	return
//}
////监听消息接收
//func (self *Service) OnClose(Function Message) {
//	self.Message = Function
//	return
//}
////载入消息
//func (self *Service) LoadMessage(data_byte []byte,data interface{}) error {
//	err := json.Unmarshal(data_byte,&data)
//	return err
//}