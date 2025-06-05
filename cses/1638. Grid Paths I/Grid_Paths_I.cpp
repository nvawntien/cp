#include <iostream>
#include <vector>
using namespace std;

const int MOD = 1e9 + 7;

int main() {
    ios_base :: sync_with_stdio(0);
    cin.tie(0); cout.tie(0);

    int n;
    cin >> n;

    vector <vector <char>> grid(n+1, vector <char> (n+1));

    for (int i = 1; i <= n; ++i) {
        for (int j = 1; j <= n; ++j) {
            cin >> grid[i][j];
        }
    }

    vector <vector <long long>> dp(n+1, vector <long long> (n+1));

    for (int i = 0; i <= n; ++i) dp[0][i] = 0, dp[i][0] = 0;
    dp[0][1] = 1;

    for (int i = 1; i <= n; ++i) {
        for (int j = 1; j <= n; ++j) {
            if (grid[i][j] == '*') dp[i][j] = 0;
            else dp[i][j] = (dp[i-1][j] % MOD + dp[i][j-1] % MOD) % MOD;
        }
    }

    cout << dp[n][n] << '\n';
    return  0;
}