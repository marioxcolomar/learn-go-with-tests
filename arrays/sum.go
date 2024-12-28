package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var res []int

	for _, numbers := range numbersToSum {
		res = append(res, Sum(numbers))
	}
	return res
}

func SumAllTails(numbersToSum ...[]int) []int {
	var res []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			res = append(res, 0)
		} else {
			tail := numbers[1:]
			res = append(res, Sum(tail))
		}
	}

	return res
}
