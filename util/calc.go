package util

type Number interface {
	int | int8 | int16 | int32 | int64 |
	uint | uint8 | uint16 | uint32 | uint64 |
	float32 | float64
}

func Sum[N Number](nums []N) N {
	var sum N
	for _, n := range nums {
		sum += n
	}
	return sum
}