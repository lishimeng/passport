package signup

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/cmd/passport/ddd/user"
	"github.com/lishimeng/passport/internal/passwd"
)

type RegisterReq struct {
	Mobile   string `json:"mobile,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}
type phoneRegisterReq struct {
	Mobile string `json:"mobile,omitempty"`
	Code   string `json:"code,omitempty"`
}

func phoneRegister(ctx iris.Context) {
	var resp app.Response
	var req phoneRegisterReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	_, err = user.GetUserInfoByUserName(req.Mobile)
	if err == nil {
		resp.Code = tool.RespCodeError
		resp.Message = "注册失败,手机号已被占用"
		tool.ResponseJSON(ctx, resp)
		return
	}
	var value string
	err = app.GetCache().Get(req.Mobile, &value)
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
	_, err = RegisterAccount(req.Mobile, "", "", "")
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "注册失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

type emailRegisterReq struct {
	Email string `json:"email,omitempty"`
	Code  string `json:"code,omitempty"`
}

func emailRegister(ctx iris.Context) {
	var resp app.Response
	var req emailRegisterReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	_, err = user.GetUserInfoByUserName(req.Email)
	if err == nil {
		resp.Code = tool.RespCodeError
		resp.Message = "注册失败,邮箱已被占用"
		tool.ResponseJSON(ctx, resp)
		return
	}
	var value string
	err = app.GetCache().Get(req.Email, &value)
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
	_, err = RegisterAccount("", req.Email, "", "")
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "注册失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
func register(ctx iris.Context) {
	var resp app.Response
	var req RegisterReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	_, err = user.GetUserInfoByThree(req.Name, req.Mobile, req.Email)
	if err == nil {
		resp.Code = tool.RespCodeError
		resp.Message = "注册失败,用户名/邮箱/手机号已被使用"
		tool.ResponseJSON(ctx, resp)
		return
	}
	info, erri := RegisterAccount(req.Mobile, req.Email, req.Password, req.Name)
	if erri != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "注册失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	account, erru := user.GetUserInfoById(info.Id)
	if erru != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "未查到记录"
		tool.ResponseJSON(ctx, resp)
		return
	}
	p := passwd.Generate(req.Password, account)
	account.Password = p
	_, err = user.UpAccount(account, "password")
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "更新密码失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
