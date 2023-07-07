package send

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/cmd/passport/ddd/user"
	"github.com/lishimeng/passport/internal/common"
	"github.com/lishimeng/passport/internal/db/model"
	"github.com/lishimeng/passport/internal/notify"
	"time"
)

type signInReq struct {
	Receiver      string `json:"receiver,omitempty"`
	CodeLoginType string `json:"codeLoginType,omitempty"`
	LoginType     string `json:"loginType,omitempty"`
}

func signInSendCodeGet(ctx iris.Context) {
	var resp app.Response
	receiver := ctx.URLParam("receiver")
	codeLoginType := ctx.URLParam("codeLoginType")
	_, err := user.GetUserInfoByUserName(receiver)
	if err != nil {
		log.Debug("receiver：", receiver)
		resp.Code = tool.RespCodeError
		resp.Message = "验证码发送失败,用户不存在！"
		tool.ResponseJSON(ctx, resp)
		return
	}
	//生成4位验证码
	var code = common.RandCode(4)
	var value string
	switch codeLoginType {
	case string(model.SmsNotifyType):
		key := string(model.SmsSighIn) + receiver
		err = app.GetCache().Get(key, &value)
		if err == nil && len(value) > 0 {
			resp.Code = tool.RespCodeError
			resp.Message = "验证码已发送,请稍后重试！"
			tool.ResponseJSON(ctx, resp)
			return
		}
		log.Info("发送短信：%s", receiver)
		log.Debug("code:%s", code)
		sms, err := notify.SighInSendSms(code, receiver)
		if err != nil || sms.Code != float64(tool.RespCodeSuccess) {
			resp.Code = tool.RespCodeError
			resp.Message = "验证码发送失败,请稍后重试！"
			tool.ResponseJSON(ctx, resp)
			return
		}
		//缓存验证码 1分钟过期 key=邮箱
		err = app.GetCache().SetTTL(key, code, time.Minute)
		if err != nil {
			log.Info("缓存验证码失败：%s", err)
		}
		break
	case string(model.MailNotifyType):
		key := string(model.EmailSighIn) + receiver
		err = app.GetCache().Get(key, &value)
		if err == nil && len(value) > 0 {
			resp.Code = tool.RespCodeError
			resp.Message = "验证码已发送,请稍后重试！"
			tool.ResponseJSON(ctx, resp)
			return
		}
		log.Info("发送邮件：%s", receiver)
		log.Debug("code:%s", code)
		mail, err := notify.SighInSendMail(code, receiver)
		if err != nil || mail.Code != float64(tool.RespCodeSuccess) {
			resp.Code = tool.RespCodeError
			resp.Message = "验证码发送失败,请稍后重试！"
			tool.ResponseJSON(ctx, resp)
			return
		}
		//缓存验证码 1分钟过期 key=邮箱
		err = app.GetCache().SetTTL(key, code, time.Minute)
		if err != nil {
			log.Info("缓存验证码失败：%s", err)
		}
		break
	default:
		resp.Code = tool.RespCodeError
		resp.Message = "验证码未发送,未匹配到发送平台！"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

func signInSendCodePost(ctx iris.Context) {
	var resp app.Response
	var req signInReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	receiver := req.Receiver
	codeLoginType := req.CodeLoginType
	_, err = user.GetUserInfoByUserName(receiver)
	if err != nil {
		log.Debug("receiver：", receiver)
		resp.Code = tool.RespCodeError
		resp.Message = "验证码发送失败,用户不存在！"
		tool.ResponseJSON(ctx, resp)
		return
	}
	//生成4位验证码
	var code = common.RandCode(4)
	switch codeLoginType {
	case string(model.SmsNotifyType):
		key := string(model.SmsSighIn) + receiver
		exist := app.GetCache().Exists(key)
		if exist {
			resp.Code = tool.RespCodeError
			resp.Message = "验证码已发送,请稍后重试！"
			tool.ResponseJSON(ctx, resp)
			return
		}
		log.Info("发送短信：%s", receiver)
		log.Debug("code:%s", code)
		sms, err := notify.SighInSendSms(code, receiver)
		if err != nil || sms.Code != float64(tool.RespCodeSuccess) {
			resp.Code = tool.RespCodeError
			resp.Message = "验证码发送失败,请稍后重试！"
			tool.ResponseJSON(ctx, resp)
			return
		}
		//缓存验证码 1分钟过期 key=邮箱
		err = app.GetCache().SetTTL(key, code, time.Minute)
		if err != nil {
			log.Info("缓存验证码失败：%s", err)
		}
		break
	case string(model.MailNotifyType):
		key := string(model.EmailSighIn) + receiver
		exist := app.GetCache().Exists(key)
		if exist {
			resp.Code = tool.RespCodeError
			resp.Message = "验证码已发送,请稍后重试！"
			tool.ResponseJSON(ctx, resp)
			return
		}
		log.Info("发送邮件：%s", receiver)
		log.Debug("code:%s", code)
		mail, err := notify.SighInSendMail(code, receiver)
		if err != nil || mail.Code != float64(tool.RespCodeSuccess) {
			resp.Code = tool.RespCodeError
			resp.Message = "验证码发送失败,请稍后重试！"
			tool.ResponseJSON(ctx, resp)
			return
		}
		//缓存验证码 1分钟过期 key=邮箱
		err = app.GetCache().SetTTL(key, code, time.Minute)
		if err != nil {
			log.Info("缓存验证码失败：%s", err)
		}
		break
	default:
		resp.Code = tool.RespCodeError
		resp.Message = "验证码未发送,未匹配到发送平台！"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
