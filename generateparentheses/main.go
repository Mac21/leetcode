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
var dp func(l, r int, s string)

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	dp = func(l, r int, s string) {
		if l == 0 && r == 0 {
			res = append(res, s)
			return
		}

		if l > 0 {
			dp(l-1, r, s+"(")
		}

		if r > l {
			dp(l, r-1, s+")")
		}
	}
	dp(n, n, "")
	return res
}
