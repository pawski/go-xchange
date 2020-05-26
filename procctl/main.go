package procctl

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func RegisterSigTerm(shouldRun *bool) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func(shouldRun *bool) {
		<-s
		fmt.Println("Finishing remaining tasks...")
		*shouldRun = false
	}(shouldRun)
}
