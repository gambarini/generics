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
		Name   string
	}
)

func NewProcessManager(name string) *ProcessManager {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	return &ProcessManager{
		Ctx:    ctx,
		cancel: cancel,
		Name:   name,
	}
}

func (proc *ProcessManager) Start() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT)  // Handling Ctrl + C
	signal.Notify(sigChan, syscall.SIGTERM) // Handling Docker stop

	log.Printf("Process %s started. CTRL+C to stop", proc.Name)
	<-sigChan
	log.Printf("Process %s terminating...", proc.Name)
	proc.cancel()
	time.Sleep(time.Second)
	log.Println("Done.")
}
