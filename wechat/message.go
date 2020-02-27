package wechat

import (
	"../api"
	"encoding/json"
	"os"
	"io"
	"net/http"
	"bytes"
	"mime/multipart"
	"path/filepath"
	"io/ioutil"
)

//素材上传返回信息
type MediaError struct {
	ErrCode int
	ErrMsg string
	Media_id string
}

//发送客服消息
func Send_Custom(data interface{},assess_token string) (err *api.CryptError) {
	jsonbyte,_ := json.Marshal(data)
	URL := api.ApiUrl(api.SEND_CUSTOM,assess_token)
	api.PostUnmarshal(URL,jsonbyte,&err)
	return
}
//模板消息
func Send_Template(data interface{},assess_token string) (err *api.CryptError) {
	jsonbyte,_ := json.Marshal(data)
	URL := api.ApiUrl(api.SEND_TEMPLATE,assess_token)
	api.PostUnmarshal(URL,jsonbyte,&err)
	return
}
//给用户打标签
func Sign_Users_Tag(data interface{},assess_token string) (err *api.CryptError) {
	jsonbyte,_ := json.Marshal(data)
	URL := api.ApiUrl(api.BATCHTAGING,assess_token)
	api.PostUnmarshal(URL,jsonbyte,&err)
	return
}
//上传临时素材
func  Media_Upload(assess_token,genre,file_path string) (media *MediaError) {
	var (
		file *os.File
		part io.Writer
		req *http.Request
		res *http.Response
	)
	file,err:= os.Open(file_path)
	if err != nil {
		return &MediaError{4201,err.Error(),""}
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if part, err = writer.CreateFormFile(genre, filepath.Base(file_path));err!=nil{
		return &MediaError{ErrCode:4455,ErrMsg:err.Error()}
	}
	_, err = io.Copy(part, file)
	if err = writer.Close();err != nil{
		return &MediaError{4201,err.Error(),""}
	}
	URL := api.ApiUrl(api.UPLOAD,assess_token,genre)
	req, err = http.NewRequest("POST", URL, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	urlQuery := req.URL.Query()
	if err != nil {
		return &MediaError{4201,err.Error(),""}
	}
	urlQuery.Add("access_token", assess_token)
	urlQuery.Add("type", genre)
	req.URL.RawQuery = urlQuery.Encode()
	client := http.Client{}
	if res, err = client.Do(req);err !=nil {
		return &MediaError{4201,err.Error(),""}
	}
	defer res.Body.Close()
	data,err := ioutil.ReadAll(res.Body)
	json.Unmarshal(data,&media)
	return
}

//上传临时素材
func  Image_Upload(assess_token,genre,file_path string) (media *MediaError) {
	var (
		file *os.File
		part io.Writer
		req *http.Request
		res *http.Response
	)
	file,err:= os.Open(file_path)
	if err != nil {
		return &MediaError{4201,err.Error(),""}
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if part, err = writer.CreateFormFile(genre, filepath.Base(file_path));err!=nil{
		return &MediaError{ErrCode:4455,ErrMsg:err.Error()}
	}
	_, err = io.Copy(part, file)
	if err = writer.Close();err != nil{
		return &MediaError{4201,err.Error(),""}
	}
	URL := api.ApiUrl(api.MEDIA_UPLOAD,assess_token,genre)
	req, err = http.NewRequest("POST", URL, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	urlQuery := req.URL.Query()
	if err != nil {
		return &MediaError{4201,err.Error(),""}
	}
	urlQuery.Add("access_token", assess_token)
	urlQuery.Add("type", genre)
	req.URL.RawQuery = urlQuery.Encode()
	client := http.Client{}
	if res, err = client.Do(req);err !=nil {
		return &MediaError{4201,err.Error(),""}
	}
	defer res.Body.Close()
	data,err := ioutil.ReadAll(res.Body)
	json.Unmarshal(data,&media)
	return
}