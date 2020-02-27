package firm

import (
	"math"
	"encoding/json"
	"encoding/xml"
	"./wechat"
	"sync"
	"fmt"
)

type HandlersChain []HandlerFunc

type HandlerFunc func(wechat *wechat.Wechat)

const abortIndex int8 = math.MaxInt8 / 2

//消息类型
type TypeEvent struct {
	MsgType,Event string
}
type WechatCallback struct {
	envet []TypeInfo //消息回调
	handlers HandlersChain //消息回调
	pool  sync.Pool
	Format string
	index   int8
}

type TypeInfo struct {
	MsgType      string
	HandlerFunc HandlerFunc
}

func NewWechatCallback() *WechatCallback {
	Callback := &WechatCallback{}
	Callback.pool.New = func() interface{} {
		return Callback.allocateContext()
	}
	return Callback
}
// allocateContext
func (self *WechatCallback) allocateContext() wechat.Wechat {
	// 构造新的上下文对象
	return wechat.Wechat{}
}
//Next::中间件
func (self *WechatCallback) Next(c *wechat.Wechat) {
	//c.index
	for self.index < int8(len(self.handlers)) {
		self.handlers[self.index](c)
		self.index++
	}
}
//提前结束请求处理
func (c *WechatCallback) Abort() {
	c.index = abortIndex
}

//添加插件
func (self *WechatCallback) Use(middleware ...HandlerFunc) {
	self.handlers = append(self.handlers, middleware...)
	return
}
//添加消息监听
func (self *WechatCallback) PushMessage(msg_type string,function HandlerFunc) {
	self.envet = append(self.envet,TypeInfo{
		MsgType:msg_type,HandlerFunc:function,
	})
}

//加载消息类型
func (self *WechatCallback) LoadMessage(data []byte,wechat2 *wechat.Wechat,Format string)  {
	wechat2.Format = Format
	wechat2.MsgType = self.LoadEvent(data ,Format)
	return
}

//消息类型
func (self *WechatCallback) LoadEvent(data []byte,Format string) string {
	var wechat2 TypeEvent
	if Format == "json"{
		json.Unmarshal(data,&wechat2)
	}else {
		xml.Unmarshal(data,&wechat2)
	}
	fmt.Println("TypeEvent",wechat2)
	if wechat2.Event != ""{
		return fmt.Sprintf("%v/%v",wechat2.MsgType,wechat2.Event)
	}
	return wechat2.MsgType
}
//开始运行
func (self *WechatCallback) Start(wec *wechat.Wechat) {
	for _,function := range self.envet {
		if function.MsgType == wec.MsgType {
			self.handlers = append(self.handlers,function.HandlerFunc)
			self.Next(wec)
		}
	}
	return
}
//线程池处理
func (self *WechatCallback) StartPool(data []byte,Format string) {
	c := self.pool.Get().(*wechat.Wechat)
	self.LoadMessage(data,c,Format)
	self.Start(c)
	self.pool.Put(c)
}