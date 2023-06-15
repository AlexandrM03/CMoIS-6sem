package trisemus

import "strings"

var trisemusTable = map[rune]rune{
	'A': 'E', 'B': 'N', 'C': 'I', 'D': 'G', 'E': 'M', 'F': 'A', 'G': 'B', 'H': 'C',
	'I': 'D', 'J': 'F', 'K': 'H', 'L': 'K', 'M': 'L', 'N': 'O', 'O': 'P', 'P': 'Q',
	'Q': 'R', 'R': 'S', 'S': 'T', 'T': 'U', 'U': 'V', 'V': 'W', 'W': 'X', 'X': 'Y',
	'Y': 'Z', 'Z': 'F', 'Ä': 'Ä', 'Ö': 'Ö', 'Ü': 'Ü', 'ß': 'ß',
}

func Encrypt(text string) string {
	var result strings.Builder
	for _, char := range text {
		encrypted, ok := trisemusTable[char]
		if ok {
			result.WriteRune(encrypted)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func Decrypt(text string) string {
	var result strings.Builder
	var found bool
	for _, char := range text {
		found = false
		for key, value := range trisemusTable {
			if char == value {
				result.WriteRune(key)
				found = true
				break
			}
		}
		if !found {
			result.WriteRune(char)
		}
	}
	return result.String()
}
