package slices

func Sum(collection []int) int {
	var sum int
	for _, v := range collection {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int

	for _, n := range numbers {

		if len(n) > 0 {
			tail := n[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}

	}

	return sums
}
