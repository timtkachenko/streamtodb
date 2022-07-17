package shutdown

import (
	"github.com/google/logger"
	"os"
	"os/signal"
	"syscall"
)

func Graceful() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-quit
	logger.Info("Shutting down service...")
}
