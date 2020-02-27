package wechat

import (
"time"
"fmt"
"sort"
"strings"
"crypto/md5"
"net/http"
"bytes"
"io/ioutil"
"encoding/xml"
"crypto/tls"
"math/rand"
"strconv"
"encoding/hex"
)


var strlength = 32
//EnterpriseXml：：企业付款
type EnterpriseXml struct {
	XMLName xml.Name 		`xml:"xml"`
	Nonce_str  string  		`xml:"nonce_str"`
	Partner_trade_no string `xml:"partner_trade_no"`
	Mchid int64 			`xml:"mchid"`
	Mch_appid string		`xml:"mch_appid"`
	Openid string 			`xml:"openid"`
	Check_name string 		`xml:"check_name"`
	Amount string 			`xml:"amount"`
	Desc string 			`xml:"desc"`
	Spbill_create_ip string `xml:"spbill_create_ip"`
	Remark string			`xml:"remark"`
	Sign string 			`xml:"sign"`
}

//WechatBonusXml：：微信红包
type WechatBonusXml  struct {
	XMLName xml.Name 			`xml:"xml"`
	Nonce_str  string  			`xml:"nonce_str"`
	Mch_billno string 			`xml:"mch_billno"`
	Mch_id string 				`xml:"mch_id"`
	Wxappid string		 		`xml:"wxappid"`
	Re_openid string 			`xml:"re_openid"`
	Send_name string 			`xml:"send_name"`
	Total_amount int64 			`xml:"total_amount"`
	Total_num int64 			`xml:"total_num"`
	Wishing string 				`xml:"wishing"`
	Client_ip string			`xml:"client_ip"`
	Act_name string 			`xml:"act_name"`
	Remark string				`xml:"remark"`
	Sign string					`xml:"sign"`
}

//微信红包结构体
type Wechat_Bouns struct {
	 Mch_id ,Send_name, Key ,Cert string
	Wechat *Wechat
	certPath string //= "/home/J-gaunghua/advanced-web/gopath/swoole/cert/wx7598ac265e3a9bf5/cert/apiclient_cert.pem"
	keyPath string //= "/home/J-gaunghua/advanced-web/gopath/swoole/cert/wx7598ac265e3a9bf5/cert/apiclient_key.pem"
	cAPath  string //= "/home/J-gaunghua/advanced-web/gopath/swoole/cert/wx7598ac265e3a9bf5/cert/apiclient_cert.pem"
}
//微信红包对象
func NewWechatBouns(Mch_id,Send_name,Key,Cert string) *Wechat_Bouns {
	return &Wechat_Bouns{
		Mch_id:Mch_id,
		Send_name:Send_name,
		Key:Key,
		Cert:Cert,
	}
}
//设置必要参数
func  (self *Wechat_Bouns) SetWechatBounsPathInfo(wec *Wechat,CertPath,KeyPath,CAPath string)  {
	self.Wechat = wec
	self.certPath = CertPath
	self.keyPath = KeyPath
	self.cAPath = CAPath
	return
}
//GetNonceStr产生随机字符串，不长于32位
func (self *Wechat_Bouns)  GetNonceStr() string {
	chaers := "abcdefghijklmnopqrstuvwxyz0123456789"
	return chaers[:strlength]
}

//GetOrderSn::获取订单编号
func (self *Wechat_Bouns)  GetOrderSn() string {
	timeint :=time.Now().Unix()
	dateday := time.Now().YearDay()
	return fmt.Sprintf("%v%v",dateday,timeint)
}

//微信红包参数格式转换
func (self *Wechat_Bouns) ToMakeMap(Bonus WechatBonusXml) map[string]interface{} {
	array := make(map[string]interface{})
	array["nonce_str"] = Bonus.Nonce_str
	array["mch_billno"] = Bonus.Mch_billno
	array["mch_id"] = Bonus.Mch_id
	array["wxappid"] = Bonus.Wxappid
	array["send_name"] = Bonus.Send_name
	array["re_openid"] = Bonus.Re_openid
	array["total_amount"] = Bonus.Total_amount
	array["total_num"] = Bonus.Total_num
	array["wishing"] = Bonus.Wishing
	array["client_ip"] = Bonus.Client_ip
	array["act_name"] = Bonus.Act_name
	array["remark"] = Bonus.Remark
	return array
}

