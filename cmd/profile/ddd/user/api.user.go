package user

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/internal/db/model"
)

type UserResp struct {
	app.Response
	Item UserInfo `json:"item"`
}

type UserInfo struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Phone  string `json:"phone,omitempty"`
	Active int    `json:"active,omitempty"`
	IsOrg  int    `json:"isOrg,omitempty"`
	IsBind int    `json:"isBind,omitempty"`
}

func GetUserInfo(ctx iris.Context) {
	var resp UserResp
	ui := ctx.Values().Get(auth.UserInfoKey)
	jwt, ok := ui.(token.JwtPayload)
	if !ok {
		return
	}
	log.Debug("code:%s", jwt.Uid)
	account, err := GetUserInfoByCode(jwt.Uid)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "未查到记录"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	resp.Item.Id = account.Id
	resp.Item.Name = account.Name
	resp.Item.Email = account.Email
	resp.Item.Active = account.Active
	resp.Item.Phone = account.Mobile
	resp.Item.IsOrg = 0
	resp.Item.IsBind = 0
	_, err = GetTenantAccountByUid(account.Id)
	if err == nil {
		resp.Item.IsOrg = 1
	}
	_, err = GetSocialAccountByAccountId(account.Id)
	if err == nil {
		resp.Item.IsBind = 1
	}
	tool.ResponseJSON(ctx, resp)
	return
}

type SocialReq struct {
	SocialAccountId string `json:"socialAccountId,omitempty"`
	SocialGroupId   string `json:"socialGroupId,omitempty"`
	Category        string `json:"category,omitempty"`
}

func BindUser(ctx iris.Context) {
	var req SocialReq
	var resp app.Response
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	ui := ctx.Values().Get(auth.UserInfoKey)
	jwt, ok := ui.(token.JwtPayload)
	if !ok {
		return
	}
	log.Debug("code:%s", jwt.Uid)
	account, err := GetUserInfoByCode(jwt.Uid)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "未查到记录"
		tool.ResponseJSON(ctx, resp)
		return
	}
	_, err = GetSocialAccountById(req.SocialAccountId, req.Category, account.Id)
	if err == nil {
		resp.Code = tool.RespCodeError
		resp.Message = req.SocialAccountId + "已绑定"
		tool.ResponseJSON(ctx, resp)
		return
	}
	var socialAccount = model.SocialAccount{
		AccountId:       account.Id,
		SocialAccountId: req.SocialAccountId,
		SocialGroupId:   req.SocialGroupId,
		Category:        model.SocialCategory(req.Category),
	}
	_, err = InsertSocialAccount(socialAccount)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "绑定失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

type BindPhoneReq struct {
	Mobile string `json:"mobile,omitempty"`
	Code   string `json:"code,omitempty"`
}

func BindPhone(ctx iris.Context) {
	var req BindPhoneReq
	var resp app.Response
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	ui := ctx.Values().Get(auth.UserInfoKey)
	jwt, ok := ui.(token.JwtPayload)
	if !ok {
		return
	}
	log.Debug("userCode:%s", jwt.Uid)
	account, err := GetUserInfoByCode(jwt.Uid)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "用户不存在"
		tool.ResponseJSON(ctx, resp)
		return
	}
	key := string(model.SmsBind) + req.Mobile
	var value string
	err = app.GetCache().Get(key, &value)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "验证码未发送，请重新发送！"
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
	var cols []string
	account.Mobile = req.Mobile
	cols = append(cols, "Mobile")
	account.MobileVerified = model.ActivateEnable
	cols = append(cols, "MobileVerified")
	account.Active = model.ActivateEnable
	cols = append(cols, "Active")
	_, err = UpAccount(account, cols...)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "绑定失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

type BindEmailReq struct {
	Email string `json:"email,omitempty"`
	Code  string `json:"code,omitempty"`
}

func BindEmail(ctx iris.Context) {
	var req BindEmailReq
	var resp app.Response
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	ui := ctx.Values().Get(auth.UserInfoKey)
	jwt, ok := ui.(token.JwtPayload)
	if !ok {
		return
	}
	log.Debug("userCode:%s", jwt.Uid)
	account, err := GetUserInfoByCode(jwt.Uid)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "用户不存在"
		tool.ResponseJSON(ctx, resp)
		return
	}
	key := string(model.EmailBind) + req.Email
	var value string
	err = app.GetCache().Get(key, &value)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "验证码未发送，请重新发送！"
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
	var cols []string
	account.Email = req.Email
	cols = append(cols, "Email")
	account.EmailVerified = model.ActivateEnable
	cols = append(cols, "EmailVerified")
	account.Active = model.ActivateEnable
	cols = append(cols, "Active")
	_, err = UpAccount(account, cols...)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "绑定失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
