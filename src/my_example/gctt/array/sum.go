package main

 
// Sum ... 
func Sum(numbers []int) (sum int){
	// part 01
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	// part 02
	for _, number := range numbers {
		sum += number
	}
	return
}
 
// SumAll ([]int{1,2}, []int{0,9}) would return []int{3, 9}
func SumAll(numbersToSum ...[]int) (sums []int) {
	// part 01
	// lengthOfNumbers := len(numbersToSum)
	// sums = make([]int, lengthOfNumbers)

	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }

	// part 02

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}
// SumAllTails ..
func SumAllTails(numbersToSum ...[]int) (sums []int){
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return 
}