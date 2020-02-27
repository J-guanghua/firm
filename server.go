package firm

import (
	"./wechat"
	"sync"
	"encoding/json"
	"fmt"
)
//消息返回信息结构
type ReturnMessage struct {
	Action string    `json:"action"` //动作
	Data struct{
		Username string     `json:"username"` //用户名称
		Avatar string       `json:"avatar"` //用户头像
		Fromusername string  `json:"fromusername"` //openid
		Id int64            `json:"id"` //id
		Type string         `json:"type"` //类型
		Content string      `json:"content"` //内容
		Mine bool           `json:"mine"`
		Appid string        `json:"appid"` //
		Groupid int         `json:"groupid"` //组id
		Last_msg string     `json:"last_msg"` //最后的消息内容
		Unread int          `json:"unread"` //新消息数
		Fromid int64        `json:"fromid"` //发给的用户id
		Timestamp int64       `json:"timestamp"` //发送时间
	} `json:"data"`
}
//数据接口实现
type SqlModels interface {
	ClientUserInfo(mine *wechat.User_Message) error //检查用户信息
	WechatUserInfo(mine *wechat.User_Message) error //检查用户信息
	ClientMessage (mine *wechat.User_Message) []byte //客服发送消息
	WechatMessage(mine *wechat.User_Message) []byte  //微信用户用户消息
	JoinMessage(mine *wechat.User_Message) []byte  //接入给缓冲区用户消息
	PushMessage(mine *wechat.User_Message) []byte  //接入给缓冲区用户消息
	NotifyMessage(mine *wechat.User_Message) []byte  //接入给缓冲区用户消息并通知
	FinishMessage(mine *wechat.User_Message) []byte  //结束会话@all
}
//请求动作
type ConnService struct {
	Action string
	envet []Type_Info //消息回调
	handlers wechat.Handlers_Chain
	Format string
	pool  sync.Pool
	Model SqlModels
}

type Type_Info struct {
	Action      string
	HandlerFunc wechat.Handler_Func
}
//http请求
type HttpByte struct {
	Action string `json:"action"`
	Data  struct{
		Errcode int `json:"errcode"`
		Errmsg string `json:"errmsg"`
	} `json:"data"`
}
//微信消息接口
func NewHttpByte(Action string, ErrCode int,ErrMsg string) []byte {
	var Http_Byte HttpByte
	Http_Byte.Action = Action
	Http_Byte.Data.Errcode = ErrCode
	Http_Byte.Data.Errmsg = ErrMsg
	data,_:=json.Marshal(Http_Byte)
	return data
}

//实例化
func NewConnService(model SqlModels) *ConnService {
	service := &ConnService{Model:model}
	service.pool.New = func() interface{} {
		return service.allocateContext()
	}
	service.StartAction()
	return service
}
func (self *ConnService) NewService(data []byte) wechat.Conn_Service {
	c := wechat.Conn_Service{}
	self.LoadMessage(data)
	c.Reset(data)
	return c
}
func (self *ConnService) ServiceData(action string,data []byte)(serv wechat.Conn_Service) {
	serv.ResetData(action,data)
	return serv
}
// allocateContext
func (self *ConnService) allocateContext() wechat.Conn_Service {
	// 构造新的上下文对象
	return wechat.Conn_Service{}
}
//添加插件
func (self *ConnService) Use(middleware ...wechat.Handler_Func) {
	self.handlers = append(self.handlers, middleware...)
	return
}
//添加消息监听
func (self *ConnService) PushAction(msg_type string,function wechat.Handler_Func) {
	self.envet = append(self.envet,Type_Info{
		Action:msg_type,HandlerFunc:function,
	})
}

//加载消息类型
func (self *ConnService) LoadMessage(data []byte)  {
	json.Unmarshal(data,&self)
	return
}
//加载消息类型
func (self *ConnService) ToHttpByte(data []byte)(httpcode HttpByte){
	json.Unmarshal(data,&httpcode)
	return
}
//开始运行
func (self *ConnService) Start(server wechat.Conn_Service) []byte {
	for _,function := range self.envet {
		if function.Action == server.Mine.Action {
			server.Handlers = append(self.handlers,function.HandlerFunc)
			return server.Next()
		}
	}
	return server.RequestByte
}
//线程池处理
func (self *ConnService) StartPool(data []byte) []byte {
	c := self.pool.Get().(wechat.Conn_Service)
	self.LoadMessage(data)
	c.Reset(data)
	return_data := self.Start(c)
	self.pool.Put(c)
	return return_data
}

//连接成功
func (self *ConnService) Open(addr int64,Function func(i int64))  {
	Function(addr)
}
//关闭连接
func (self *ConnService) Close(addr int64)  {

}
//设置请求动作
func (self *ConnService) StartAction()  {
	//客服发送的消息
	self.PushAction(wechat.CLIENT_MESSAGE, func(service *wechat.Conn_Service) []byte {
		if err := self.Model.ClientUserInfo(&service.Mine.Data);err != nil {
			return NewHttpByte("errMessage",4022,err.Error())
		}
		return self.Model.ClientMessage(&service.Mine.Data)
	})
	//微信用户发的消息
	self.PushAction(wechat.WECHAY_MESSAGE, func(service *wechat.Conn_Service) []byte {
		if err := self.Model.WechatUserInfo(&service.Mine.Data);err != nil {
			return NewHttpByte("errMessage",4022,err.Error())
		}
		return self.Model.WechatMessage(&service.Mine.Data)
	})

	//手动接入用户消息
	self.PushAction("joinMessage", func(service *wechat.Conn_Service) []byte {
		fmt.Println("joinMessage",service.Mine.Data)
		if err := self.Model.WechatUserInfo(&service.Mine.Data);err != nil {
			return NewHttpByte("errMessage",4022,err.Error())
		}
		return self.Model.JoinMessage(&service.Mine.Data)
	})

	//添加消息，转接消息
	self.PushAction("pushMessage", func(service *wechat.Conn_Service) []byte {
		fmt.Println("pushMessage",service.Mine.Data)
		if err := self.Model.WechatUserInfo(&service.Mine.Data);err != nil {
			return NewHttpByte("errMessage",4022,err.Error())
		}
		return self.Model.PushMessage(&service.Mine.Data)
	})
	//手动接入用户消息 ,并通知
	self.PushAction("notifyMessage", func(service *wechat.Conn_Service) []byte {
		fmt.Println("notifyMessage",service.Mine.Data)
		if err := self.Model.WechatUserInfo(&service.Mine.Data);err != nil {
			return NewHttpByte("errMessage",4022,err.Error())
		}
		return self.Model.NotifyMessage(&service.Mine.Data)
	})
	self.PushAction("finishMessage", func(service *wechat.Conn_Service) []byte {
		return self.Model.FinishMessage(&service.Mine.Data)
	})
}