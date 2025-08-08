/*
    author: n.vantien
    from: University of Engineering and Technology
*/

#include <bits/stdc++.h>
using namespace std;

const int MOD = 1e9 + 7;

int main() {
    ios_base :: sync_with_stdio(0);
    cin.tie(0); cout.tie(0);

    int n, m;
    cin >> n >> m;

    vector <vector<char>> arr(n+1, vector<char>(m+1, 0));

    for (int i = 1; i <= n; ++i) {
        for (int j = 1; j <= m; ++j) {
            cin >> arr[i][j];
        }
    }

    vector <vector<int>> dp(n+1, vector <int> (m+1, 0));

    dp[1][1] = 1;

    for (int i = 1; i <= n; ++i) {
        for (int j = 1; j <= m; ++j) {
            if (i == 1 && j == 1) continue;
            if (arr[i][j] == '#') continue;
            dp[i][j] = (dp[i-1][j] % MOD + dp[i][j-1] % MOD) % MOD;
        }
    }

    cout << dp[n][m];

    return 0;
}
