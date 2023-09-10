package piscine

func RecursiveFactorial(nb int) int {
	if nb < 25 {
		if nb == 0 {
			return 1
		}
		if nb < 0 {
			return 0
		}
		return nb * RecursiveFactorial(nb-1)
	} else {
		return 0
	}
}
