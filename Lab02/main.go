package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	entropy, information := calculateEntropy(file)
	fmt.Printf("Entropy: %.2f bits/character\n", entropy)
	fmt.Printf("Overall information: %.2f\n\n", information)

	binaryFile, err := os.Open("file.bin")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer binaryFile.Close()

	binaryEntropy := calculateEntropyBinary(binaryFile)
	fmt.Printf("Binary entropy: %.2f bits/bit\n\n", binaryEntropy)

	errorFile, err := os.Open("fio.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	errorEntropy := calculateEntropyWithError(errorFile, 0.1)
	fmt.Printf("Error entropy (0.1): %.2f bits/character\n\n", errorEntropy)
	errorFile.Close()

	errorFile, err = os.Open("fio.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	errorEntropy = calculateEntropyWithError(errorFile, 0.5)
	fmt.Printf("Error entropy (0.5): %.2f bits/character\n\n", errorEntropy)
	errorFile.Close()

	errorFile, err = os.Open("fio.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	errorEntropy = calculateEntropyWithError(errorFile, 1)
	fmt.Printf("Error entropy (1): %.2f bits/character\n\n", errorEntropy)
	errorFile.Close()
}
