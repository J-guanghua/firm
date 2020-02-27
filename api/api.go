package api

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
	"bytes"
	"encoding/json"
	"net/url"
	"crypto/tls"
)

const GET_TOKEN        	= "/cgi-bin/gettoken?corpid=%s&corpsecret=%s" //获取token
const EXTERNLOCNTACT     ="/cgi-bin/externalcontact/get?access_token=%s&external_userid=%s" //获取客户详情
const GET_FOLLOW_USER_LIST  = "/cgi-bin/externalcontact/get_follow_user_list?access_token=%s" //获取配置了客户联系功能的成员列表
const EXTERNALCONTACT_LIST  = "/cgi-bin/externalcontact/list?access_token=%s&userid=%s" //获取客户列表
const CONVERT_TO_OPENID  ="/cgi-bin/externalcontact/convert_to_openid?access_token=%s"
const REMARK  				 = "/cgi-bin/externalcontact/remark?access_token=%s" //修改客户备注信息
const MARK_TAG  			 = "/cgi-bin/externalcontact/remark?access_token=%s" //修改客户备注信息
const GET_CORP_TAG_LIST  = "/cgi-bin/externalcontact/get_corp_tag_list?access_token=%s" //获取企业标签库
const REDIRECT_URL        = "/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=STATE#wechat_redirect"//构造网页授权链接
const GET_USER_INFO_BY_CODE = "/cgi-bin/user/getuserinfo?access_token=%s&code=%s" //该接口用于根据code获取成员信息
const MESSAGE_SEND       = "/cgi-bin/message/send?access_token=%s"//发送应用消息
const USER_GET            = "/cgi-bin/user/get?access_token=%s&userid=%s" //读取成员信息
const GET_JSAPI_TICKET  = "/cgi-bin/get_jsapi_ticket?access_token=%s" //获取企业的jsapi_ticket
const GET_TICKET         = "/cgi-bin/ticket/get?access_token=%s" //获取应用的jsapi_ticket
const JSCODE2_SESSION   = "/cgi-bin/miniprogram/jscode2session?access_token=%s&js_code=%s&grant_type=authorization_code"
const ADD_CONTACT_WAY   = "/cgi-bin/externalcontact/add_contact_way?access_token=%s" //配置客户联系「联系我」方式
const MEDIA_UPLOAD        = "/cgi-bin/media/upload?access_token=%s&type=%s"
const USERID_TO_OPENID  = "/cgi-bin/user/convert_to_openid?access_token=%s" //userid与openid互换
const USER_CREATE        = "/cgi-bin/user/create?access_token=ACCESS_TOKEN"
const USER_UPDATE        ="/cgi-bin/user/update?access_token=ACCESS_TOKEN"
const USER_DELETE        = "/cgi-bin/user/delete?access_token=ACCESS_TOKEN"
const USER_BATCH_DELETE = "/cgi-bin/user/batchdelete?access_token=ACCESS_TOKEN"
const USER_SIMPLE_LIST  = "/cgi-bin/user/simplelist?access_token=ACCESS_TOKEN"
const USER_LIST          = "/cgi-bin/user/list?access_token=ACCESS_TOKEN"
const OPENID_TO_USERID  = "/cgi-bin/user/convert_to_userid?access_token=ACCESS_TOKEN"
const USER_AUTH_SUCCESS = "/cgi-bin/user/authsucc?access_token=ACCESS_TOKEN"
const DEPARTMENT_CREATE = "/cgi-bin/department/create?access_token=ACCESS_TOKEN"
const DEPARTMENT_UPDATE = "/cgi-bin/department/update?access_token=ACCESS_TOKEN"
const DEPARTMENT_DELETE = "/cgi-bin/department/delete?access_token=ACCESS_TOKEN"
const DEPARTMENT_LIST   = "/cgi-bin/department/list?access_token=ACCESS_TOKEN"
const TAG_CREATE        = "/cgi-bin/tag/create?access_token=ACCESS_TOKEN"
const TAG_UPDATE        = "/cgi-bin/tag/update?access_token=ACCESS_TOKEN"
const TAG_DELETE        = "/cgi-bin/tag/delete?access_token=ACCESS_TOKEN"
const TAG_GET_USER      = "/cgi-bin/tag/get?access_token=ACCESS_TOKEN"
const TAG_ADD_USER      = "/cgi-bin/tag/addtagusers?access_token=ACCESS_TOKEN"
const TAG_DELETE_USER   = "/cgi-bin/tag/deltagusers?access_token=ACCESS_TOKEN"
const TAG_GET_LIST      = "/cgi-bin/tag/list?access_token=ACCESS_TOKEN"
const BATCH_JOB_GET_RESULT = "/cgi-bin/batch/getresult?access_token=ACCESS_TOKEN"
const BATCH_INVITE      = "/cgi-bin/batch/invite?access_token=ACCESS_TOKEN"
const AGENT_GET         = "/cgi-bin/agent/get?access_token=ACCESS_TOKEN"
const AGENT_SET         = "/cgi-bin/agent/set?access_token=ACCESS_TOKEN"
const AGENT_GET_LIST    = "/cgi-bin/agent/list?access_token=ACCESS_TOKEN"
const MENU_CREATE       = "/cgi-bin/menu/create?access_token=ACCESS_TOKEN"
const MENU_GET          = "/cgi-bin/menu/get?access_token=ACCESS_TOKEN"
const MENU_DELETE       = "/cgi-bin/menu/delete?access_token=ACCESS_TOKEN"
const MEDIA_GET         = "/cgi-bin/media/get?access_token=ACCESS_TOKEN"
const GET_USER_DETAIL   = "/cgi-bin/user/getuserdetail?access_token=ACCESS_TOKEN"
const GET_CHECKIN_OPTION  = "/cgi-bin/checkin/getcheckinoption?access_token=ACCESS_TOKEN"
const GET_CHECKIN_DATA     = "/cgi-bin/checkin/getcheckindata?access_token=ACCESS_TOKEN"
const GET_APPROVAL_DATA    = "/cgi-bin/corp/getapprovaldata?access_token=ACCESS_TOKEN"
const GET_INVOICE_INFO     = "/cgi-bin/card/invoice/reimburse/getinvoiceinfo?access_token=ACCESS_TOKEN"
const UPDATE_INVOICE_STATUS = "/cgi-bin/card/invoice/reimburse/updateinvoicestatus?access_token=ACCESS_TOKEN"
const BATCH_UPDATE_INVOICE_STATUS = "/cgi-bin/card/invoice/reimburse/updatestatusbatch?access_token=ACCESS_TOKEN"
const BATCH_GET_INVOICE_INFO = "/cgi-bin/card/invoice/reimburse/getinvoiceinfobatch?access_token=ACCESS_TOKEN"
const GET_PRE_AUTH_CODE      = "/cgi-bin/service/get_pre_auth_code?suite_access_token=SUITE_ACCESS_TOKEN"
const SET_SESSION_INFO       = "/cgi-bin/service/set_session_info?suite_access_token=SUITE_ACCESS_TOKEN"
const GET_PERMANENT_CODE     = "/cgi-bin/service/get_permanent_code?suite_access_token=SUITE_ACCESS_TOKEN"
const GET_AUTH_INFO           = "/cgi-bin/service/get_auth_info?suite_access_token=SUITE_ACCESS_TOKEN"
const GET_ADMIN_LIST          = "/cgi-bin/service/get_admin_list?suite_access_token=SUITE_ACCESS_TOKEN"
const GET_USER_INFO_BY_3RD   = "/cgi-bin/service/getuserinfo3rd?suite_access_token=SUITE_ACCESS_TOKEN"
const GET_USER_DETAIL_BY_3RD = "/cgi-bin/service/getuserdetail3rd?suite_access_token=SUITE_ACCESS_TOKEN"
const GET_LOGIN_INFO          = "/cgi-bin/service/get_login_info?access_token=PROVIDER_ACCESS_TOKEN"
const GET_REGISTER_CODE       = "/cgi-bin/service/get_register_code?provider_access_token=PROVIDER_ACCESS_TOKEN"
const SET_AGENT_SCOPE          = "/cgi-bin/agent/set_scope"
const SET_CONTACT_SYNC_SUCCESS = "/cgi-bin/sync/contact_sync_success"
const REQUEST_URL  = "https://qyapi.weixin.qq.com"

