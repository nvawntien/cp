class Solution {
public:
    bool check(int n) {
        while (n) {
            if (!(n % 10)) return false;
            n /= 10;
        }
        return true;
    }
    vector<int> getNoZeroIntegers(int n) {
        vector <int> ans;
        int a, b;
        for (int i = 1; i < n; ++i) {
            if (check(i) && check(n-i)) {
                 a = i;
                 b = n-i;
            }
        }

        ans.push_back(a);
        ans.push_back(b);
        return ans;
    }
};