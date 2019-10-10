package sundry

import "fmt"

func Mem(){
	aSlice := make([]int, 10)
	aSlice[0] = 10
	aSlice[3] = 100
	bSlice := aSlice // aSlice 和 bSlice同一地址
	fmt.Println(aSlice, bSlice)
	bSlice[0] = 1
	aSlice[3] = 111
	fmt.Println(aSlice, bSlice)
	cSlice := make([]int, len(aSlice)) // 独立地址单元
	copy(cSlice, aSlice)
	cSlice[0] = 1010
	aSlice[3] = 108
	fmt.Println(aSlice, bSlice, cSlice)
}
