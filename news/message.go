package news

import (
	"../api"
	"encoding/json"
)
//发送应用消息
func MessageSend(assess_token string,data map[string]interface{}) (array api.PortData,err *api.CryptError) {
	URL := api.HostUrl(api.MESSAGE_SEND,assess_token)
	jsonbyte,_ := json.Marshal(data)
	return array,api.PostUnmarshal(URL,jsonbyte,&array)
}
