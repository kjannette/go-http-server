package os

import (
	"context"
	"os"
	"syscall"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestSignalListener_CancelsOnSIGTERM(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	canceled := make(chan struct{})
	go func() {
		<-ctx.Done()
		close(canceled)
	}()

	SignalListener(zap.NewNop(), cancel)

	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf("FindProcess: %v", err)
	}
	if err := proc.Signal(syscall.SIGTERM); err != nil {
		t.Fatalf("Signal SIGTERM: %v", err)
	}

	select {
	case <-canceled:
	case <-time.After(2 * time.Second):
		t.Fatal("context was not canceled after SIGTERM")
	}
}
