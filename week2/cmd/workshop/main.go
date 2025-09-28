package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"time"
	"week2/logger"
)

type Counter struct {
	rw    sync.RWMutex
	value int
}

func (c *Counter) Increment() {
	c.rw.RLock()
	defer c.rw.RUnlock()

	c.value++
}
func (c *Counter) Get() int {
	c.rw.Lock()
	defer c.rw.Unlock()

	<-time.After(1 * time.Second)

	return c.value
}

func main() {
	ctx := context.Background()

	ctx = logger.ContextWithTimeStamp(ctx, time.Now())

	if err := run(ctx); err != nil {
		logger.Errorf(ctx, "%value", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	var cnt Counter

	go counter(ctx, &cnt)

	for i := 0; i < 10; i++ {
		go observe(ctx, &cnt)
	}

	<-ctx.Done()
	return ctx.Err()
}

func counter(ctx context.Context, c *Counter) {
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			c.Increment()
		case <-ctx.Done():
			return
		}
	}
}

func observe(ctx context.Context, c *Counter) {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			logger.Infof(ctx, "counter: %d", c.Get())
		case <-ctx.Done():
			return
		}
	}
}
