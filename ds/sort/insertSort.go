package sort

func InsertSort(data []int){
	for i := 1; i < len(data); i++{
		p := i - 1 // 当前轮次待插入元素前一个元素索引
		for ;p >= 0 && data[p] > data[p+1];{
			data[p], data[p+1] = data[p+1], data[p]
			p--
		}
	}
}
