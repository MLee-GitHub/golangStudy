package sundry

import (
	"context"
	"fmt"
	"sync"
	"time"
)


func work(ctx context.Context, wg *sync.WaitGroup) error{
	defer wg.Done() // wg应为引用类型（或者不穿入wg，使用闭包类型直接使用外部wg变量），如若不然，只是wg的副本Done，而原wg将一直Wait，从而出现deadlock

	for i := 0; i <= 1000; i++{
		select {
		case <-time.After(time.Second * 1):
			fmt.Println("res:", i)
			case <-ctx.Done(): // 1.调用方主动cancel 2.达到deadline 3.父context的done关闭
				fmt.Println("res(done):", i)
				return ctx.Err()
		}
	}
	return nil
}




