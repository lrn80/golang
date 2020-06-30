package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	l, max := 0, 0
	m := make(map[byte]int)
	for r, ch := range []byte(s) {
		lastR, ok := m[ch]
		if ok && lastR >= l {
			l = lastR + 1
		}
		if r-l+1 > max {
			max =  r - l + 1
		}
		m[ch] = r
	}

	return max;
}

func main() {
	re := lengthOfLongestSubstring("aaaaccc")
	fmt.Println(re);
}
