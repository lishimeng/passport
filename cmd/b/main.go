package main

import (
	"fmt"
	"github.com/lishimeng/app-starter/buildscript"
)

func main() {
	err := buildscript.Generate("lishimeng",
		buildscript.Application{
			Name:    "passport",
			AppPath: "cmd/passport",
			HasUI:   true,
		},
		buildscript.Application{
			Name:    "passport-admin",
			AppPath: "cmd/admin",
			HasUI:   false,
		})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
