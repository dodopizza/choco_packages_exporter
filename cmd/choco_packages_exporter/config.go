package main

import (
	"fmt"
	"strconv"
)

var appConfigVersion = "0.0.000" // go build -ldflags "-X main.appConfigVersion=1.2.345"
var appConfig = AppConfig{
	appName: "choco_packages_exporter",
	version: appConfigVersion,
}

type AppConfig struct {
	appName string
	version string
	port    *int
	debug   *bool
}

func (a *AppConfig) getHttpAddress() string {
	return ":" + strconv.Itoa(*a.port)
}

func (a *AppConfig) getAppVersion() {
	fmt.Println(a.version)
}
