package firm

import 	(
	"../firm/token"
	"../firm/api"
	"../firm/external"
	"log"
	"fmt"
)
//企业内部开发
type Company struct {
	CorpId string //企业ID
	Secret string //企业客户的Secret
	Apply_Secret string //企业应用的Secret
	Contacts_Secret string //通讯录的Secret
}
//返回企业实例
func NewCompany(CorpId,Secret string) *Company {
	return &Company{
		CorpId:CorpId,
		Secret:Secret,
		Apply_Secret:"kwIFRBk6VanvFIziM1Aeo-R4aXYTbg5LBoIFB6Stz0A",
		Contacts_Secret:"FSbBFTyntWJaIeLo3GTSnQwSgGv9SDY1gljopm93_Rc",
	}
}
//设置企业应用的Secret
func (self *Company) Set_Apply_Secret(apply_secret string)  {
	self.Apply_Secret = apply_secret
}
//获取企业内部开发Access_token
func (self *Company) access_token(Secret string) string {
	access_token,err := token.Get_Token(self.CorpId,Secret)
	if err != nil{
		log.Fatal(fmt.Sprintf("access_token" +
			"获取失败:状态码：%v,返回信息：%v",err.ErrCode,err.ErrMsg))
	}
	return access_token
}
//获取客户信息
func (self *Company) ExternalUserid(external_userid string) (api.PortData,*api.CryptError) {

	return external.External_Userid(self.access_token(self.Secret),external_userid)
}
//读取成员信息
func (self *Company) User_get(userid string) (api.PortData,*api.CryptError)  {

	return external.User_Get(self.access_token(self.Secret),userid)
}