package token

import (
	"../api"
	"fmt"
	"crypto/sha1"
	"io"
	"math/rand"
	"time"
)

//获取Access_token结构
type Token struct {
	Errcode int
	Errmsg string
	Access_token string
	Expires_in int64
}
var ContainerToken = make(map[string]string) //存放Token
func init()  {
	go func() {
		report_ticker := time.NewTicker(time.Second * time.Duration(7200))
		for range report_ticker.C {
			fmt.Println("hello Token 清除",ContainerToken)
			for key,_:=range ContainerToken{
				delete(ContainerToken,key)
			}
		}
	}()
}

//获取企业的jsapi_ticket
type Ticket struct {
	Errcode int
	Errmsg string
	Ticket string
	Expires_in int64
}
//Code2Session
type CodeSession struct {
	Errcode int
	Errmsg string
	Session_key string
	Corpid string
	Userid int64
}
func TickerReturnToken(URL ,appid string) (access_token string,err *api.CryptError) {
	var wec Token
	if val,ok := ContainerToken[appid];ok {
		return val,nil
	}
	if err = api.GetUnmarshal(URL,&wec);err == nil && wec.Access_token != "" {
		ContainerToken[appid] = wec.Access_token
		return wec.Access_token,nil
	}
	return access_token,err
}

//str2sha1：：sha1加密
func Str2sha1(data string)string{
	t:= sha1.New()
	io.WriteString(t,data)
	return fmt.Sprintf("%x",t.Sum(nil))
}
//生成随机字符串
func  GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
//获取企业微信Access_token
func Get_Token(corpid ,corpsecret string)(string,*api.CryptError){
	var wechat Token
	URL := api.HostUrl(api.GET_TOKEN,corpid,corpsecret)
	return wechat.Access_token,api.GetUnmarshal(URL,&wechat)
}

//获取微信公众号平台Access_token
func GetToken(appid ,secret string)(string,*api.CryptError){
	URL := api.ApiUrl(api.TOKEN,appid,secret)
	return TickerReturnToken(URL,appid)
}



//获取企业的jsapi_ticket
func Jsapi_Ticket(access_token string) (string,*api.CryptError) {
	var wechat Ticket
	URL := api.HostUrl(api.GET_JSAPI_TICKET,access_token)
	return wechat.Ticket,api.GetUnmarshal(URL,&wechat)
}
//获取AgentId的jsapi_ticket
func AgentId_Ticket(access_token string) (string,*api.CryptError) {
	var wechat Ticket
	URL := api.HostUrl(api.GET_TICKET,access_token)
	return wechat.Ticket,api.GetUnmarshal(URL,&wechat)
}

//Code2Session
func Code2_Session(access_token ,js_code string) (wechat CodeSession,err *api.CryptError) {
	URL := api.HostUrl(api.JSCODE2_SESSION,access_token,js_code)
	return wechat,api.GetUnmarshal(URL,&wechat)
}