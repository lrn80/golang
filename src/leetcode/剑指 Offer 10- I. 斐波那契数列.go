package main

func fib(n int) int {
	if n <= 1 {
		return n
	}
	pre := 0
	next := 1
	fn := 0
	for i := 2; i <= n; i++ {
		fn = (pre + next) % 1000000007
		pre = next
		next = fn
	}

	return fn
}

func main() {
	
}
