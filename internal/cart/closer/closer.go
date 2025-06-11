package closer

import (
	"cart/internal/cart/logger"
	"context"
	"os"
	"os/signal"
	"sync"
)

type Closer interface {
	Add(f ...func() error)
	CloseAll()
	Wait()
	Signal()
}

type closer struct {
	sync.Mutex
	once     sync.Once
	funcs    []func() error
	done     chan struct{}
	shutdown chan os.Signal
}

func NewCloser(signals ...os.Signal) Closer {
	c := &closer{
		done:     make(chan struct{}),
		shutdown: make(chan os.Signal, 1),
	}

	if len(signals) > 0 {
		go func() {
			// теперь заданные сигналы записываются в shutdown
			signal.Notify(c.shutdown, signals...)
			// ожидает, когда придет сигнал
			<-c.shutdown
			// теперь все сигналы снова не записываются в shutdown
			signal.Stop(c.shutdown)
			c.CloseAll()
		}()
	}

	return c
}

func (c *closer) Add(f ...func() error) {
	c.Lock()
	c.funcs = append(c.funcs, f...)
	c.Unlock()
}

func (c *closer) CloseAll() {
	var once sync.Once
	once.Do(func() {
		ctx := context.Background()
		logger.Info(ctx, "graceful shutdown started")
		defer logger.Info(ctx, "graceful shutdown finished")

		defer close(c.done)

		c.Lock()
		funcs := c.funcs
		c.Unlock()

		for i := len(funcs) - 1; i >= 0; i-- {
			err := c.funcs[i]()
			if err != nil {
				logger.Errorf(ctx, "failed to close some func from shutdown: %v", err)
			}
		}
	})
}

func (c *closer) Wait() {
	<-c.done
}

func (c *closer) Signal() {
	c.shutdown <- os.Interrupt
}
