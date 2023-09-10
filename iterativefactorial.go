package piscine

func IterativeFactorial(nb int) int {
	result1 := 1
	for i := 0; i <= nb; i++ {
		result1 *= i
	}
	return result1
}
