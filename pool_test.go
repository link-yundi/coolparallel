package coolparallel

import (
	"github.com/link-yundi/ylog"
	"testing"
	"time"
)

/**
------------------------------------------------
Created on 2022-11-03 23:52
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

// ========================== 协程池 ==========================
func TestPool(t *testing.T) {
	p := NewParallelPool(8)
	task := func(arg any) {
		time.Sleep(2 * time.Second)
		ylog.Info(time.Now())
	}
	for i := 0; i < 32; i++ {
		p.AddTask(task, nil)
	}
	p.Wait()
}
