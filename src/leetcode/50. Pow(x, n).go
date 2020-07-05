package main

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	return helper(x, n / 2)
}

func helper(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	res := helper(x, n)

	if n % 2 == 0{
		return res * res
 	} else {
 		return  res * res * x
	}
}

func main() {
	
}
