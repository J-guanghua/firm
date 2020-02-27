package external

import (
	"../api"
)
//配置客户联系「联系我」方式
func Add_Contact_Way(assess_token string) (array api.PortData,err *api.CryptError) {
	URL := api.HostUrl(api.ADD_CONTACT_WAY,assess_token)
	return array,api.PostUnmarshal(URL,nil,&array)
}