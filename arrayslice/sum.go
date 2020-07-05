package arrayslice

// Sum : calculate sum of two integers
func Sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

// SumAll : 可变参数，计算各切片元素和
func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)

	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return
}

// SumAllTails : 计算多个切片不包括第一个元素的和
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)
	for index, numbers := range numbersToSum {
		// tail := numbers[1:]
		if len(numbers) == 0 {
			sums[index] = 0
		} else {
			sums[index] = Sum(numbers[1:])
		}
	}
	return
}
