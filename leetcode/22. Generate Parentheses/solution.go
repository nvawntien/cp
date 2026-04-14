func generateParenthesis(n int) []string {
    ans := []string{}

    var backtrack func (string, int, int)
    backtrack = func (str string, open, close int) {
        if len(str) == 2 * n {
            ans = append(ans, str)
            return
        }

        if open < n {
            backtrack(str + "(", open+1, close)
        }
      
        if close < open {
            backtrack(str + ")", open, close+1)
        }
    }

    backtrack("", 0, 0)

    return ans
}