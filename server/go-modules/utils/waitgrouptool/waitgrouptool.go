package waitgrouptool

import (
	"sync"
)

type WaitGroupTool struct {
	Size    int
	current chan struct{}
	wg      sync.WaitGroup
}

func NewLimit(size int) WaitGroupTool {
	return WaitGroupTool{
		Size:    size,
		current: make(chan struct{}, size),
		wg:      sync.WaitGroup{},
	}
}

func (wgt *WaitGroupTool) Add() {
	select {
	case wgt.current <- struct{}{}:
		break
	}
	wgt.wg.Add(1)
}

func (wgt *WaitGroupTool) Done() {
	<-wgt.current
	wgt.wg.Done()
}

func (wgt *WaitGroupTool) Wait() {
	wgt.wg.Wait()
}
