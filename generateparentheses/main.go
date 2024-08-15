package main

/*
22. Generate Parentheses Medium
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]

Input: n = 2
Output: ["()()", "(())"]

Constraints:

1 <= n <= 8
*/
var dp func(i int) []string

func generateParenthesis(n int) []string {
	cache := make([]string, 2)
	cache[0] = ""
	cache[1] = "()"
	dp = func(i int) []string {
		if i > 0 && i < len(cache) {
			return cache[i : i+1]
		}

		pp := dp(i - 1)
		currentLevelCacheIndex := len(cache)
		for _, p := range pp {
			nest := "(" + p + ")"
			cache = append(cache, nest)
			concatf := "()" + p
			if nest != concatf {
				cache = append(cache, concatf)
			}
			concatb := p + "()"
			if concatf != concatb {
				cache = append(cache, concatb)
			}
		}
		return cache[currentLevelCacheIndex:]
	}
	return dp(n)
}
