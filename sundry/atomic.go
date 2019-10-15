package sundry

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// https://zhuanlan.zhihu.com/p/31122953
// 原子操作能够确保任一时刻只有一个goroutine对变量进行操作，善用 atomic 能有效避免程序中出现大量的锁sync.Mutex操作
// 原子操作：1、增或减 2、比较并交换 3、载入 4、存储 5、交换

func OpAtomic(){
	var wg sync.WaitGroup
	var a int32 = 52
	for i := 0; i < 10; i++{
		wg.Add(1)
		go func(a *int32) {
			defer wg.Done()
			atomic.AddInt32(a, 1)
		}(&a)
		//go func(a int32){
		//	atomic.StoreInt32(&a, 10)
		//}(a)
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&a))
}