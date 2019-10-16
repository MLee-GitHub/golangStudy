package sort

func SheelSort(data []int){
	for gap := len(data)/2; gap > 0 ; gap /= 2{
		for i := gap; i < len(data); i++{
			for j := i - gap; j >=0 && data[j] > data[j+gap]; j -= gap {
				data[j], data[j+gap] = data[j+gap], data[j]
			}
		}
	}
}