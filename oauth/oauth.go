package oauth

import (
	"../api"
)

//构造网页授权链接
func Redirect_Uri(appid,redirect_uri,scope string) string {
	return api.HostUrl(api.REDIRECT_URL,appid,redirect_uri,scope)
}

//该接口用于根据code获取成员信息
func GetUserInfo(assess_token,code string)(array api.PortData,err *api.CryptError) {
	URL := api.HostUrl(api.GET_USER_INFO_BY_CODE,assess_token)
	return array,api.GetUnmarshal(URL,&array)
}

//微信接口获取用户信息
func Qrconnect_Info(openid ,accesstoken string) (array api.PortData,err *api.CryptError) {
	URL := api.ApiUrl(api.USER_INFO,openid ,accesstoken)
	return array,api.GetUnmarshal(URL,&array)
}

//返回授权url链接
func Auth_Url(appid,redirect string) string {
	return api.OpenUrl(api.QRCONNECT,appid,redirect)
}

//获取用户信息数据
func Auth_Data(appid, secret ,code string) (array api.PortData,err *api.CryptError) {
	URL := api.ApiUrl(api.QRCONNECT_OAUTH2,appid, secret ,code)
	return array,api.GetUnmarshal(URL,&array)
}