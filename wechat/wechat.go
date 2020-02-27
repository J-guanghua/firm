package wechat

import (
	"../token"
	"regexp"
	"../api"
	"sort"
	"encoding/json"
	"encoding/xml"
	"../encrypt"
	"../oauth"
	"time"
	"strconv"
)

type Wechat struct {
	Wechat_id int64
	Appid string
	Secret string
	Token string
	Format string
	MsgType string
}
//实例化
func NewWechat(token,appid,secret string) *Wechat {
	return &Wechat{
		Appid:appid,
		Secret:secret,
		Token:token,
	}
}
//回调URL验证
func (self *Wechat) CheckSignature(msg_signature,timestamp,nonce string) bool {
	var tokens = self.Token
	tmps:=[]string{tokens,timestamp,nonce}
	sort.Strings(tmps)
	tmpStr:=tmps[0]+tmps[1]+tmps[2]
	tmp := token.Str2sha1(tmpStr)
	if tmp == msg_signature {
		return  true
	}
	return false
}
//获取Accesstoken
func (self *Wechat) GetAccesstoken() string {
	if ok,_:=regexp.Match("^http.*://.*", []byte(self.Secret));ok{
		AccessToken := struct {
			Access_token string
		}{}
		api.GetUnmarshal(self.Secret,&AccessToken)
		token.ContainerToken[self.Appid] = AccessToken.Access_token
		return AccessToken.Access_token
	}
	access_token,_ := token.GetToken(self.Appid,self.Secret)
	return access_token
}

//批量发消息
func (self *Wechat) BatchMessage(array []interface{}) (list []*api.CryptError) {
	for _,keyword :=range array {
		keyword := keyword.(map[string]interface{}) //接入失败回复
		list = append(list,self.SendWechat(keyword["data"],keyword["msg_type"].(string)))
	}
	return list
}

//发送客服消息
func (self *Wechat) SendCustom(data interface{}) (*api.CryptError) {
	access_token := self.GetAccesstoken()
	return self.RefreshAccesstoken(Send_Custom(data,access_token))
}

//更新hAccesstoken
func (self *Wechat) RefreshAccesstoken(errcode *api.CryptError) *api.CryptError {

	if errcode.ErrCode ==  40001 || errcode.ErrCode != 42001 {
		delete(token.ContainerToken,self.Appid)
	}
	return errcode
}
//发送模板消息
func (self *Wechat) SendTemplate(data interface{}) (*api.CryptError) {
	access_token := self.GetAccesstoken()
	return self.RefreshAccesstoken(Send_Template(data,access_token))
}

//发送微信消息
func (self *Wechat) SendWechat(data interface{},msgType string) (*api.CryptError) {
	if msgType == "template" {
		return self.SendTemplate(data)
	}
	return self.SendCustom(data)
}

//上传素材
func (self *Wechat) Material(genre string,file_path string) *MediaError {
	access_token := self.GetAccesstoken()
	return Media_Upload(access_token,genre,file_path)
}

//批量个用户打标签
func (self *Wechat) SignUsersTag(array map[string]interface{}) *api.CryptError {
	access_token := self.GetAccesstoken()
	return Sign_Users_Tag(array,access_token)
}
//微信接口获取用户信息
func (self *Wechat) QrconnectInfo(openid string) (array api.PortData,err *api.CryptError) {
	access_token := self.GetAccesstoken()
	return oauth.Qrconnect_Info(openid,access_token)
}

//返回授权url
func (self *Wechat) AuthUrl(redirect string) string {
	return oauth.Auth_Url(self.Appid,redirect)
}

//获取用户信息数据
func (self *Wechat) AuthData(code string) (array api.PortData,err *api.CryptError) {
	return oauth.Auth_Data(self.Appid,self.Secret,code)
}

