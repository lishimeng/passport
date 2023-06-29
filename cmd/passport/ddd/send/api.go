package send

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/internal/common"
	"github.com/lishimeng/passport/internal/db/model"
	"github.com/lishimeng/passport/internal/notify"
	"time"
)

func GetSendCode(ctx iris.Context) {
	var resp app.Response
	receiver := ctx.URLParam("receiver")
	codeLoginType := ctx.URLParam("codeLoginType")
	var value string
	err := app.GetCache().Get(receiver, &value)
	if err != nil {
		log.Info("获取缓存验证码失败%s：%s", receiver, err)
	}
	if len(value) > 0 {
		log.Debug("缓存Code:%s,%s", receiver, value)
		resp.Code = tool.RespCodeError
		resp.Message = "验证码已发送,请稍后重试！"
		tool.ResponseJSON(ctx, resp)
		return
	}
	//生成4位验证码
	var code = common.RandCode(4)
	log.Debug("code:%s", code)
	switch codeLoginType {
	case string(model.SmsNotifyType):
		log.Info("发送短信：%s", receiver)
		sms, err := notify.SighInSendSms(code, receiver)
		if err != nil || sms.Code != float64(tool.RespCodeSuccess) {
			resp.Code = tool.RespCodeError
			resp.Message = "验证码发送失败,请稍后重试！"
			tool.ResponseJSON(ctx, resp)
			return
		}
		break
	case string(model.MailNotifyType):
		log.Info("发送邮件：%s", receiver)
		mail, err := notify.SighInSendMail(code, receiver)
		if err != nil || mail.Code != float64(tool.RespCodeSuccess) {
			resp.Code = tool.RespCodeError
			resp.Message = "验证码发送失败,请稍后重试！"
			tool.ResponseJSON(ctx, resp)
			return
		}
		break
	default:
		resp.Code = tool.RespCodeError
		resp.Message = "验证码未发送,请稍后重试！"
		tool.ResponseJSON(ctx, resp)
		return
	}
	//缓存验证码 3分钟过期 key=邮箱
	err = app.GetCache().SetTTL(receiver, code, time.Minute)
	if err != nil {
		log.Info("缓存验证码失败：%s", err)
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

type sendCodeReq struct {
	Receiver      string `json:"receiver,omitempty"`
	CodeLoginType string `json:"codeLoginType,omitempty"`
	LoginType     string `json:"loginType,omitempty"`
}

func PostSendCode(ctx iris.Context) {
	var resp app.Response
	var req sendCodeReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	receiver := req.Receiver
	codeLoginType := req.CodeLoginType
	var value string
	err = app.GetCache().Get(receiver, &value)
	if err != nil {
		log.Info("获取缓存验证码失败%s：%s", receiver, err)
	}
	//生成4位验证码
	var code = common.RandCode(4)
	log.Debug("code:%s", code)
	if len(value) > 0 {
		resp.Code = tool.RespCodeError
		resp.Message = "验证码已发送,请稍后重试！"
		tool.ResponseJSON(ctx, resp)
		return
	}
	switch codeLoginType {
	case string(model.SmsNotifyType):
		log.Info("发送短信：%s", receiver)
		sms, err := notify.SighInSendSms(code, receiver)
		if err != nil || sms.Code != float64(tool.RespCodeSuccess) {
			resp.Code = tool.RespCodeError
			resp.Message = "验证码发送失败,请稍后重试！"
			tool.ResponseJSON(ctx, resp)
			return
		}
		break
	case string(model.MailNotifyType):
		log.Info("发送邮件：%s", receiver)
		mail, err := notify.SighInSendMail(code, receiver)
		if err != nil || mail.Code != float64(tool.RespCodeSuccess) {
			resp.Code = tool.RespCodeError
			resp.Message = "验证码发送失败,请稍后重试！"
			tool.ResponseJSON(ctx, resp)
			return
		}
		break
	}
	//缓存验证码 3分钟过期 key=邮箱
	err = app.GetCache().SetTTL(receiver, code, time.Minute*3)
	if err != nil {
		log.Info("缓存验证码失败：%s", err)
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
