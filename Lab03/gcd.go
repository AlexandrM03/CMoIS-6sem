package main

func gcd2(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd2(b, a%b)
}

func gcd3(a, b, c int) int {
	return gcd2(gcd2(a, b), c)
}
