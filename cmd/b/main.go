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
			HasUI:   false,
		},
		buildscript.Application{
			Name:    "passport-profile",
			AppPath: "cmd/profile",
			HasUI:   true,
		},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
