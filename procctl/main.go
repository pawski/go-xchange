package procctl

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func RegisterSigTerm() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		<-s
		fmt.Println("Shutting down gracefully.")
		// clean up here
		os.Exit(0)
	}()
}
