package main

func myPow2(x float64, n int) float64 {
	if x == 0 || n == 1{
		return x
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	pow := 1.0

	for n != 0  {
		if n & 1 == 1 {
			pow = pow * x
		}

		x = x * x

		n = n >> 1
	}

	return pow
}
func main() {
	
}
