class Solution {
public:
    bool rotateString(string s, string goal) {
        int n = s.size(), m = goal.size();
        if (n != m) return false;

        string text = s + s; 
        vector<int> lps(n, 0);

        auto build = [&]() {
            for (int i = 1; i < n; i++) {
                int j = lps[i-1];
                while (j > 0 && goal[i] != goal[j]) {
                    j = lps[j-1];
                }
                if (goal[i] == goal[j]) ++j;
                lps[i] = j;
            }
        };

        build(); 
        int j = 0;
        int len = text.size();

        for (int i = 0; i < len; ++i) {
            while (j > 0 && text[i] != goal[j]) {
                j = lps[j-1];
            }
            if (text[i] == goal[j]) ++j;
            
            if (j == n) return true; 
        }

        return false;
    }
};