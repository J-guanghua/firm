package firm

import (
	"./token"
	"time"
	"net/http"
	"fmt"
	"../firm/api"
)

type JsSdk struct {
	Firm *Company
}

func NewJsSdk(CorpId,Secret string) *JsSdk {
	return &JsSdk{
		Firm:NewCompany(CorpId,Secret),
	}
}
//获取企业的jsapi_ticket
func (self *JsSdk) JsapiTicket() (string,*api.CryptError) {
	access_token := self.Firm.access_token(self.Firm.Secret)
	return  token.Jsapi_Ticket(access_token)
}
//Code2_Session
func (self *JsSdk) Code2Session(js_code string) (token.CodeSession,*api.CryptError) {
	access_token := self.Firm.access_token(self.Firm.Secret)
	return token.Code2_Session(access_token,js_code)
}
//JS-SDK 授权参数
func (self *JsSdk) Sign_Package(r *http.Request) (api.PortData,*api.CryptError) {
	var data = make(api.PortData);var err *api.CryptError
	data["appId"] = self.Firm.CorpId
	data["timestamp"] = time.Now().Unix()
	data["nonceStr"] = token.GetRandomString(16)
	data["jsapi_ticket"],err = self.JsapiTicket()
	signature := fmt.Sprintf("jsapi_ticket=%v&noncestr=%s&timestamp=%v&url=%v",
		data["jsapi_ticket"],data["nonceStr"],data["timestamp"],"https://" + r.Host + r.RequestURI)
	data["signature"] = token.Str2sha1(signature)
	return data,err

}

