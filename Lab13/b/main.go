package main

import (
	"crypto/elliptic"
	"fmt"
	"math/big"

	"gitlab.com/elktree/ecc"
)

type Point struct {
	x, y *big.Int
}

func NewPoint(x, y int64) *Point {
	return &Point{
		x: big.NewInt(x),
		y: big.NewInt(y),
	}
}

func (p *Point) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func (p *Point) Add(q *Point, a, b, prime *big.Int) *Point {
	if p.IsInfinity() {
		return q
	}
	if q.IsInfinity() {
		return p
	}

	if p.Equals(q) || p.Equals(q.Negate()) {
		return p.Double(a, b, prime)
	}

	lambda := new(big.Int)
	lambda.Sub(q.y, p.y)
	lambda.Mod(lambda, prime)

	if lambda.Cmp(big.NewInt(0)) == 0 {
		return NewPoint(0, 0)
	}

	lambda.Mul(lambda.ModInverse(lambda, prime), new(big.Int).Sub(q.x, p.x))
	lambda.Mod(lambda, prime)

	x3 := new(big.Int)
	x3.Exp(lambda, big.NewInt(2), prime)
	x3.Sub(x3, p.x)
	x3.Sub(x3, q.x)
	x3.Mod(x3, prime)

	y3 := new(big.Int)
	y3.Sub(p.x, x3)
	y3.Mul(lambda, y3)
	y3.Sub(y3, p.y)
	y3.Mod(y3, prime)

	return &Point{x3, y3}
}

func (p *Point) Subtract(q *Point, a, b, prime *big.Int) *Point {
	qNeg := q.Negate()
	return p.Add(qNeg, a, b, prime)
}

func (p *Point) Multiply(k *big.Int, a, b, prime *big.Int) *Point {
	result := NewPoint(0, 0)
	accumulator := NewPoint(p.x.Int64(), p.y.Int64())

	for i := k.BitLen() - 1; i >= 0; i-- {
		result = result.Add(result, a, b, prime)
		if k.Bit(i) == 1 {
			result = result.Add(accumulator, a, b, prime)
		}
		accumulator = accumulator.Add(accumulator, a, b, prime)
	}

	return result
}

func (p *Point) Double(a, b, prime *big.Int) *Point {
	if p.IsInfinity() {
		return p
	}

	if p.y.Cmp(big.NewInt(0)) == 0 {
		return NewPoint(0, 0)
	}

	lambda := new(big.Int)
	lambda.Exp(p.x, big.NewInt(2), prime)
	lambda.Mul(lambda, big.NewInt(3))
	lambda.Add(lambda, a)
	lambda.Mod(lambda, prime)

	if lambda.Cmp(big.NewInt(0)) == 0 {
		return NewPoint(0, 0)
	}

	lambda.Mul(lambda.ModInverse(lambda, prime), big.NewInt(2).Mul(p.y, prime))
	lambda.Mod(lambda, prime)

	x3 := new(big.Int)
	x3.Exp(lambda, big.NewInt(2), prime)
	x3.Sub(x3, big.NewInt(2).Mul(p.x, prime))
	x3.Mod(x3, prime)

	y3 := new(big.Int)
	y3.Sub(p.x, x3)
	y3.Mul(lambda, y3)
	y3.Sub(y3, p.y)
	y3.Mod(y3, prime)

	return &Point{x3, y3}
}

func (p *Point) Negate() *Point {
	if p.IsInfinity() {
		return p
	}
	return &Point{p.x, new(big.Int).Neg(p.y)}
}

func (p *Point) Equals(q *Point) bool {
	if p.IsInfinity() && q.IsInfinity() {
		return true
	}
	return p.x.Cmp(q.x) == 0 && p.y.Cmp(q.y) == 0
}

func (p *Point) IsInfinity() bool {
	return p.x.Cmp(big.NewInt(0)) == 0 && p.y.Cmp(big.NewInt(0)) == 0
}

func main() {
	a := big.NewInt(-1)
	b := big.NewInt(1)
	prime := big.NewInt(751)
	pointP := NewPoint(3, 6)
	pointQ := NewPoint(5, 1)
	k := big.NewInt(8)
	l := big.NewInt(5)

	resultA := pointP.Multiply(k, a, b, prime)
	resultB := pointP.Add(pointQ, a, b, prime)
	resultC := pointP.Multiply(k, a, b, prime).Add(pointQ.Multiply(l, a, b, prime), a, b, prime).Subtract(pointQ, a, b, prime)
	resultD := pointP.Subtract(pointQ, a, b, prime).Add(pointP, a, b, prime)

	fmt.Println("kP:", resultA)
	fmt.Println("P + Q:", resultB)
	fmt.Println("kP + lQ - R:", resultC)
	fmt.Println("P - Q + R:", resultD)

	// Encrypt and decrypt message
	fmt.Println()
	pub, priv, _ := ecc.GenerateKeys(elliptic.P521())

	plaintext := "Mozolevskiy"

	fmt.Println("Plaintext:", plaintext)
	encrypted, _ := pub.Encrypt([]byte(plaintext))
	fmt.Println("Encrypted:", encrypted)
	decrypted, _ := priv.Decrypt(encrypted)
	fmt.Println("Decrypted:", string(decrypted))

	// Sign and verify message
	fmt.Println()
	pub, priv, _ = ecc.GenerateKeys(elliptic.P384())

	plaintext = "secret secrets are no fun, secret secrets hurt someone"
	sig, _ := priv.SignMessage([]byte(plaintext))
	fmt.Println("Signature:", sig)

	verified, _ := pub.VerifyMessage([]byte(plaintext), sig)
	fmt.Println(verified)
}
