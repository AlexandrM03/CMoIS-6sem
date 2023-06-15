package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"unicode"
	"unicode/utf8"
)

func calculateEntropy(file *os.File) (float64, float64) {
	freq := make(map[rune]int)
	scanner := bufio.NewScanner(bufio.NewReader(file))
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		line := scanner.Text()
		for len(line) > 0 {
			r, size := utf8.DecodeRuneInString(line)
			line = line[size:]
			if unicode.IsLetter(r) {
				freq[unicode.ToLower(r)]++
			}
		}
	}

	fmt.Println("Frequency: ")
	for k, v := range freq {
		fmt.Print("(", string(k), ": ", v, "); ")
	}
	fmt.Println()

	err := createReport(freq, "Entropy")
	if err != nil {
		fmt.Println("Error writing report:", err)
	}
	total := findTotal(freq)

	fmt.Println("Probability: ")
	prob := make(map[rune]float64)
	for symbol, count := range freq {
		prob[symbol] = float64(count) / float64(total)
		fmt.Printf("(%c: %0.2f); ", symbol, prob[symbol])
	}
	fmt.Println()
	entropy := shennonFano(prob)

	return entropy, entropy * float64(total)
}

func calculateEntropyBinary(file *os.File) float64 {
	freq := make(map[byte]int)

	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		panic(err)
	}

	for i := 0; i < n; i++ {
		for j := 7; j >= 0; j-- {
			b := (buffer[i] >> uint(j)) & 1
			freq[byte(b)]++
		}
	}

	fmt.Println("Frequency: ")
	for k, v := range freq {
		fmt.Print("(", k, ": ", v, "); ")
	}
	fmt.Println()

	total := findTotal(freq)

	fmt.Println("Probability: ")
	prob := make(map[byte]float64)
	for bit, count := range freq {
		prob[bit] = float64(count) / float64(total)
		fmt.Printf("(%d: %0.2f); ", bit, prob[bit])
	}
	fmt.Println()

	return shennonFano(prob)
}

func calculateEntropyWithError(file *os.File, p float64) float64 {
	freq := make(map[rune]int)
	scanner := bufio.NewScanner(bufio.NewReader(file))
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		line := scanner.Text()
		for len(line) > 0 {
			r, size := utf8.DecodeRuneInString(line)
			line = line[size:]
			if unicode.IsLetter(r) {
				freq[unicode.ToLower(r)]++
			}
		}
	}

	fmt.Println("Frequency: ")
	for k, v := range freq {
		fmt.Print("(", string(k), ": ", v, "); ")
	}
	fmt.Println()

	total := findTotal(freq)

	fmt.Println("Probability: ")
	prob := make(map[rune]float64)
	for symbol, count := range freq {
		prob[symbol] = float64(count) / float64(total)
		fmt.Printf("(%c: %0.2f); ", symbol, prob[symbol])
	}
	fmt.Println()

	q := 1 - p
	if q == 0 || p == 0 {
		return 0
	}

	conditional := 1 - (-p*math.Log2(p) - q*math.Log2(q))
	entropy := 0.0
	for _, p := range prob {
		entropy += p*math.Log2(p) - conditional
	}

	return -entropy
}

func findTotal[T rune | byte](freq map[T]int) int {
	total := 0
	for _, count := range freq {
		total += count
	}
	return total
}

func shennonFano[T rune | byte](prob map[T]float64) float64 {
	entropy := 0.0
	for _, p := range prob {
		entropy -= p * math.Log2(p)
	}

	return entropy
}
