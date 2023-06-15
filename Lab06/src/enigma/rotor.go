package enigma

type Rotor struct {
	Alphabet      []rune
	Shift         int
	CountOfShifts int
}

func NewRotor(alphabet []rune, shift int) *Rotor {
	return &Rotor{Alphabet: alphabet, Shift: shift}
}

func (r *Rotor) ShiftRotor(turnsCount int) int {
	turns := 0
	if r.CountOfShifts < len(r.Alphabet) {
		r.CountOfShifts += r.Shift
	} else {
		r.CountOfShifts = len(r.Alphabet) - r.CountOfShifts
		turns++
	}

	for i := 0; i < r.Shift+turnsCount; i++ {
		temp := ' '
		for j := 0; j < len(r.Alphabet)-1; j++ {
			if j == 0 {
				temp = r.Alphabet[len(r.Alphabet)-1]
				r.Alphabet[len(r.Alphabet)-1] = r.Alphabet[j]
				r.Alphabet[j] = r.Alphabet[j+1]
			} else {
				r.Alphabet[j] = r.Alphabet[j+1]
			}

			if j == len(r.Alphabet)-2 {
				r.Alphabet[len(r.Alphabet)-2] = temp
			}
		}
	}

	return turns
}

func (r *Rotor) PickStartPosition(startPosition rune) {
	for i := 0; i < len(r.Alphabet); i++ {
		if r.Alphabet[i] == startPosition {
			r.ShiftRotor(i)
			break
		}
	}
}
