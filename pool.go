package coolparallel

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
	work chan func(arg any) // 任务
	num  chan int           // 数量
	wg   *sync.WaitGroup
}

func NewParallelPool(size int) *ParallelPool {
	return &ParallelPool{
		work: make(chan func(arg any)),
		num:  make(chan int, size),
		wg:   &sync.WaitGroup{},
	}
}

// 往协程池中添加任务
func (pp *ParallelPool) AddTask(task func(arg any), arg any) {
	select {
	case pp.work <- task:
		pp.wg.Add(1)
	case pp.num <- 1:
		pp.wg.Add(1)
		go pp.worker(task, arg)
	}
}

// 执行任务
func (pp *ParallelPool) worker(task func(arg any), arg any) {
	for {
		task(arg)
		pp.wg.Done()
		task = <-pp.work
	}
}

// 关闭通道
func (pp *ParallelPool) close() {
	close(pp.work)
	close(pp.num)
}

func (pp *ParallelPool) Wait() {
	pp.wg.Wait()
	pp.close()
}
