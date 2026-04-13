class Solution {
public:
    int maxSum(vector<int>& nums) {
        set <int> s;

        int maxVal = INT_MIN;

        for (int &a : nums) {
            if (a <= 0) {
                maxVal = max(a, maxVal);
                continue;
            }
            s.insert(a);
        }

        if (!s.size()) return maxVal;
        int sum = 0;

        for (int a : s) sum += a;
        return max(sum, maxVal); 
    }
};