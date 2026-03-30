#include <iostream>
#include <vector>
using namespace std;

const int MOD = 1e9 + 7;

int main() {
    ios_base :: sync_with_stdio(0);
    cin.tie(0); cout.tie(0);

    int n;
    cin >> n;
    vector <long long> dp(n+1, 0);

    // dp[0] = 1, dp[1] = 1, dp[2] = 2, dp[n] = dp[n-1] + dp[n-2] .... dp[n-6]
    dp[0] = 1, dp[1] = 1, dp[2] = 2, dp[3] = 4, dp[4] = 8, dp[5] = 16, dp[6] = 32;
    for (int i = 7; i <= n; ++i) {
        dp[i] = (dp[i-1]%MOD + dp[i-2]%MOD + dp[i-3]%MOD + dp[i-4]%MOD + dp[i-5]%MOD + dp[i-6]%MOD);
    }

    cout << dp[n] % MOD;
    return 0;
}