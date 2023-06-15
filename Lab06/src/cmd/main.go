package main

import (
	"bufio"
	"enigma/src/enigma"
	"fmt"
	"os"
)

func main() {
	originalAlphabet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	alphabetA := []rune{'f', 'k', 'q', 'h', 't', 'l', 'x', 'o', 'c', 'b', 'j', 's', 'p', 'd', 'z', 'r', 'a', 'm', 'e', 'w', 'n', 'i', 'u', 'y', 'g', 'v'}
	alphabetB := []rune{'a', 'j', 'd', 'k', 's', 'i', 'r', 'u', 'x', 'b', 'l', 'h', 'w', 't', 'm', 'c', 'q', 'g', 'z', 'n', 'p', 'y', 'f', 'v', 'o', 'e'}
	alphabetC := []rune{'e', 's', 'o', 'v', 'p', 'z', 'j', 'a', 'y', 'q', 'u', 'i', 'r', 'h', 'x', 'l', 'n', 'f', 't', 'g', 'k', 'd', 'c', 'm', 'w', 'b'}

	rotorL := enigma.NewRotor(alphabetA, 1)
	rotorM := enigma.NewRotor(alphabetB, 0)
	rotorR := enigma.NewRotor(alphabetC, 1)

	reflectorMap := map[rune]rune{
		'a': 'y',
		'b': 'r',
		'c': 'u',
		'd': 'h',
		'e': 'q',
		'f': 's',
		'g': 'l',
		'i': 'p',
		'j': 'x',
		'k': 'n',
		'm': 'o',
		't': 'z',
		'v': 'w',
	}
	reflector := enigma.NewReflector(reflectorMap)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose starting position for rotor L (a-z): ")
	startPositionL, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	reader.ReadString('\n')
	rotorL.PickStartPosition(startPositionL)

	fmt.Print("Choose starting position for rotor M (a-z): ")
	startPositionM, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	reader.ReadString('\n')
	rotorM.PickStartPosition(startPositionM)

	fmt.Print("Choose starting position for rotor R (a-z): ")
	startPositionR, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	reader.ReadString('\n')
	rotorR.PickStartPosition(startPositionR)

	fmt.Print("Choose message to encrypt: ")
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	messageRunes := []rune(message)
	messageRunes = messageRunes[:len(messageRunes)-2]
	encryptedMessage := make([]rune, len(messageRunes))

	var encryptedSymbol rune
	for i := 0; i < len(messageRunes); i++ {
		temp := messageRunes[i]
		encryptedSymbol = rotorR.Alphabet[SliceIndex(len(rotorR.Alphabet), rotorR.Alphabet, temp)]
		encryptedSymbol = rotorM.Alphabet[SliceIndex(len(rotorM.Alphabet), rotorM.Alphabet, encryptedSymbol)]
		encryptedSymbol = rotorL.Alphabet[SliceIndex(len(rotorL.Alphabet), rotorL.Alphabet, encryptedSymbol)]

		encryptedSymbol = reflector.GetLetter(encryptedSymbol)

		encryptedSymbol = originalAlphabet[SliceIndex(len(originalAlphabet), rotorL.Alphabet, encryptedSymbol)]
		encryptedSymbol = originalAlphabet[SliceIndex(len(originalAlphabet), rotorM.Alphabet, encryptedSymbol)]
		encryptedSymbol = originalAlphabet[SliceIndex(len(originalAlphabet), rotorR.Alphabet, encryptedSymbol)]

		encryptedMessage[i] = encryptedSymbol

		rotorL.ShiftRotor(rotorM.ShiftRotor(rotorR.ShiftRotor(0)))
	}

	fmt.Printf("\nEncrypted message: %s", string(encryptedMessage))
}

func SliceIndex(limit int, alphabet []rune, symbol rune) int {
	var xIndex int
	for i := 0; i < limit; i++ {
		if alphabet[i] == symbol {
			return i
		}
		if alphabet[i] == 'x' {
			xIndex = i
		}
	}

	return xIndex
}
