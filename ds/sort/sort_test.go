package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

// https://zhuanlan.zhihu.com/p/57088609
const(
	num = 10
	strategy = 0
)
var data, dataCpy []int

func init(){
	rand.Seed(time.Now().Unix())
}

func getData(num int, strategy int) []int{
	data := make([]int, num)
	switch strategy {
	case 0:
		for i := 0; i < num; i++{
			data[i] = rand.Intn(1000)
		}
	case 1:
		for i := 0; i < num; i++{
			_, _ = fmt.Scan(&data[i])
		}
	default:
		data = rand.Perm(num)
	}
	return data
}

func print(type_ string, before []int, after []int){
	fmt.Printf("type: %s\nbefore: %v\nafter: %v\n", type_, before, after)
}


func AssertEqual(t *testing.T, a, b interface{}){
	if a != b {
		t.Errorf("not equal: %v, %v", a, b)
	}else {
		fmt.Println("equal")
	}
}

func TestSelectSort(t *testing.T){
	data = getData(num, strategy)
	dataCpy = make([]int, len(data))
	copy(dataCpy, data)
	SelectSort(data)
	print("select sort", dataCpy, data)
	AssertEqual(t, true, sort.IntsAreSorted(data))
}

func TestInsertSort(t *testing.T) {
	data = getData(num, strategy)
	dataCpy = make([]int, len(data))
	copy(dataCpy, data)
	InsertSort(data)
	print("insert sort", dataCpy, data)
	AssertEqual(t, true, sort.IntsAreSorted(data))
}

func TestBubbleSort(t *testing.T) {
	data = getData(num, strategy)
	dataCpy := make([]int, len(data))
	copy(dataCpy, data)
	BubbleSort(data)
	print("bubble sort", dataCpy, data)
	AssertEqual(t, true, sort.IntsAreSorted(data))
}
