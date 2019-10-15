package sort

func SelectSort(data []int){
	for i := 0; i < len(data); i++{
		minIndex := i // 本轮最小索引位置，初始为本轮第一个元素索引
		for j := i + 1; j < len(data); j++{
			if data[j] < data[minIndex]{
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
}
