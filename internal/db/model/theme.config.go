package model

import "github.com/lishimeng/app-starter"

// ThemeConfig 配置
type ThemeConfig struct {
	app.Pk
	ConfigPage        string `orm:"column(config_page);null"`      //配置所属页
	ConfigTheme       string `orm:"column(config_theme);null"`     //配置字段名称
	ConfigName        string `orm:"column(config_name);null"`      //配置字段名称
	ConfigContent     string `orm:"column(config_content);null"`   //配置字段内容
	ConfigContentType string `orm:"column(config_name_Type);null"` //配置字段内容类型
	app.TableChangeInfo
}

type ThemeConfigContentType string

const (
	BooleanConfigContentType ThemeConfigContentType = "boolean"
	NumberConfigContentType  ThemeConfigContentType = "int"
	StringConfigContentType  ThemeConfigContentType = "string"
)
