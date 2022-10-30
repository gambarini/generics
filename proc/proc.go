package proc

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type (
	ProcessManager struct {
		Ctx    context.Context
		cancel context.CancelFunc
	}
)

func NewProcessManager() *ProcessManager {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	return &ProcessManager{
		Ctx:    ctx,
		cancel: cancel,
	}
}

func (proc *ProcessManager) Start() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT)  // Handling Ctrl + C
	signal.Notify(sigChan, syscall.SIGTERM) // Handling Docker stop

	log.Println("Process started. CTRL+C to stop")
	<-sigChan
	log.Println("Process terminating...")
	proc.cancel()
	time.Sleep(time.Second)
	log.Println("Done.")
}
