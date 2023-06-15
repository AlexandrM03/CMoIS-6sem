package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(-1)
	b := big.NewInt(1)
	p := big.NewInt(751)

	for x := big.NewInt(516); x.Cmp(big.NewInt(550)) <= 0; x = new(big.Int).Add(x, big.NewInt(1)) {
		y := new(big.Int).Sub(new(big.Int).Exp(x, big.NewInt(3), p), new(big.Int).Mul(x, a))
		y.Add(y, b)
		y.Mod(y, p)

		if isOnCurve(x, y, a, b, p) {
			fmt.Printf("(%v, %v)\n", x, y)
		}
	}
}

func isOnCurve(x, y, a, b, p *big.Int) bool {
	left := new(big.Int).Exp(y, big.NewInt(2), p)
	right := new(big.Int).Exp(x, big.NewInt(3), p)
	right.Mul(right, x)
	right.Add(right, new(big.Int).Mul(x, a))
	right.Add(right, b)
	right.Mod(right, p)

	return left.Cmp(right) == 0
}
