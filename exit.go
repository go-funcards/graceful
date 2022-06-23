package graceful

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Exit(callback func()) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	signal.Stop(quit)
	fmt.Println()
	fmt.Println("CTRL-C command received. Exiting...")

	if callback != nil {
		callback()
	}

	fmt.Println("Gracefully exited.")
	os.Exit(0)
}
