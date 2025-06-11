package db

import (
	"context"
	"errors"
	"sync/atomic"
)

type Client interface {
	GetReaderPool() Pool
	GetWriterPool() Pool
	GetMasterPool() Pool
	GetSyncPool() Pool
	Close() error
}

type client struct {
	ctx        context.Context
	masterPool Pool
	syncPool   Pool
	counter    atomic.Uint64
}

func NewClient(ctx context.Context, masterDsn string, syncDsn string) (Client, error) {
	masterPool, err := NewPool(ctx, masterDsn)
	if err != nil {
		return nil, errors.New("failed to connect to master: " + err.Error())
	}

	syncPool, err := NewPool(ctx, syncDsn)
	if err != nil {
		return nil, errors.New("failed to connect to sync pool: " + err.Error())
	}

	return &client{
		ctx:        ctx,
		masterPool: masterPool,
		syncPool:   syncPool,
	}, nil
}

func (c *client) GetReaderPool() Pool {
	res := c.counter.Add(1)
	if res%2 == 0 {
		return c.masterPool
	}

	return c.syncPool
}

func (c *client) GetWriterPool() Pool {
	return c.masterPool
}

func (c *client) GetMasterPool() Pool {
	return c.masterPool
}

func (c *client) GetSyncPool() Pool {
	return c.syncPool
}

func (c *client) Close() error {
	err := c.masterPool.Close()
	if err != nil {
		return err
	}

	err = c.syncPool.Close()
	if err != nil {
		return err
	}

	return nil
}
