class Solution {
public:
    string largestGoodInteger(string num) {
        char cur = 'a', ans = ' ';

        int count = 0;

        for (char c : num) {
            if (c == cur) count++;
            else count = 1;

            if (count == 3) {
                ans = max(ans, c);
            }

            cur = c;
        } 

        return  (ans == ' ') ? "":string(3, ans);
    }
};