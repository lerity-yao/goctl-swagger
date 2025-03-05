package main

import (
	"fmt"
	"github.com/lerity-yao/goctl-swagger/generate"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
)

func main() {
	fileName := "swaggerTest.json"
	host := "localhost:8080"
	basepath := "/v1"
	schemes := "https"

	var plugin plugin.Plugin

	var info struct {
		ApiFilePath string
		Style       string
		Dir         string
	}

	info.ApiFilePath = "./example/user.api"
	info.Style = "api"
	info.Dir = "./"

	plugin.ApiFilePath = info.ApiFilePath
	plugin.Style = info.Style
	plugin.Dir = info.Dir
	api, err := parser.Parse(info.ApiFilePath)
	plugin.Api = api

	if err != nil {
		fmt.Println(err, "create plugin")
	}

	err = generate.Do(fileName, host, basepath, schemes, &plugin)
	if err != nil {
		fmt.Println(err)
	}
}
