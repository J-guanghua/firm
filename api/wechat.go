package api

import "fmt"

const TOKEN  = "/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s" //获取token
const SEND_CUSTOM  = "/cgi-bin/message/custom/send?access_token=%s"  //发送客服消息
const UPLOAD  = "/cgi-bin/media/upload?access_token=%s&type=%s"      //上传素材
const SEND_TEMPLATE   = "cgi-bin/message/template/send?access_token=%s"   //发送模板消息
const BATCHTAGING   = "/cgi-bin/tags/members/batchtagging?access_token=%s" //批量打标签
const USER_INFO = "/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN" //获取用户信息
const QRCONNECT =  "/connect/qrconnect?appid=%s&response_type=code&redirect_uri=%s" //扫码登陆
const QRCONNECT_OAUTH2 =  "/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code" //扫码授权用户信息
const API_REQUEST_URL  = "https://api.weixin.qq.com" //微信公众号
const OPEN_REQUEST_URL  = "https://open.weixin.qq.com" //开发平台

func ApiUrl(url string,cols ...string)  string {
	length := len(cols)
	if length == 1 { url = fmt.Sprintf(url,cols[0]) }
	if length == 2 { url = fmt.Sprintf(url,cols[0],cols[1]) }
	if length == 3 { url = fmt.Sprintf(url,cols[0],cols[1],cols[2]) }
	if length == 4 { url = fmt.Sprintf(url,cols[0],cols[1],cols[2],cols[3]) }
	return API_REQUEST_URL + url
}

func OpenUrl(url string,cols ...string)  string {
	length := len(cols)
	if length == 1 { url = fmt.Sprintf(url,cols[0]) }
	if length == 2 { url = fmt.Sprintf(url,cols[0],cols[1]) }
	if length == 3 { url = fmt.Sprintf(url,cols[0],cols[1],cols[2]) }
	if length == 4 { url = fmt.Sprintf(url,cols[0],cols[1],cols[2],cols[3]) }
	return OPEN_REQUEST_URL + url
}