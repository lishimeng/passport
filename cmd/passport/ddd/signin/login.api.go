package signin

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/cmd/passport/ddd/signup"
	"github.com/lishimeng/passport/cmd/passport/ddd/user"
	"github.com/lishimeng/passport/internal/common"
	"github.com/lishimeng/passport/internal/db/model"
	"github.com/lishimeng/passport/internal/notify"
	"github.com/lishimeng/passport/internal/passwd"
	"time"
)

type LoginReq struct {
	Password  string `json:"password,omitempty"`
	UserName  string `json:"userName,omitempty"`
	Code      string `json:"code,omitempty"`
	LoginType string `json:"loginType,omitempty"` //登录方式-pc/app/wx
}

type CodeLoginReq struct {
	UserName      string `json:"userName,omitempty"`
	Code          string `json:"code,omitempty"`
	CodeLoginType string `json:"codeLoginType,omitempty"` //登录方式-sms/mail
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
	info, err := user.GetUserInfoByUserName(req.UserName)
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
	var provider token.JwtProvider
	err = factory.Get(&provider)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token 服务异常"
		tool.ResponseJSON(ctx, resp)
		return
	}
	tokenPayload := token.JwtPayload{
		Org:    "org_1",
		Dept:   "dep_1",
		Uid:    info.Code,
		Client: req.LoginType,
	}
	tokenContent, err := provider.Gen(tokenPayload)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token生成失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	resp.Token = string(tokenContent)
	tool.ResponseJSON(ctx, resp)
}

func codeLogin(ctx iris.Context) {
	var resp LoginResp
	var req CodeLoginReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	var value string
	err = app.GetCache().Get(req.UserName, &value)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "请先获取验证码"
		tool.ResponseJSON(ctx, resp)
		return
	}
	log.Info("code:%s,%s", value, req.Code)
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
	//验证码匹配激活邮箱或手机
	var cols []string
	switch req.CodeLoginType {
	case string(model.SmsNotifyType):
		info.MobileVerified = model.ActivateEnable
		cols = append(cols, "MobileVerified")
		info.Active = model.ActivateEnable
		cols = append(cols, "Active")
		_, err = signup.UpAccount(info, cols...)
	case string(model.MailNotifyType):
		info.EmailVerified = model.ActivateEnable
		cols = append(cols, "EmailVerified")
		info.Active = model.ActivateEnable
		cols = append(cols, "Active")
		_, err = signup.UpAccount(info, cols...)
	}
	resp.Code = tool.RespCodeSuccess
	var provider token.JwtProvider
	err = factory.Get(&provider)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token 服务异常"
		tool.ResponseJSON(ctx, resp)
		return
	}
	tokenPayload := token.JwtPayload{
		Org:    "org_1",
		Dept:   "dep_1",
		Uid:    info.Code,
		Client: req.CodeLoginType,
	}
	tokenContent, err := provider.Gen(tokenPayload)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token生成失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Token = string(tokenContent)
	tool.ResponseJSON(ctx, resp)
}

func sendCode(ctx iris.Context) {
	var resp app.Response
	mail := ctx.URLParam("mail")
	codeLoginType := ctx.URLParam("codeLoginType")
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
	switch codeLoginType {
	case string(model.SmsNotifyType):
		//todo
	case string(model.MailNotifyType):
		sendMail, err := notify.SighInSendMail(code, mail)
		if err != nil || sendMail.Code != float64(tool.RespCodeSuccess) {
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

type openLoginReq struct {
	SocialAccountId string `json:"socialAccountId,omitempty"`
	LoginType       string `json:"loginType,omitempty"`
}

func openLogin(ctx iris.Context) {
	var resp LoginResp
	var req openLoginReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	socialAccount, err := user.GetSocialAccountById(req.SocialAccountId)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "未绑定"
		tool.ResponseJSON(ctx, resp)
		return
	}
	account, err := user.GetUserInfoById(socialAccount.Id)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "未查到记录"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	var provider token.JwtProvider
	err = factory.Get(&provider)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token 服务异常"
		tool.ResponseJSON(ctx, resp)
		return
	}
	tokenPayload := token.JwtPayload{
		Org:    "org_1",
		Dept:   "dep_1",
		Uid:    account.Code,
		Client: req.LoginType,
	}
	tokenContent, err := provider.Gen(tokenPayload)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token生成失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Token = string(tokenContent)
	tool.ResponseJSON(ctx, resp)
	return
}
