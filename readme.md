# CoolParallel

协程池，控制并发数，复用

### 安装

```sh
go get -u github.com/link-yundi/coolparallel
```

### 示例

```go
import (
	"github.com/link-yundi/ylog"
	"time"
)

func main() {
    p := NewParallelPool(8)
	task := func() {
		time.Sleep(2 * time.Second)
		ylog.Info(time.Now())
	}
	for i := 0; i < 32; i++ {
		p.AddTask(task)
	}
	p.Wait()
}
```

