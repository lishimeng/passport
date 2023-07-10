package theme

// ThemeConfig
// 动态修改页面配置

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport/internal/db/model"
	"strconv"
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

type configResp struct {
	app.Response
	Data interface{} `json:"data,omitempty"`
}

func themeConfig(ctx iris.Context) {
	var resp configResp
	configPage := ctx.URLParamDefault("configPage", "")
	log.Info("configPage:%s", configPage)
	var themeConfigs []model.ThemeConfig
	qs := app.GetOrm().Context.QueryTable(new(model.ThemeConfig))
	if len(configPage) > 0 {
		_, err := qs.Filter("ConfigPage", configPage).All(&themeConfigs)
		if err != nil {
			log.Info("query fail : %s", err)
		}
	} else {
		_, err := qs.All(&themeConfigs)
		if err != nil {
			log.Info("query fail : %s", err)
		}
	}
	config := make(map[string]interface{})
	if len(themeConfigs) > 0 {
		for _, item := range themeConfigs {
			switch item.ConfigContentType {
			case string(model.NumberConfigContentType):
				config[item.ConfigName], _ = strconv.Atoi(item.ConfigContent)
				break
			case string(model.BooleanConfigContentType):
				config[item.ConfigName], _ = strconv.ParseBool(item.ConfigContent)
				break
			default:
				config[item.ConfigName] = item.ConfigContent
				break
			}
		}
	}
	resp.Data = config
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
