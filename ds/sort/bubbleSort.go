package sort

func BubbleSort(data []int){
	for i := 0; i < len(data); i++{
		flag := true // 当前轮次是否有交换标记
		for j := 0; j < len(data) - i - 1; j++{
			if data[j] > data[j+1]{
				data[j], data[j+1] = data[j+1], data[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}
