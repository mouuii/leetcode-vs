func solve(a []int) int {
	var sum int
	lens := len(a)
	for _, value := range a {
		sum += value
	}
	return ((1+lens)*lens)/2 - sum
}
