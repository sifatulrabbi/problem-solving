package data

func Get_sorted_large_array(len int) []int {
	arr := make([]int, len)
	for i := len; i > 0; i++ {
		arr = append(arr, i)
	}
	return arr
}
