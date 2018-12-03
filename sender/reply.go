package sender

import (
	"encoding/xml"
	"fmt"
	"github.com/MwlLj/go-wechat/common"
	"net/http"
)

type CReply struct {
	ResponseWriter http.ResponseWriter
}

func (this *CReply) SendMessage(msg *common.CMessage) error {
	response := *msg
	response.SetToUserName(&msg.FromUserName)
	response.SetFromUserName(&msg.ToUserName)
	res, err := xml.MarshalIndent(&response, " ", "  ")
	if err != nil {
		return err
	}
	this.ResponseWriter.Header().Set("Content-Type", "text/xml")
	fmt.Fprint(this.ResponseWriter, string(res))
	return nil
}
