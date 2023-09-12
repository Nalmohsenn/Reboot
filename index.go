package piscine

func Index(s string, toFind string) int {
	index := -1
	index2 := 0
	count := 0
	indexposition := -1
	if len(s) > 0 && len(toFind) > 0 {
		if len(s) == len(toFind) {
			for i := 0; i <= len(s); i++ {
				if Compare(s, toFind) == 0 {
					index = i
					return index
				}
			}
		} else if len(s) > len(toFind) {
			for i := 0; i < len(s); i++ {
				if s[i] == toFind[0] && len(toFind) > 1 {
					indexposition = i
					for l := 1; l < len(toFind); l++ {
						if s[i+l] != toFind[l] {
							indexposition = -1
						}
					}
					index = indexposition
				} else if s[i] == toFind[0] && len(toFind) == 1 {
					index = i
					break
				}
			}
			if len(toFind) > 1 && index != -1 {
				for i := index; i < len(s); i++ {
					if s[i] == toFind[index2] {
						count++
						index2++
						if count == len(toFind) {
							return index
						}
					} else {
						break
					}
				}
				if count == len(toFind) {
					return index
				}
			} else {
				return index
			}
		}
	} else {
		return -1
	}
	return index
}
