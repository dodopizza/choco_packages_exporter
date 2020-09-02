# Choco packages exporter

![Choco packages exporter](https://github.com/dodopizza/choco_packages_exporter/workflows/Mainline/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/dodopizza/choco_packages_exporter)](https://goreportcard.com/report/github.com/dodopizza/choco_packages_exporter)
[![codecov](https://codecov.io/gh/dodopizza/choco_packages_exporter/branch/master/graph/badge.svg)](https://codecov.io/gh/dodopizza/choco_packages_exporter)
![GitHub Release](https://img.shields.io/github/release/dodopizza/choco_packages_exporter.svg?style=flat)

Prometheus exporter for Chocolatey packages

## Metrics

`Choco packages exporter` exposes only one metric `winserver_choco_packages` with all installed packages and their versions as labels by pattern `package_name`=`package_version`.  Value of this metric is always `0`

Metrics are avaliable on `<exporter host>:<port>/metrics`

## Startup options

Name | Descrition
------------ | -------------
--help | Help information
--port | Metrics port (default 9804)
--service | Service actions: install,remove,start,stop
--version | Display an app version

## Installing as Windows service

1. Download release binary to desired exporter directory (eg: C:\choco_packages_exporter)
2. Run terminal (cmd, powershell) as administrator and execute the binary to read avaliable options ( eg: `.\choco_packages_exporter.exe --service install --port 1234` )
3. Done
