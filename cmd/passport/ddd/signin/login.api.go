package signin

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/cmd/passport/ddd/user"
	"github.com/lishimeng/passport/internal/common"
	"github.com/lishimeng/passport/internal/db/model"
	"github.com/lishimeng/passport/internal/etc"
	"github.com/lishimeng/passport/internal/gentoken"
	"github.com/lishimeng/passport/internal/passwd"
	"github.com/lishimeng/passport/internal/sendmessage"
	"time"
)

type LoginReq struct {
	Password  string `json:"password,omitempty"`
	UserName  string `json:"userName,omitempty"`
	Code      string `json:"code,omitempty"`
	LoginType string `json:"loginType,omitempty"`
}

type LoginResp struct {
	app.Response
	Token string `json:"token,omitempty"`
	Uid   int    `json:"uid,omitempty"`
}

func login(ctx iris.Context) {
	var resp LoginResp
	var req LoginReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	info, err := user.GetUserInfoByName(req.UserName)
	log.Debug("info:%s", info)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "用户名或密码错误"
		tool.ResponseJSON(ctx, resp)
		return
	}
	p := passwd.Verify(req.Password, info)
	log.Info("password:%s,%s", info.Password, p)
	if !p {
		resp.Code = tool.RespCodeError
		resp.Message = "用户名或密码错误"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	resp.Uid = info.Id
	log.Info("Uid:%s", info.Id)
	token, err := gentoken.GenToken("", string(info.Id), req.LoginType)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token生成失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Token = token
	tool.ResponseJSON(ctx, resp)
}

func codeLogin(ctx iris.Context) {
	var resp LoginResp
	var req LoginReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	//登录方式 sms-短信验证，mail-邮箱验证
	var value string
	err = app.GetCache().Get(req.UserName, &value)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "请先获取验证码"
		tool.ResponseJSON(ctx, resp)
		return
	}
	log.Info("Uid:%s", value, req.Code)
	if value != req.Code {
		resp.Code = tool.RespCodeError
		resp.Message = "验证码不正确"
		tool.ResponseJSON(ctx, resp)
		return
	}
	info, err := user.GetUserInfoByUserName(req.UserName)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "请先绑定邮箱/手机"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	resp.Uid = info.Id
	log.Info("Uid:%s", info.Id)
	token, err := gentoken.GenToken("", string(info.Id), req.LoginType)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token生成失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Token = token
	tool.ResponseJSON(ctx, resp)
}

func sendCode(ctx iris.Context) {
	var resp app.Response
	mail := ctx.URLParam("mail")
	loginType := ctx.URLParam("loginType")
	var value string
	err := app.GetCache().Get(mail, &value)
	if err != nil {
		log.Info("获取缓存验证码失败%s：%s", mail, err)
	}
	//生成4位验证码
	var code = common.RandCode(4)
	log.Debug("code:%s", value)
	if len(value) > 0 {
		resp.Code = tool.RespCodeError
		resp.Message = "验证码已发送,请稍后重试！"
		tool.ResponseJSON(ctx, resp)
		return
	}
	switch loginType {
	case string(model.SmsNotifyType):
		//todo
	case string(model.MailNotifyType):
		var sms = sendmessage.Req{
			Template:      etc.Config.Notify.MailTemplate,
			TemplateParam: "{\"verrificationCode\":\"" + code + "\"}",
			Subject:       "验证码",
			Receiver:      mail,
		}
		response, err := sendmessage.SendMail(sms)
		if err != nil || response.Code != float64(tool.RespCodeSuccess) {
			resp.Code = tool.RespCodeError
			resp.Message = "验证码发送失败,请稍后重试！"
			tool.ResponseJSON(ctx, resp)
			return
		}
	}
	//缓存验证码 3分钟过期 key=邮箱
	err = app.GetCache().SetTTL(mail, code, time.Minute*3)
	if err != nil {
		log.Info("缓存验证码失败：%s", err)
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
