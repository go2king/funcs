package gorunlimit

import (
	"sync"
)

type Monitor struct {
	routines chan int
	wg       *sync.WaitGroup
}

// NewMoni 获取一个实例
func NewMoni(size int) *Monitor {
	if size <= 0 {
		size = 0
	}
	return &Monitor{
		routines: make(chan int, size),
		wg:       &sync.WaitGroup{},
	}
}
func (t *Monitor) Add() {
	t.wg.Add(1)
	t.routines <- 1
}
func (t *Monitor) Done() {
	<-t.routines
	t.wg.Done()
}
func (t *Monitor) Wait() {
	t.wg.Wait()
}

type MonitorWrapper struct {
	moni *Monitor
}

// NewMoniWrap 生成一个管理实例
func NewMoniWrap(size int) *MonitorWrapper {
	return &MonitorWrapper{
		moni: NewMoni(size),
	}
}

// Wrap 执行体
func (t *MonitorWrapper) Wrap(cb func()) {
	t.moni.Add()
	go func() {
		defer t.moni.Done()
		cb()
	}()
}

// Wait 等待执行结束
func (t *MonitorWrapper) Wait() {
	t.moni.Wait()
}
//var genes =[]struct{Id:“id”，Name:“name”}
//mi := gorunlimit.NewMoniWrap(10)
//for _, gene := range genes {
//tmpid := gene.Id
//mi.Wrap(func() {
//proc(tmpid)
//})
//}
//mi.Wait()