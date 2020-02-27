package encrypt

import "encoding/xml"

type ImgMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	PicUrl CDATAText
	MediaId CDATAText
}

