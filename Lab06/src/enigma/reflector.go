package enigma

type Reflector struct {
	Alphabet map[rune]rune
}

func NewReflector(alphabet map[rune]rune) *Reflector {
	return &Reflector{Alphabet: alphabet}
}

func (r *Reflector) GetLetter(letter rune) rune {
	for key, value := range r.Alphabet {
		if key == letter {
			return value
		} else if value == letter {
			return key
		}
	}

	return 'x'
}
