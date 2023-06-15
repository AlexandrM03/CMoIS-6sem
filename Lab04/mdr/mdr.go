package mdr

import "strings"

var alphabet = [...]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'Ä', 'Ö', 'Ü', 'ß'}

func Encrypt(text string, key int) string {
	var result strings.Builder
	for _, char := range text {
		index := indexOf(char, alphabet)
		if index >= 0 {
			encryptedIndex := (index + key) % len(alphabet)
			result.WriteRune(alphabet[encryptedIndex])
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func Decrypt(text string, key int) string {
	var result strings.Builder
	for _, char := range text {
		index := indexOf(char, alphabet)
		if index >= 0 {
			decryptedIndex := (index - key + len(alphabet)) % len(alphabet)
			result.WriteRune(alphabet[decryptedIndex])
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func indexOf(char rune, arr [30]rune) int {
	for i, c := range arr {
		if c == char {
			return i
		}
	}
	return -1
}
