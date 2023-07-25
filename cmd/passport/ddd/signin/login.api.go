package signin

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/cmd/passport/ddd/user"
	"github.com/lishimeng/passport/internal/db/model"
	"github.com/lishimeng/passport/internal/passwd"
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
	LoginType     string `json:"loginType,omitempty"`     //登录方式-pc/app/wx
}

type LoginResp struct {
	app.Response
	Token string `json:"token,omitempty"`
	Uid   int    `json:"uid,omitempty"`
}

func getToken(uid int, uCode, loginType string) (tokenVal []byte, err error) {
	var tokenPayload token.JwtPayload
	tokenPayload.Uid = uCode
	tokenPayload.Client = loginType
	tenantAccount, err := user.GetTenantAccountByUid(uid)
	if err != nil {
		// 不在组织里
		log.Debug("不在组织里:%s", uCode)
		tokenVal, err = genToken(tokenPayload)
		return
	}
	tenant, err := user.GetTenantById(tenantAccount.Org)
	if err == nil {
		tokenPayload.Org = tenant.Code
	} else {
		log.Debug("组织不存在, id:%d", tenantAccount.Org)
	}
	tokenVal, err = genToken(tokenPayload)
	return
}

func genToken(payload token.JwtPayload) (content []byte, err error) {
	var provider token.JwtProvider
	err = factory.Get(&provider)
	if err != nil {
		return
	}
	log.Info("tokenPayload: %s", payload)
	content, err = provider.Gen(payload)
	return
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
	tokenContent, err := getToken(info.Id, info.Code, req.LoginType)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token获取失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	//cache token
	go func() {
		_ = saveToken(tokenContent)
	}()
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
	info, err := user.GetUserInfoByUserName(req.UserName)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "请先绑定邮箱/手机"
		tool.ResponseJSON(ctx, resp)
		return
	}
	//验证码匹配激活邮箱或手机
	var cols []string
	var value string
	switch req.CodeLoginType {
	case string(model.SmsNotifyType):
		key := string(model.SmsSighIn) + req.UserName
		err = app.GetCache().Get(key, &value)
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
		info.MobileVerified = model.ActivateEnable
		cols = append(cols, "MobileVerified")
		info.Active = model.ActivateEnable
		cols = append(cols, "Active")
		_, err = user.UpAccount(info, cols...)
		//登录成功时清除验证码缓存
		app.GetCache().Set(key, "")
	case string(model.MailNotifyType):
		key := string(model.EmailSighIn) + req.UserName
		err = app.GetCache().Get(key, &value)
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
		info.EmailVerified = model.ActivateEnable
		cols = append(cols, "EmailVerified")
		info.Active = model.ActivateEnable
		cols = append(cols, "Active")
		_, err = user.UpAccount(info, cols...)
		//登录成功时清除验证码缓存
		app.GetCache().Set(key, "")
	default:
		resp.Code = tool.RespCodeError
		resp.Message = "未匹配到登录平台"
		tool.ResponseJSON(ctx, resp)
		return
	}
	tokenContent, err := getToken(info.Id, info.Code, req.LoginType)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token获取失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	//cache token
	go func() {
		_ = saveToken(tokenContent)
	}()
	resp.Code = tool.RespCodeSuccess
	resp.Token = string(tokenContent)
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

	socialAccount, err := user.GetSocialAccountById(req.SocialAccountId, req.LoginType, 0)
	if err != nil {
		go func() {
			// TODO save social account
		}()
		resp.Code = tool.RespCodeError
		resp.Message = "未绑定"
		tool.ResponseJSON(ctx, resp)
		return
	}
	account, err := user.GetUserInfoById(socialAccount.AccountId)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "未查到记录"
		tool.ResponseJSON(ctx, resp)
		return
	}
	tokenContent, err := getToken(account.Id, account.Code, req.LoginType)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "token获取失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	//cache token
	go func() {
		_ = saveToken(tokenContent)
	}()
	resp.Code = tool.RespCodeSuccess
	resp.Uid = account.Id
	resp.Token = string(tokenContent)
	tool.ResponseJSON(ctx, resp)
	return
}
