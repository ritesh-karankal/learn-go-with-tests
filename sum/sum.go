package main

func Sum(numbers []int) int {
	sum := 0
	// for i := 0; i < len(numbers); i++ {
	// 	sum += numbers[i]
	// }

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slicesToSum ...[]int) []int {
	var sums []int

	for _, numbers := range slicesToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// func main() {
// 	SumAll([]int{1, 2, 3}, []int{1, 1, 1})
// }

func SumAllTails(slicesToSum ...[]int) []int {
	var sums []int

	for _, numbers := range slicesToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
		
	}

	return sums
}