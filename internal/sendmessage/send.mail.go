package sendmessage

import (
	"bytes"
	"encoding/json"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/internal/etc"
	"io/ioutil"
	"net/http"
	"time"
)

type Req struct {
	Template      string      `json:"template"`           // template of this mail
	TemplateParam interface{} `json:"params,omitempty"`   // template params
	Subject       string      `json:"subject,omitempty"`  // mail's subject
	Sender        string      `json:"sender,omitempty"`   // mail send account on the platform
	Receiver      string      `json:"receiver,omitempty"` // receiver list(with comma if multi)
	Cc            string      `json:"cc,omitempty"`       // cc list(with comma if multi)
}

type Resp struct {
	Code      interface{} `json:"code"`
	MessageId int         `json:"messageId,omitempty"`
}

func SendMail(sms Req) (mailResp Resp, err error) {
	client := &http.Client{Timeout: 8 * time.Second}
	jsonStr, _ := json.Marshal(sms)
	url := etc.Config.Notify.Host + etc.Config.Notify.Mail
	log.Info("url:%s", url)
	log.Info("jsonStr:%s", jsonStr)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Info("client Post err")
		return
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(result, &mailResp)
	if err != nil {
		log.Info("response Unmarshal err, %+v", mailResp)
		return
	}
	log.Info("sendMail response: %s", mailResp)
	return
}