//拼接请求的url
func HostUrl(url string,cols ...string)  string {
	length := len(cols)
	if length == 1 { url = fmt.Sprintf(url,cols[0]) }
	if length == 2 { url = fmt.Sprintf(url,cols[0],cols[1]) }
	if length == 3 { url = fmt.Sprintf(url,cols[0],cols[1],cols[2]) }
	if length == 4 { url = fmt.Sprintf(url,cols[0],cols[1],cols[2],cols[3]) }
	return REQUEST_URL + url
}
//接口数据接收
type PortData map[string]interface{}
//错误信息接收
type CryptError struct{
	ErrCode int
	ErrMsg string
}
func NewCryptError(err_code int, err_msg string) * CryptError{
	return &CryptError{ErrCode:err_code, ErrMsg: err_msg}
}

//返回响应结构数据 GET
func GetUnmarshal(url string,data interface{}) *CryptError {
	var err error;var body []byte
	if body,err = GetCurl(url);err == nil {
		if err = json.Unmarshal(body,&data);err != nil {
			return NewCryptError(4019,err.Error())
		}
		return nil
	}
	return NewCryptError(4011,err.Error())
}

//返回响应结构数据 POST
func PostUnmarshal(url string,haed []byte,data interface{}) *CryptError{
	var err error
	if haed,err = PostCurl(url,haed);err == nil {
		if err = json.Unmarshal(haed,&data);err != nil {
			return NewCryptError(4019,err.Error())
		}
		return nil
	}
	return NewCryptError(4011,err.Error())
}

//GET请求的url
func GetCurl(url string) (body []byte,err error) {
	var (
		resp *http.Response
	)
	if resp, err = http.Get(url);err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	if body, err = ioutil.ReadAll(resp.Body);err != nil {
		return nil,err
	}
	return body,nil
}

//POST请求的url
func PostCurl(url string,data []byte) (body []byte,err error) {
	timeout := time.Duration(5 * time.Second)//超时时间5ms
	client := &http.Client{
		Timeout: timeout,
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset","UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	if body,err = ioutil.ReadAll(resp.Body);err != nil {
		return nil,err
	}
	return body,nil
}

//代理
func Client() *http.Client {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("https://app.huwaishequ.com:443")
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy:           proxy,
	}
	return &http.Client{Transport: tr}
}