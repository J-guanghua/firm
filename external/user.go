package external

import (
	"../api"
	"encoding/json"
)

//userid与openid互换
func Userid_To_Openid(assess_token string,data []byte) (array api.PortData,err *api.CryptError) {
	URL := api.HostUrl(api.USERID_TO_OPENID,assess_token)
	return array,api.GetUnmarshal(URL,&array)
}

//获取客户联系功能的成员列表
func Follow_User_List(assess_token string,data []byte) (array api.PortData,err *api.CryptError) {
	URL := api.HostUrl(api.GET_FOLLOW_USER_LIST,assess_token)
	return array,api.GetUnmarshal(URL,&array)
}

//获取客户列表
func External_Contact_List(assess_token string,userid string) (array api.PortData,err *api.CryptError) {
	URL := api.HostUrl(api.EXTERNALCONTACT_LIST,assess_token)
	return array,api.GetUnmarshal(URL,&array)
}

//获取客户详情
func External_Userid(assess_token,external_userid string) (array api.PortData,err *api.CryptError) {
	URL := api.HostUrl(api.EXTERNLOCNTACT,assess_token ,external_userid)
	return array,api.GetUnmarshal(URL,&array)
}

//外部联系人openid转换
func Convert_To_Openid(assess_token string,external_userid string) (array api.PortData,err *api.CryptError) {
	var data = struct {
		External_userid string
	}{external_userid}
	jsonbyte,_ := json.Marshal(data)
	URL := api.HostUrl(api.CONVERT_TO_OPENID,assess_token)
	return array,api.PostUnmarshal(URL,jsonbyte,&array)
}

//读取成员
func User_Get(assess_token string,userid string) (array api.PortData,err *api.CryptError) {
	URL := api.HostUrl(api.USER_GET,assess_token ,userid)
	return array,api.GetUnmarshal(URL,&array)
}
