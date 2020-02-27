package external

import (
	"../api"
	"encoding/json"
)
//获取企业标签库
func Corp_Tag_List(assess_token string,tag_id []string) (array api.PortData,err *api.CryptError) {
	var data = struct {
		Tag_id []string
	}{tag_id}
	jsonbyte,_ := json.Marshal(data)
	URL := api.HostUrl(api.GET_CORP_TAG_LIST,assess_token )
	return array,api.PostUnmarshal(URL,jsonbyte,&array)
}

//企业可通过此接口修改指定用户添加的客户的备注信息。
func Remark(assess_token string) (array api.PortData,err *api.CryptError) {
	URL := api.HostUrl(api.REMARK,assess_token )
	return array,api.PostUnmarshal(URL,nil,&array)
}

//编辑客户企业标签
func Mark_Tag(assess_token ,userid,external_userid string,add_tag,remove_tag []string) (array api.PortData,err *api.CryptError) {
	var data = struct {
		Userid,External_userid string
		Add_tag,Remove_tag []string
	}{userid,external_userid,add_tag,remove_tag}
	jsonbyte,_ := json.Marshal(data)
	URL := api.HostUrl(api.MARK_TAG,assess_token)
	return array,api.PostUnmarshal(URL,jsonbyte,&array)
}