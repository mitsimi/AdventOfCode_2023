package util

import "strconv"

func Atoi(str string) int {
	num, err := strconv.Atoi(str)
	CheckErr(err)
	return num
}

func Itoa(i int) string {
	return strconv.Itoa(i)
}

func Atoi_Arr(stra []string) []int {
	arr := make([]int, len(stra))

	for i, v := range stra {
		arr[i] = Atoi(v)
	}

	return arr
}

func Itoa_Arr(iarr []int) []string {
	arr := make([]string, len(iarr))

	for i, v := range iarr {
		arr[i] = Itoa(v)
	}

	return arr
}
