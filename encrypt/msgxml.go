package encrypt

import "encoding/xml"

const (
	TEXT = "text"
	IMAGE = "image"
	VOICE = "voice"
	VIDEO = "video"
	LINK = "link"
	EVENT = "event"
	LOCATION = "location"
	ENTER_AGENT = "enter_agent"
	SUBSCRIBE = "subscribe"
	ADD_EXTTERNAL_CONTACT = "add_external_contact"
	ADD_HALF_EXTERNAL_CONTACT = "add_half_external_contact"
	DEL_EXTERNAL_CONTACT = "del_external_contact"
	DEL_FOLLOW_USER = "del_follow_user"
)
type MsgType struct {
	XMLName xml.Name `xml:"xml"`
	MsgType string  `xml:"MsgType"`
	ChangeType string `xml:"ChangeType"`
}
type Base struct {
	FromUserName string `xml:"FromUserName"`
	ToUserName   string `xml:"ToUserName"`
	MsgType      string `xml:"MsgType"`
	CreateTime   int64 `xml:"CreateTime"`
}

//添加企业客户事件
type AddExternalContact struct{
	XMLName xml.Name `xml:"xml"`
	Base
	Event string `xml:"Event"`
	ChangeType string `xml:"ChangeType"`
	UserID string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
	State string `xml:"State"`
	WelcomeCode string `xml:"WelcomeCode"`
}
//外部联系人免验证添加成员事件
type AddHalfExternalContact struct{
	XMLName xml.Name `xml:"xml"`
	Base
	Event string `xml:"Event"`
	ChangeType string `xml:"ChangeType"`
	UserID string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
	State string `xml:"State"`
	WelcomeCode string `xml:"WelcomeCode"`
}
//删除企业客户事件
type DelExternalContact struct{
	XMLName xml.Name `xml:"xml"`
	Base
	Event string `xml:"Event"`
	ChangeType string `xml:"ChangeType"`
	UserID string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
}
//删除跟进成员事件
type DelFollowUser struct{
	XMLName xml.Name `xml:"xml"`
	Base
	Event string `xml:"Event"`
	ChangeType string `xml:"ChangeType"`
	UserID string `xml:"UserID"`
	ExternalUserID string `xml:"ExternalUserID"`
}
// -------------------------------- 接收应用消息 -------------------------------- //
//文本消息
type TextMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	Content string `xml:"Content"`
}
//图片消息
type ImageMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	PicUrl  string `xml:"PicUrl"`
	MediaId  string `xml:"MediaId"`
}

//语音消息
type VoiceMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	MediaId string `xml:"MediaId"`
	Format string `xml:"Format"`
}
//视频消息
type VideoMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	MediaId string `xml:"MediaId"`
	ThumbMediaId string `xml:"ThumbMediaId"`
}
//位置消息
type LocationMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	Scale int64 `xml:"Scale"`
	Location_X string `xml:"Location_X"`
	Location_Y string `xml:"Location_Y"`
	Label string `xml:"Label"`
}
//链接消息
type LinkMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	Title string `xml:"Title"`
	Description string `xml:"Description"`
	Url string `xml:"Url"`
	PicUrl string `xml:"PicUrl"`
}
// -------------------------------- 事件消息-------------------------------- //

//成员关注及取消关注事件
type SubscribeMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	EventKey string `xml:"EventKey"`
	Event string `xml:"Event"`
}
//点击小程序按钮
type TempsessionMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	SessionFrom string `xml:"SessionFrom"`
}

//小程序卡片
type MiniprogrampagenMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	Title string		`xml:"Title"`		//发送小程序卡片标题
	PagePath  string   	`xml:"PagePath"`    //发送小程序卡片路径
	ThumbUrl  string    	`xml:"ThumbUrl"`	//发送小程序图片url
}

//微信二维码
type QrcodeMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	EventKey string `xml:"EventKey"`
	Event string `xml:"Event"`
}
//进入应用
type Enter_AgentMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	Event string `xml:"Event"`
	EventKey string `xml:"Event"`
}
// -------------------------------- 被动回复消息格式 -------------------------------- //
type CDATAText struct {
	Text string `xml:",innerxml"`
}
type Passivity struct {
	ToUserName   CDATAText
	FromUserName CDATAText
	MsgType      CDATAText
	CreateTime   int64
}

//文本消息
type PassivityText struct {
	XMLName xml.Name `xml:"xml"`
	Passivity
	Content CDATAText
}
//图片消息
type PassivityImage struct {
	XMLName xml.Name `xml:"xml"`
	Passivity
	Image struct{
		MediaId CDATAText
	}
}

//语音消息
type PassivityVoice struct {
	XMLName xml.Name `xml:"xml"`
	Passivity
	Voice struct{
		MediaId CDATAText
	}
}
//视频消息
type PassivityVideo struct {
	XMLName xml.Name `xml:"xml"`
	Passivity
	Video struct{
		MediaId CDATAText
		Title CDATAText
		Description CDATAText
	}
}
func Value2CDATA(v string) CDATAText {
	return CDATAText{"<![CDATA[" + v + "]]>"}
}


//微信客服接入结构体
type PassivityCustom struct {
	XMLName xml.Name  `xml:"xml"`
	Passivity
}
