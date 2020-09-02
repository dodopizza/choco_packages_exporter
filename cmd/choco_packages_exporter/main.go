package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/eventlog"
	"github.com/dodopizza/choco_packages_exporter/internal/exportersvc"
	"github.com/dodopizza/choco_packages_exporter/internal/choco"
)

var (
	//svclog = debug.New("choco_packages_exporter")
	svclog, _ = eventlog.Open("choco_packages_exporter")
	chocoPackages = choco.New(svclog)
)

func getArguments() {
	show_version := flag.Bool("version", false, "Display an app version")
	show_help := flag.Bool("help", false, "Help information")
	svc_action := flag.String("service", "", "Service actions: install,remove,start,stop")
	appConfig.port = flag.Int("port", 9804, "Metrics port")
	flag.Parse()

	if *show_version == true {
		appConfig.getAppVersion()
		os.Exit(0)
	}

	if *show_help == true {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *svc_action != "" {
		switch strings.ToLower(*svc_action) {
		case "install":
			exportersvc.InstallService(appConfig.appName, appConfig.appName, "--port", strconv.Itoa(*appConfig.port))
			fmt.Println("Service installed")
		case "remove":
			exportersvc.ControlService(appConfig.appName, svc.Stop, svc.Stopped)
			fmt.Println("Service stopped")
			exportersvc.RemoveService(appConfig.appName)
			fmt.Println("Service removed")
		case "start":
			exportersvc.StartService(appConfig.appName)
			fmt.Println("Service started")
		case "stop":
			exportersvc.ControlService(appConfig.appName, svc.Stop, svc.Stopped)
			fmt.Println("Service stopped")
		default:
			flag.PrintDefaults()
		}
		os.Exit(0)
	}
}

func init() {
	fmt.Sprintf("Hello")
	getArguments()
}

func main() {
	if !(exportersvc.IsInteractiveSession(svclog)){
		exportersvc.ProvideService(svclog, &appConfig.appName)
		serveMetrics()
		exportersvc.WaitForServiceStoppedStatus(svclog, &appConfig.appName)
	} else {
		flag.PrintDefaults()
	}
}
