class Solution {
public:
    long long maximumTotalDamage(vector<int>& power) {
        int n = power.size();
        sort(power.begin(), power.end());
        vector <long long> dp(n, 0), dp1(n, 0);

        for (int i = 0; i < n; i++) {
            if (i > 0 && power[i] == power[i-1]) {
                dp[i] = dp[i-1] + power[i];
                dp1[i] = max(dp1[i-1], dp[i]);
                continue;
            }

            int low = 0, high = i-1, index = -1;
            while (low <= high) {
                int mid = low + high >> 1;
                if (power[mid] <= power[i] - 3) {
                    index = mid;
                    low = mid+1;
                }
                else {
                    high = mid-1;
                }
            }

            if (index != -1) dp[i] = dp1[index] + power[i];
            else dp[i] += power[i];

            dp1[i] = max(dp1[max(i-1,0)], dp[i]);
        }
        
        return dp1[n-1];
    }
};