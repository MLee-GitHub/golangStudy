package sundry

import "fmt"
import "time"

func ChanNilArr() {
	c1, c2 := make(chan struct{}), make(chan struct{})
	go useZeroArray(c1)
	go useZeroStruct(c2)
	_, _ = <-c1, <-c2
	fmt.Println("main exit.")
}

// 长度为0的数组并不占用内存空间，只是通过管道来达到协程间的通信同步
func useZeroArray(c chan struct{}) {
	c1 := make(chan [0]int)
	go func() {
		fmt.Println("sub function in zeroArray.")
		c1 <- [0]int{}
		time.Sleep(time.Duration(2) * time.Second)
	}()
	<-c1
	fmt.Println("zeroArray exit.")
	c <- struct{}{}
}

// 使用无类型匿名结构体(推荐)
func useZeroStruct(c chan struct{}) {
	c1 := make(chan struct{})
	go func() {
		fmt.Println("sub function in zeroStruct.")
		c1 <- struct{}{} // struct{}匿名结构体(类型),{}表示结构体取值
	}()
	<-c1
	fmt.Println("zeroStruct exit.")
	c <- struct{}{}
}
