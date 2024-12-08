package main

import "slices"

func FindSubstring(s string, words []string) []int {
	res := make([]int, 0)
	window := len(words) * len(words[0])
	for i := 0; i <= len(s)-window; i++ {
		subStr := s[i : window+i]
		if InWords(subStr, words) {
			res = append(res, i)
		}
	}
	return res
}

func InWords(subStr string, words []string) bool {
	clone := slices.Clone(words)
	window := len(words[0])

	for i := 0; i < len(words); i++ {
		index := slices.Index(clone, subStr[i*window:i*window+window])
		if index != -1 {
			clone = slices.Delete(clone, index, index+1)
			if len(clone) == 0 {
				return true
			}
		}
	}
	return false
}
