package piscine

func Index(s string, toFind string) int {
	length := len(s)
	sublength := len(toFind)

	for i := 0; i < length-sublength; i++ {
		if s[i:i+sublength] == toFind {
			return i
		}
	}
	return -1
}
