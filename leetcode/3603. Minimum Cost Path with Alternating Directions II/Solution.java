class Solution {
    public long minCost(int m, int n, int[][] waitCost) {
        long[][] dp = new long[m][n];

        dp[0][0] = 1;

        for (int i = 1; i < n; i++) {
            dp[0][i] = dp[0][i-1] + i+1 + waitCost[0][i];
        }

        for (int i = 1; i < m; i++) {
            dp[i][0] = dp[i-1][0] + i+1 + waitCost[i][0];
        }

        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                dp[i][j] = Math.min(dp[i-1][j], dp[i][j-1]) + (i+1) * (j+1) + waitCost[i][j];
            }
        }

        return dp[m-1][n-1] - waitCost[m-1][n-1];
    }
}