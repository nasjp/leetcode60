package main

// 22. Generate Parentheses
// https://leetcode.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	return helperGenerateParenthesis("", 0, 0, n, make([]string, 0))
}

func helperGenerateParenthesis(str string, opens, closes, n int, result []string) []string {
	if len(str) == n*2 {
		return append(result, str)
	}

	if opens < n {
		result = append(helperGenerateParenthesis(str+"(", opens+1, closes, n, result))
	}
	if closes < opens {
		result = append(helperGenerateParenthesis(str+")", opens, closes+1, n, result))
	}
	return result
}
