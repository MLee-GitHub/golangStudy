package main

import (
	"flag"
	"fmt"
)

func main() {
	var which = flag.Int("w", 0, "specified functioin to test")
	flag.Parse()

	switch *which {
	case 0:
		strUtf8("你好，世界。golang")
	case 1:
		src := []int{10, 20, 30, 40}
		data := []int{11, 22, 33}
		n, newSrc, err := insertByIndex(2, data, src)
		if err != nil {
			fmt.Printf("insert fail: %v\n", err.Error())
		} else {
			fmt.Printf("insert succ(%d): %#v\n", n, newSrc)
		}
	default:
		fmt.Println("unsupported test.")
	}

}

func strUtf8(str string) {
	fmt.Printf("%#v\n", []byte(str)) // 每个字节都是一个ASCII码，汉字采用3个字节
	fmt.Println("\xe4\xbd\xa0")
	fmt.Println("\xe5\xa5\xbd")
}

// 在索引i处依次插入data切片中数据
func insertByIndex(index int, data []int, src []int) (int, []int, error) {
	if index < 0 || index > len(src) {
		return 0, nil, fmt.Errorf("index out of range.")
	}
	src = append(src, make([]int, len(data))...)
	copy(src[index+len(data):], src[index:])
	n := copy(src[index:], data)
	return n, src, nil
}
