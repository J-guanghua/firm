package firm

import (
	"./encrypt"
	"fmt"
	"encoding/xml"
	"time"
)

type Encrypt struct {
	Firm *Company
	Token string
	EncodingAeskey string
	ReceiverId string
	encrypt *encrypt.WXBizMsgCrypt

}

//实例化
func NewEncrypt(token, encodingAeskey, receiverId,secret string) *Encrypt {
	return &Encrypt{
		Firm:NewCompany(receiverId,secret),
		encrypt:encrypt.NewWXBizMsgCrypt(token, encodingAeskey, receiverId, encrypt.XmlType),
	}
}
//验证回调URL
func (self *Encrypt) VerifyURL(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr string) []byte {
	echoStr, cryptErr := self.encrypt.VerifyURL(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr)
	if nil != cryptErr {
		fmt.Println("verifyUrl fail", cryptErr)
		return nil
	}
	fmt.Println("verifyUrl success echoStr", string(echoStr))
	return echoStr
}

//解密Encrypt
func (self *Encrypt) DecryptMsg(msg_signature, timestamp, nonce string, post_data []byte) (msgtype encrypt.MsgType,xmls []byte){

	msg, cryptErr := self.encrypt.DecryptMsg(msg_signature, timestamp, nonce, post_data)
	if nil != cryptErr {
		fmt.Println("DecryptMsg fail", cryptErr)
		return msgtype,nil
	}
	xml.Unmarshal(msg,&msgtype)
	return msgtype,msg
}

//解密Encrypt转interface
func (self *Encrypt) DecryptMsgInterface(msg_signature, timestamp, nonce string, post_data []byte) (msgtype encrypt.MsgType,xmls interface{}){
	//解密Encryp
	msg_type, data := self.DecryptMsg(msg_signature, timestamp, nonce, post_data)
	//企业事件
	if msgtype.MsgType == encrypt.EVENT {
		return msgtype,self.ClientDecryptMsg(msg_type,data)
	}
	return msgtype,self.DecryptMsgStruct(msg_type,data) //应用消息
}

//事件消息信息转struct结构
func (self *Encrypt) DecryptMsgStruct(msgtype encrypt.MsgType,data []byte) interface{}{
	if msgtype.MsgType == encrypt.TEXT {
		return self.TextMessage(data)
	}
	if msgtype.MsgType == encrypt.IMAGE {
		return self.ImageMessage(data)
	}
	if msgtype.MsgType == encrypt.VIDEO {
		return self.VideoMessage(data)
	}
	if msgtype.MsgType == encrypt.VOICE {
		return self.VoiceMessage(data)
	}
	return nil
}

//企业客户事件信息struct结构
func (self *Encrypt) ClientDecryptMsg(msgtype encrypt.MsgType,data []byte) interface{}{
	if msgtype.ChangeType == encrypt.ADD_EXTTERNAL_CONTACT {
		return self.AddExternalContact(data)
	}
	if msgtype.ChangeType == encrypt.ADD_HALF_EXTERNAL_CONTACT {
		return self.AddHalfExternalContact(data)
	}
	if msgtype.ChangeType == encrypt.DEL_EXTERNAL_CONTACT {
		return self.DelExternalContact(data)
	}
	if msgtype.ChangeType == encrypt.DEL_FOLLOW_USER {
		return self.DelFollowUser(data)
	}
	return nil
}
//--------------------------------消息事件-------------------------------//
//文本消息
func (self *Encrypt) TextMessage(data []byte) encrypt.TextMessage {
	var TextMessage encrypt.TextMessage
	xml.Unmarshal(data,&TextMessage)
	return TextMessage
}
//图片消息
func (self *Encrypt) ImageMessage(data []byte) encrypt.ImageMessage {
	var ImageMessage encrypt.ImageMessage
	xml.Unmarshal(data,&ImageMessage)
	return ImageMessage
}
//语音消息
func (self *Encrypt) VideoMessage(data []byte) encrypt.VideoMessage {
	var VideoMessage encrypt.VideoMessage
	xml.Unmarshal(data,&VideoMessage)
	return VideoMessage
}
//视频消息
func (self *Encrypt) VoiceMessage(data []byte) encrypt.VoiceMessage {
	var VoiceMessage encrypt.VoiceMessage
	xml.Unmarshal(data,&VoiceMessage)
	return VoiceMessage
}
//添加企业客户事件
func (self *Encrypt) AddExternalContact(data []byte) encrypt.AddExternalContact {
	var AddExternalContact encrypt.AddExternalContact
	xml.Unmarshal(data,&AddExternalContact)
	return AddExternalContact
}
//外部联系人免验证添加成员事件
func (self *Encrypt) AddHalfExternalContact(data []byte) encrypt.AddHalfExternalContact {
	var AddHalfExternalContact encrypt.AddHalfExternalContact
	xml.Unmarshal(data,&AddHalfExternalContact)
	return AddHalfExternalContact
}
//删除企业客户事件
func (self *Encrypt) DelExternalContact(data []byte) encrypt.DelExternalContact {
	var DelExternalContact encrypt.DelExternalContact
	xml.Unmarshal(data,&DelExternalContact)
	return DelExternalContact
}
//删除跟进成员事件
func (self *Encrypt) DelFollowUser(data []byte) encrypt.DelFollowUser {
	var DelFollowUser encrypt.DelFollowUser
	xml.Unmarshal(data,&DelFollowUser)
	return DelFollowUser
}
//--------------------------------被动回复消息-------------------------------//

//回复文本消息
func (self *Encrypt) ReplyText(FromUserName,ToUserName,Content string) []byte {
	var PassivityText encrypt.PassivityText
	PassivityText.ToUserName =  encrypt.Value2CDATA(ToUserName)
	PassivityText.FromUserName = encrypt.Value2CDATA(FromUserName)
	PassivityText.CreateTime = time.Now().Unix()
	PassivityText.MsgType = encrypt.Value2CDATA("text")
	PassivityText.Content =  encrypt.Value2CDATA(Content)
	if xmlbytes,err :=xml.Marshal(PassivityText);err == nil {
		return xmlbytes
	}
	return nil
}