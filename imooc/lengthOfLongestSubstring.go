package main

import "fmt"

// "abcabcbb"

func lengthOfLongestSubstring(s string) int {
	lastOccured := make(map[byte]int)
	start, maxLength := 0, 0
	for i, ch := range []byte(s) {

		lastI, ok := lastOccured[ch]

		if ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccured[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(
		lengthOfLongestSubstring("abcabcbb"),
	)
}
