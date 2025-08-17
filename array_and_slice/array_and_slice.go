package array

func GetLength3IntArray() [3]int {
	return [3]int{}
}

func GetIntSlice() []int {
	return []int{}
}

func AddElementToSlice(input int, slice []int) []int {
	return append(slice, input)
}

func MakeSlice(length int, capacity int) []int {
	return make([]int, length, capacity)
}