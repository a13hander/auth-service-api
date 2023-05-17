package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
)

var closer = newCloser()

type closerCb struct {
	mu    sync.Mutex
	once  sync.Once
	done  chan struct{}
	funcs []func() error
}

// New returns new closerCb, if []os.Signal is specified closerCb will automatically call CloseAll when one of signals is received from OS
func newCloser(sig ...os.Signal) *closerCb {
	c := &closerCb{done: make(chan struct{})}
	if len(sig) > 0 {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, sig...)
			<-ch
			signal.Stop(ch)
			c.closeAll()
		}()
	}
	return c
}

// Add func to closer
func (c *closerCb) add(f ...func() error) {
	c.mu.Lock()
	c.funcs = append(c.funcs, f...)
	c.mu.Unlock()
}

// Wait blocks until all closer functions are done
func (c *closerCb) wait() {
	<-c.done
}

// CloseAll calls all closer functions
func (c *closerCb) closeAll() {
	c.once.Do(func() {
		defer close(c.done)

		c.mu.Lock()
		funcs := c.funcs
		c.funcs = nil
		c.mu.Unlock()

		// call all closerCb funcs async
		errs := make(chan error, len(funcs))
		for _, f := range funcs {
			go func(f func() error) {
				errs <- f()
			}(f)
		}

		for i := 0; i < cap(errs); i++ {
			if err := <-errs; err != nil {
				log.Println(context.Background(), "error returned from closerCb: %v", err)
			}
		}
	})
}
