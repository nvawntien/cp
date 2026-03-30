class Solution {
public:
    int maximumGain(string s, int x, int y) {
        deque <int> dq;
        string first = "ab", second = "ba";
        int n = s.length();

        vector <bool> check(n, false);

        if (x < y) {
            swap(x, y);
            swap(first, second);
        }

        int ans = 0;

        for (int i = 0; i < n; i++) {
            if (dq.size() && char(s[dq.back()]) == first[0] && s[i] == first[1]) {
                ans += x;
                check[i] = true;
                check[dq.back()] = true;
                dq.pop_back();
                continue;
            }
            dq.push_back(i);
        }

        while (dq.size()) dq.pop_back();

         for (int i = 0; i < n; i++) {
            if (check[i]) continue;
            if (dq.size() && char(s[dq.back()]) == second[0] && s[i] == second[1]) {
                ans += y;
                dq.pop_back();
                continue;
            }
            dq.push_back(i);
        }
        //for (int i = 0; i < n; i++) cout << check[i] << ' ';
        return ans;
    }
};