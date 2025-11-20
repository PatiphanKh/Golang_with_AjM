package cal

func Add(x ...int) (int, float32 ) {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	avg := float32(sum)/float32(len(x))
	return  sum, avg
}