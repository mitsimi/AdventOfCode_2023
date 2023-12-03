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