//--------------------------------消息事件-------------------------------//
//用户关注
func (self *Wechat) SubscribeMessage(data []byte) encrypt.SubscribeMessage {
	var SubscribeMessage encrypt.SubscribeMessage
	if self.Format == "json" {
		json.Unmarshal(data,&SubscribeMessage)
	} else {
		xml.Unmarshal(data,&SubscribeMessage)
	}
	return SubscribeMessage
}
//文本消息
func (self *Wechat) TextMessage(data []byte) encrypt.TextMessage {
	var TextMessage encrypt.TextMessage
	if self.Format == "json" {
		json.Unmarshal(data,&TextMessage)
	} else {
		xml.Unmarshal(data,&TextMessage)
	}
	return TextMessage
}
//图片消息
func (self *Wechat) ImageMessage(data []byte) encrypt.ImageMessage {
	var ImageMessage encrypt.ImageMessage
	if self.Format == "json" {
		json.Unmarshal(data,&ImageMessage)
	} else {
		xml.Unmarshal(data,&ImageMessage)
	}
	return ImageMessage
}
//语音消息
func (self *Wechat) VideoMessage(data []byte) encrypt.VideoMessage {
	var VideoMessage encrypt.VideoMessage
	if self.Format == "json" {
		json.Unmarshal(data,&VideoMessage)
	} else {
		xml.Unmarshal(data,&VideoMessage)
	}
	return VideoMessage
}
//视频消息
func (self *Wechat) VoiceMessage(data []byte) encrypt.VoiceMessage {
	var VoiceMessage encrypt.VoiceMessage
	if self.Format == "json" {
		json.Unmarshal(data,&VoiceMessage)
	} else {
		xml.Unmarshal(data,&VoiceMessage)
	}
	return VoiceMessage
}
//地理位置消息
func (self *Wechat) LocationMessage(data []byte) encrypt.LocationMessage {
	var LocationMessage encrypt.LocationMessage
	if self.Format == "json" {
		json.Unmarshal(data,&LocationMessage)
	} else {
		xml.Unmarshal(data,&LocationMessage)
	}
	return LocationMessage
}
//链接消息
func (self *Wechat) LinkMessage(data []byte) encrypt.LinkMessage {
	var LinkMessage encrypt.LinkMessage
	if self.Format == "json" {
		json.Unmarshal(data,&LinkMessage)
	} else {
		xml.Unmarshal(data,&LinkMessage)
	}
	return LinkMessage
}
//小程序卡片
func (self *Wechat) ReplyMiniprogrampagen(data []byte) encrypt.MiniprogrampagenMessage {
	var MiniprogrampagenMessage encrypt.MiniprogrampagenMessage
	if self.Format == "json" {
		json.Unmarshal(data,&MiniprogrampagenMessage)
	} else {
		xml.Unmarshal(data,&MiniprogrampagenMessage)
	}
	return MiniprogrampagenMessage
}
//点击小程序客服按钮
func (self *Wechat) TempsessionMessage(data []byte) encrypt.TempsessionMessage {
	var TempsessionMessage encrypt.TempsessionMessage
	if self.Format == "json" {
		json.Unmarshal(data,&TempsessionMessage)
	} else {
		xml.Unmarshal(data,&TempsessionMessage)
	}
	return TempsessionMessage
}

//回复文本消息
func (self *Wechat) ReplyText(FromUserName,ToUserName,Content string) []byte {
	var PassivityText encrypt.PassivityText
	PassivityText.ToUserName =  encrypt.Value2CDATA(ToUserName)
	PassivityText.FromUserName = encrypt.Value2CDATA(FromUserName)
	PassivityText.CreateTime = time.Now().Unix()
	PassivityText.MsgType = encrypt.Value2CDATA("text")
	PassivityText.Content =  encrypt.Value2CDATA(Content)
	if xmlbytes,err := xml.Marshal(PassivityText);err == nil {
		return xmlbytes
	}
	return nil
}

//回复转接客服
func (self *Wechat) ReplyCustom(FromUserName,ToUserName,Content string) []byte {
	var PassivityText encrypt.PassivityCustom
	PassivityText.ToUserName =  encrypt.Value2CDATA(ToUserName)
	PassivityText.FromUserName = encrypt.Value2CDATA(FromUserName)
	PassivityText.CreateTime = time.Now().Unix()
	PassivityText.MsgType = encrypt.Value2CDATA("transfer_customer_service")
	if xmlbytes,err := xml.Marshal(PassivityText);err == nil {
		return xmlbytes
	}
	return nil
}

//返回系统客服的content参数
func (self *Wechat) ServiceData(data[]byte)( string,string) {
	array := make(map[string]interface{})
	if self.MsgType == "text" {
		message := self.TextMessage(data)
		array["msgtype"] = message.MsgType
		array["content"] = message.Content
		array["createtime"] = strconv.FormatInt(message.CreateTime,10)
		data,_:=json.Marshal(array)
		return message.FromUserName ,string(data)
	}
	if self.MsgType == "image" {
		message := self.ImageMessage(data)
		array["msgtype"] = message.MsgType
		array["picurl"] = message.PicUrl
		array["createtime"] = strconv.FormatInt(message.CreateTime,10)
		data,_:=json.Marshal(array)
		return message.FromUserName ,string(data)
	}
	if self.MsgType == "voice" {
		message := self.VoiceMessage(data)
		array["msgtype"] = message.MsgType
		array["mediaid"] = message.MediaId
		array["createtime"] = message.CreateTime
		data,_:=json.Marshal(array)
		return message.FromUserName ,string(data)
	}
	if self.MsgType == "location" {
		message := self.LocationMessage(data)
		array["msgtype"] = message.MsgType
		array["label"] = message.Label
		array["createtime"] = strconv.FormatInt(message.CreateTime,10)
		data,_:=json.Marshal(array)
		return message.FromUserName ,string(data)
	}
	if self.MsgType == "miniprogrampage" {
		message := self.ReplyMiniprogrampagen(data)
		array["msgtype"] = message.MsgType
		array["pagepath"] = message.PagePath
		array["thumburl"] = message.ThumbUrl
		data,_:=json.Marshal(array)
		return message.FromUserName ,string(data)
	}
	return "",""
}