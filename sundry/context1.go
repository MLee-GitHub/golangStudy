package sundry

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Message struct {
	netId int
	Data  string
}

type ServerConn struct {
	sendCh   chan Message
	handleCh chan Message
	wg       *sync.WaitGroup
	ctx      context.Context
	cancel   context.CancelFunc
	netId    int
}

// 读消息
func readLoop(c *ServerConn, wg *sync.WaitGroup) {
	netId, _ := c.ctx.Value("key").(int)
	handlerCh := c.handleCh
	//ctx, _ := context.WithCancel(c.ctx)
	//cDone := ctx.Done()
	cDone := c.ctx.Done()
	defer wg.Done()

	for {
		time.Sleep(time.Second * 1)
		select {
		case <-cDone:
			fmt.Println("readLoop close")
			return
		default:
			handlerCh <- Message{netId, "Hello world"}
		}
	}
}

func handleLoop(c *ServerConn, wg *sync.WaitGroup) {
	handlerCh := c.handleCh
	sendCh := c.sendCh
	//ctx, _ := context.WithCancel(c.ctx)
	//cDone := ctx.Done()
	cDone := c.ctx.Done()
	defer wg.Done()

	for {
		select {
		case handleData, ok := <-handlerCh:
			if ok {
				handleData.netId++
				handleData.Data = "I am whole world"
				sendCh <- handleData
			}
		case <-cDone:
			fmt.Println("handleLoop close")
			return
		}
	}
}

func writeLoop(c *ServerConn, wg *sync.WaitGroup) {
	sendCh := c.sendCh
	//ctx, _ := context.WithCancel(c.ctx)
	//cDone := ctx.Done()
	cDone := c.ctx.Done()
	defer wg.Done()

	for {
		select {
		case sendData, ok := <-sendCh:
			if ok {
				fmt.Println(sendData)
			}
		case <-cDone:
			fmt.Println("writeLoop close")
			return
		}
	}
}