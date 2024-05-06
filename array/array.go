package array

func Sum(numbers []int) int {
	sum := 0
	for _, t := range numbers {
		sum += t
	}
	return sum
}

func SumAll(args ...[]int) []int {
	var sums []int
	for _, numbers := range args {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumTails(args ...[]int) []int {
	var sums []int
	for _, numbers := range args {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
