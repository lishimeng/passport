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

// 发送手机短信
type SmsInfo struct {
	Template      string      `json:"template"`         // template of this mail
	TemplateParam interface{} `json:"params"`           // template params
	Sender        string      `json:"sender,omitempty"` // sms send account on the platform
	Signature     string      `json:"signature,omitempty"`
	Receiver      string      `json:"receiver"` // receiver list(with comma if multi)
}

type SmsResp struct {
	Code      interface{} `json:"code"`
	MessageId int         `json:"messageId,omitempty"`
}

func SendSms(sms SmsInfo) (smsResp SmsResp, err error) {
	client := &http.Client{Timeout: 8 * time.Second}
	jsonStr, _ := json.Marshal(sms)
	url := etc.Config.Notify.Host + etc.Config.Notify.Sms
	log.Info("url:%s", url)
	log.Info("jsonStr:%s", jsonStr)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Info("client Post err")
		return
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(result, &smsResp)
	if err != nil {
		log.Info("response Unmarshal err, %+v", smsResp)
		return
	}
	log.Info("sendMail response: %s", smsResp)
	return
}