//GetMakeSign:签名，本函数不覆盖sign成员变量，如要设置签名需要调用SetSign方法赋值
func (self *Wechat_Bouns) GetMakeSign(array map[string]interface{}, key string) string {

	sorted_keys := make([]string, 0)
	for v,_:=range array {
		sorted_keys = append(sorted_keys,v)
	}
	sort.Strings(sorted_keys)
	var buff string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", array[k])
		if value != "" {
			buff = buff + k + "=" + value + "&"
		}
	}
	buff = strings.Trim(buff,"&")
	buff = buff + "&key="+key
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(buff))
	cipherStr := md5Ctx.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}

//RandInt64::int64: 生成的随机数
func (self *Wechat_Bouns)  RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(max-min) + min
}

//ConsoleSendBonus::发送微信红包
func (self *Wechat_Bouns)  ConsoleSendBonus(Bonus WechatBonusXml,total_amount string) (WechatBonusXml,[]byte,error) {

	//随机金额
	if index := strings.Index(total_amount,"~");index > 0 {
		money := strings.Split(total_amount,"~")
		min,_ := strconv.ParseInt(money[0],10,64)
		max,_ := strconv.ParseInt(money[1],10,64)
		Bonus.Total_amount = self.RandInt64(min,max)
	} else {
		Bonus.Total_amount,_ = strconv.ParseInt(total_amount,10,64)
	}
	Bonus.Nonce_str 	= self.GetNonceStr()
	Bonus.Mch_billno 	= self.GetOrderSn()
	Bonus.Mch_id 		= self.Mch_id
	Bonus.Wxappid 		= self.Wechat.Appid
	Bonus.Send_name 	= self.Send_name
	Bonus.Total_num 	= 1
	Bonus.Client_ip 	= "127.0.0.1"
	Bonus.Remark 		= "恭喜你获取微信红包！"
	array := self.ToMakeMap(Bonus)
	Bonus.Sign 			= self.GetMakeSign(array,self.Key)
	byte,err := self.GetBonusXml(Bonus)
	return Bonus,byte,err
}

//postXmlCurl：：xml curl请求
func (self *Wechat_Bouns) PostXmlCurl(url string, xmlData []byte) ([]byte, error) {

	tlsConfig, err := self.getTLSConfig()
	if err != nil {
		return nil, err
	}
	tr := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: tr}

	resp, err := client.Post(url,"text/xml",bytes.NewBuffer(xmlData))
	if err != nil {
		return nil, err
	}
	Bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return Bytes,nil
}

//getTLSConfig：：加载证书
func (self *Wechat_Bouns) getTLSConfig() (*tls.Config, error) {

	var TlsConfig *tls.Config
	//加载证书
	cert, err := tls.LoadX509KeyPair(self.Cert + "/apiclient_cert.pem", self.Cert + "/apiclient_key.pem")
	if err != nil {
		fmt.Println("getTLSConfig-1", err)
		return nil, err
	}
	//caData, err := ioutil.ReadFile(wechatCAPath)
	//if err != nil {
	//	fmt.Println("getTLSConfig-2", err)
	//	return nil, err
	//}
	//pool := x509.NewCertPool()
	//pool.AppendCertsFromPEM(caData)
	TlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		//RootCAs:      pool,
	}
	return TlsConfig, nil
}

//GetBonusXml::发送红包
func (self *Wechat_Bouns) GetBonusXml(Bonus WechatBonusXml) ([]byte,error)  {
	url := "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack"
	PostXml,_ := xml.Marshal(Bonus)
	return self.PostXmlCurl(url,PostXml)
}



