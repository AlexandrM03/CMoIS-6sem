package enigma

type Enimga struct {
	RotorA    *Rotor
	RotorB    *Rotor
	RotorC    *Rotor
	Reflector *Reflector
}

func NewEnigma(rotorA *Rotor, rotorB *Rotor, rotorC *Rotor, reflector *Reflector) *Enimga {
	return &Enimga{RotorA: rotorA, RotorB: rotorB, RotorC: rotorC, Reflector: reflector}
}

func NewEnigmaGen() *Enimga {
	alphabetA := []rune{'f', 'k', 'q', 'h', 't', 'l', 'x', 'o', 'c', 'b', 'j', 's', 'p', 'd', 'z', 'r', 'a', 'm', 'e', 'w', 'n', 'i', 'u', 'y', 'g', 'v'}
	alphabetB := []rune{'a', 'j', 'd', 'k', 's', 'i', 'r', 'u', 'x', 'b', 'l', 'h', 'w', 't', 'm', 'c', 'q', 'g', 'z', 'n', 'p', 'y', 'f', 'v', 'o', 'e'}
	alphabetC := []rune{'e', 's', 'o', 'v', 'p', 'z', 'j', 'a', 'y', 'q', 'u', 'i', 'r', 'h', 'x', 'l', 'n', 'f', 't', 'g', 'k', 'd', 'c', 'm', 'w', 'b'}

	rotorL := NewRotor(alphabetA, 1)
	rotorM := NewRotor(alphabetB, 0)
	rotorR := NewRotor(alphabetC, 1)

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

	reflector := NewReflector(reflectorMap)

	return NewEnigma(rotorL, rotorM, rotorR, reflector)
}
