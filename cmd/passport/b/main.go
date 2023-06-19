package main

import (
	"fmt"
	"github.com/lishimeng/app-starter/buildscript"
)

func main() {
	err := buildscript.Generate("passport",
		"lishimeng",
		"cmd/passport/main.go", true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
