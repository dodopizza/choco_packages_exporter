package choco

import (
	"golang.org/x/sys/windows/svc/debug"
	"golang.org/x/sys/windows/svc/eventlog"
	"os/exec"
	"regexp"
	"runtime"
)

type ChocoPackages struct {
	logger *eventlog.Log
	list   []ChocoPackageInfo
}

type ChocoPackageInfo struct {
	name    string
	version string
}

func New(logger *eventlog.Log) *ChocoPackages {
	return &ChocoPackages{
		logger: logger,
	}
}

func (c ChocoPackageInfo) GetName() string {
	return c.name
}

func (c ChocoPackageInfo) GetVersion() string {
	return c.version
}

func (c ChocoPackages) GetPackages() []ChocoPackageInfo{
	updatePackages(&c)
	return c.list
}

func updatePackages(packages *ChocoPackages) {
	packages.logger.Info(1,"Updating choco packages info")
	packages.list = []ChocoPackageInfo{}
	var packagesStr string
	if runtime.GOOS == "windows" {
		packagesStr = executeCommandWithOutput(packages.logger, getChocoExecutable(), "list", "-lo", "-r")
	} else {
		packages.logger.Info(1,"This exporter was designed for Windows OS. Using example output instead of real choco")
		packagesStr = executeCommandWithOutput(packages.logger, "cat", "examples/choco_packages_output_demo.txt")
	}
	for _, match := range extractPackageInfoFromPackagesMultilineString(packagesStr) {
		var p = ChocoPackageInfo{
			name:    match[1],
			version: match[2],
		}
		packages.list = append(packages.list, p)
	}
}

func getChocoExecutable() string {
	ps, _ := exec.LookPath("choco.exe")
	return ps
}

func extractPackageInfoFromPackagesMultilineString(str string) [][]string {
	var re = regexp.MustCompile(`/^(?m)(?P<name>.+)\|[v ]?(?P<version>[\d.]+)$/m`)
	return re.FindAllStringSubmatch(str, -1)
}

func executeCommandWithOutput(logger debug.Log, cmd string, arg ...string) string {
	logger.Info(1, "Executing " + cmd + " command")
	out, err := exec.Command(cmd, arg...).Output()
	if err != nil {
		logger.Error( 1, "executeCommandWithOutput - Error executing command " + err.Error() )
	}
	return string(out)
}