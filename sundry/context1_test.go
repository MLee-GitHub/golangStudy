package sundry

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestContext1(t *testing.T) {
	fmt.Println("start read...")
	conn := &ServerConn{
		sendCh:   make(chan Message),
		handleCh: make(chan Message),
		wg:       &sync.WaitGroup{},
		netId:    100,
	}

	conn.ctx, conn.cancel = context.WithCancel(context.WithValue(context.Background(), "key", conn.netId))
	loopers := []func(*ServerConn, *sync.WaitGroup){readLoop, writeLoop, handleLoop}

	for _, looper := range loopers {
		conn.wg.Add(1)
		go looper(conn, conn.wg) // 开routine
	}

	go func() {
		time.Sleep(time.Second * 3)
		conn.cancel()
	}()

	conn.wg.Wait()
	fmt.Println("send end.")
}
