package theme

// ThemeConfig
// 动态修改页面配置

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter/tool"
)

type IndexData struct {
	Meta
	Header
	Footer
}

type Meta struct {
	Title    string `json:"title,omitempty"`
	Keywords string `json:"keywords,omitempty"`
}

// Header 顶部
type Header struct {
	GlobalTitle string `json:"globalTitle,omitempty"`
	Login
	Logo
}

type Logo struct {
	LogoMini string `json:"logoMini,omitempty"`
}

// Login 登录界面配置
type Login struct {
	GlobalViceTitle    string `json:"globalViceTitle,omitempty"`
	GlobalViceTitleMsg string `json:"globalViceTitleMsg,omitempty"`
}

type Footer struct {
	FooterName string `json:"footerName,omitempty"`
	CopyRight  string `json:"copyRight,omitempty"`
}

func themeConfig(ctx iris.Context) {
	var resp IndexData

	// TODO 从db获取配置, 如没有值设置"", 确保response中没有值
	// TODO ui 根据response替换相应属性, 没有值的属性继续使用默认值
	tool.ResponseJSON(ctx, resp)
}
