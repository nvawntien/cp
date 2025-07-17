class Solution {
    public int maximumLength(int[] nums, int k) {
        int n = nums.length;

        int[][] dp = new int[n][k];

        int ans = 0;

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < i; j++) {
                int mod = (nums[i] + nums[j]) % k;
                dp[i][mod] = dp[j][mod] + 1;
                ans = Math.max(ans, dp[i][mod] + 1);
            }
        }

        return ans;
    }
}