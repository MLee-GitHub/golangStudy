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

//type redisJob struct {
//	uuid string
//}

//func NewVideoPacketServer () *VideoPacketServer {
//	v := &VideoPacketServer{
//		redisChan: make(chan *redisJob, 10000) // 异步任务
//		msgChan: make(chan *Msg, 10000)
//	}
//	go v.ProcessWriteToRedis()
//	go v.ProcessWriteToBoss()
//	return v
//}
//func (v *VideoPacketServer)ProcessWriteToRedis() {
//	for {
//		select {
//		case j := <-v.redisJob:
//			xxxxxx // 消费j
//		}
//	}
//}

type Msg struct {
	handleMsg chan []byte
	sendMsg chan []byte
	ctx context.Context
	cancel context.CancelFunc
}

func (v *VideoPacketServer)invoke(ctx context.Context, req []byte) []byte{
	v.redisChan <- &redisJob{uuid: uuid}

}