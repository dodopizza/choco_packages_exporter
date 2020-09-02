package exportersvc

import (
	"fmt"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/eventlog"
)

var svcStopCh = make(chan bool)

type ChocoPackagesExporterService struct {
	stopCh chan<- bool
	logger *eventlog.Log
}

func (s *ChocoPackagesExporterService) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	changes <- svc.Status{State: svc.StartPending}
	changes <- svc.Status{State: svc.Running, Accepts: svc.AcceptStop | svc.AcceptShutdown}
loop:
	for {
		select {
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				changes <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				s.stopCh <- true
				break loop
			default:
				s.logger.Error(1, fmt.Sprintf("unexpected control request #%d", c))
			}
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return
}

func IsInteractiveSession(logger *eventlog.Log) bool {
	isInteractive, err := svc.IsAnInteractiveSession()
	if err != nil {
		logger.Error(1, fmt.Sprintf("Failed to determine if we are running in an interactive session: %v", err))
	}
	return isInteractive
}

func ProvideService(logger *eventlog.Log, appName *string){
	go func() {
		err := svc.Run(*appName, &ChocoPackagesExporterService{stopCh: svcStopCh, logger:logger})
		if err != nil {
			logger.Error(1, fmt.Sprintf("Failed to start service: %v", err))
		}
	}()
}

func WaitForServiceStoppedStatus(logger *eventlog.Log, appName *string){
	for {
		if <-svcStopCh {
			logger.Error(1, "Shutting down " + *appName)
			break
		}
	}
}