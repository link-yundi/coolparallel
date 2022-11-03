package coolmq

import "sync"

/**
------------------------------------------------
Created on 2022-11-03 22:25
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

// ========================== 协程池 ==========================
type ParallelPool struct {
	work chan func() // 任务
	num  chan int    // 数量
	wg   *sync.WaitGroup
}

func NewParallelPool(size int) *ParallelPool {
	return &ParallelPool{
		work: make(chan func()),
		num:  make(chan int, size),
		wg:   &sync.WaitGroup{},
	}
}

// 往协程池中添加任务
func (pp *ParallelPool) AddTask(task func()) {
	select {
	case pp.work <- task:
		pp.wg.Add(1)
	case pp.num <- 1:
		pp.wg.Add(1)
		go pp.worker(task)
	}
}

// 执行任务
func (pp *ParallelPool) worker(task func()) {
	for {
		task()
		pp.wg.Done()
		task = <-pp.work
	}
}

func (pp *ParallelPool) Close() {
	close(pp.work)
	close(pp.num)
}

func (pp *ParallelPool) Wait() {
	pp.wg.Wait()
}
