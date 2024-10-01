package util

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GracefulShutdown(ctx context.Context, timeout time.Duration, operations map[string]func(context.Context) error) <-chan struct{} {
	wait := make(chan struct{})

	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		defer close(s)

		log.Println("shutting down application")

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %v sec has been elapsed, force exit\n", timeout.Seconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		for key, op := range operations {
			innerKey := key
			innerOp := op

			log.Printf("cleaning up: %v", innerKey)

			if err := innerOp(ctx); err != nil {
				log.Printf("%s: cleanin up failed: %e", innerKey, err)
				return
			}

			log.Printf("%s was shutdown gracefully", innerKey)
		}

		close(wait)
	}()

	return wait
}
