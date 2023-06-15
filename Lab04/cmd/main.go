package main

import (
	"crypto4/mdr"
	"crypto4/report"
	"crypto4/trisemus"
	"fmt"
	"os"
	"time"
)

func main() {
	showTrisemus()
	showMdr()

	fmt.Println("Done.")
}

func showTrisemus() {
	inputFilename := "input.txt"
	encryptedFilename := "encrypted_tris.txt"
	decryptedFilename := "decrypted_tris.txt"

	inputFile, err := os.Open(inputFilename)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	defer inputFile.Close()

	inputStat, err := inputFile.Stat()
	if err != nil {
		fmt.Println("Error reading input file stats:", err)
		return
	}

	inputBytes := make([]byte, inputStat.Size())
	_, err = inputFile.Read(inputBytes)
	if err != nil {
		fmt.Println("Error reading input file data:", err)
		return
	}

	start := time.Now()
	encrypted := trisemus.Encrypt(string(inputBytes))
	elapsed := time.Since(start)
	err = report.BuildCharHistogram(encrypted, "Trisemus", "trisemus_encr.xlsx")
	if err != nil {
		fmt.Println("Error building char histogram:", err)
		return
	}
	fmt.Println("Encryption trisemus took", elapsed)

	encryptedFile, err := os.Create(encryptedFilename)
	if err != nil {
		fmt.Println("Error creating encrypted file:", err)
		return
	}
	defer encryptedFile.Close()

	_, err = encryptedFile.WriteString(encrypted)
	if err != nil {
		fmt.Println("Error writing encrypted file:", err)
		return
	}

	encryptedFile, err = os.Open(encryptedFilename)
	if err != nil {
		fmt.Println("Error reading encrypted file:", err)
		return
	}
	defer encryptedFile.Close()

	encryptedStat, err := encryptedFile.Stat()
	if err != nil {
		fmt.Println("Error reading encrypted file stats:", err)
		return
	}

	encryptedBytes := make([]byte, encryptedStat.Size())
	_, err = encryptedFile.Read(encryptedBytes)
	if err != nil {
		fmt.Println("Error reading encrypted file data:", err)
		return
	}

	start = time.Now()
	decrypted := trisemus.Decrypt(string(encryptedBytes))
	elapsed = time.Since(start)
	report.BuildCharHistogram(decrypted, "Trisemus", "trisemus_decr.xlsx")
	fmt.Println("Decryption trisemus took", elapsed)

	decryptedFile, err := os.Create(decryptedFilename)
	if err != nil {
		fmt.Println("Error creating decrypted file:", err)
		return
	}
	defer decryptedFile.Close()

	_, err = decryptedFile.WriteString(decrypted)
	if err != nil {
		fmt.Println("Error writing decrypted file:", err)
		return
	}
}

func showMdr() {
	inputFilename := "input.txt"
	encryptedFilename := "encrypted_mdr.txt"
	decryptedFilename := "decrypted_mdr.txt"

	inputFile, err := os.Open(inputFilename)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	defer inputFile.Close()

	inputStat, err := inputFile.Stat()
	if err != nil {
		fmt.Println("Error reading input file stats:", err)
		return
	}

	inputBytes := make([]byte, inputStat.Size())
	_, err = inputFile.Read(inputBytes)
	if err != nil {
		fmt.Println("Error reading input file data:", err)
		return
	}

	start := time.Now()
	encrypted := mdr.Encrypt(string(inputBytes), 7)
	elapsed := time.Since(start)
	report.BuildCharHistogram(encrypted, "Mdr", "mdr_encr.xlsx")
	fmt.Println("Encryption mdr took", elapsed)

	encryptedFile, err := os.Create(encryptedFilename)
	if err != nil {
		fmt.Println("Error creating encrypted file:", err)
		return
	}
	defer encryptedFile.Close()

	_, err = encryptedFile.WriteString(encrypted)
	if err != nil {
		fmt.Println("Error writing encrypted file:", err)
		return
	}

	encryptedFile, err = os.Open(encryptedFilename)
	if err != nil {
		fmt.Println("Error reading encrypted file:", err)
		return
	}
	defer encryptedFile.Close()

	encryptedStat, err := encryptedFile.Stat()
	if err != nil {
		fmt.Println("Error reading encrypted file stats:", err)
		return
	}

	encryptedBytes := make([]byte, encryptedStat.Size())
	_, err = encryptedFile.Read(encryptedBytes)
	if err != nil {
		fmt.Println("Error reading encrypted file data:", err)
		return
	}

	start = time.Now()
	decrypted := mdr.Decrypt(string(encryptedBytes), 7)
	elapsed = time.Since(start)
	report.BuildCharHistogram(decrypted, "Mdr", "mdr_decr.xlsx")
	fmt.Println("Decryption mdr took", elapsed)

	decryptedFile, err := os.Create(decryptedFilename)
	if err != nil {
		fmt.Println("Error creating decrypted file:", err)
		return
	}
	defer decryptedFile.Close()

	_, err = decryptedFile.WriteString(decrypted)
	if err != nil {
		fmt.Println("Error writing decrypted file:", err)
		return
	}
}
