package mylib

// Average 平均値を返す
func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}

	return int(total / len(s))
}
