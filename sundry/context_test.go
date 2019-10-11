package sundry

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWork(t *testing.T){
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	fmt.Println("start work.")
	var wg = new(sync.WaitGroup)
	wg.Add(1) // 控制routine完成工作
	go work(ctx, wg)
	wg.Wait()

	fmt.Println("end work.")
}
