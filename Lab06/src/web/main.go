package main

import (
	"enigma/src/enigma"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type options struct {
	RotorA string `json:"rotorA"`
	RotorB string `json:"rotorB"`
	RotorC string `json:"rotorC"`

	Message string `json:"message"`
}

func main() {
	originalAlphabet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Post("/", func(c *fiber.Ctx) error {
		var opts options
		if err := c.BodyParser(&opts); err != nil {
			return err
		}

		enigma := enigma.NewEnigmaGen()
		enigma.RotorA.PickStartPosition(rune(opts.RotorA[0]))
		enigma.RotorB.PickStartPosition(rune(opts.RotorB[0]))
		enigma.RotorC.PickStartPosition(rune(opts.RotorC[0]))

		messageRunes := []rune(opts.Message)
		encryptedMessage := make([]rune, len(messageRunes))

		var encryptedSymbol rune
		for i := 0; i < len(messageRunes); i++ {
			temp := messageRunes[i]
			encryptedSymbol = enigma.RotorA.Alphabet[SliceIndex(len(enigma.RotorA.Alphabet), enigma.RotorA.Alphabet, temp)]
			encryptedSymbol = enigma.RotorB.Alphabet[SliceIndex(len(enigma.RotorB.Alphabet), enigma.RotorB.Alphabet, encryptedSymbol)]
			encryptedSymbol = enigma.RotorC.Alphabet[SliceIndex(len(enigma.RotorC.Alphabet), enigma.RotorC.Alphabet, encryptedSymbol)]

			encryptedSymbol = enigma.Reflector.GetLetter(encryptedSymbol)

			encryptedSymbol = originalAlphabet[SliceIndex(len(originalAlphabet), enigma.RotorA.Alphabet, encryptedSymbol)]
			encryptedSymbol = originalAlphabet[SliceIndex(len(originalAlphabet), enigma.RotorB.Alphabet, encryptedSymbol)]
			encryptedSymbol = originalAlphabet[SliceIndex(len(originalAlphabet), enigma.RotorC.Alphabet, encryptedSymbol)]

			encryptedMessage[i] = encryptedSymbol

			enigma.RotorA.ShiftRotor(enigma.RotorB.ShiftRotor(enigma.RotorC.ShiftRotor(0)))
		}

		return c.JSON(fiber.Map{
			"message": string(encryptedMessage),
		})
	})

	app.Listen(":3000")
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
