class Solution {
public:
    vector<int> maxValue(vector<int>& nums) {
        int n = nums.size();
        vector<int> ans(n), pre(n);   
        pre[0] = nums[0];
        int suf = INT_MAX;

        for (int i = 1; i < n; ++i) pre[i] = max(pre[i-1], nums[i]);

        for (int i = n-1; i >= 0; --i) {
            if (pre[i] > suf) ans[i] = ans[i+1];
            else ans[i] = pre[i];
            suf = min(suf, nums[i]);
        }

        return ans;        
    }
};