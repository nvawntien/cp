#include <iostream>
#include <vector>
using namespace std;

const int MOD = 1e9 + 7;

int main() {
    ios_base :: sync_with_stdio(0);
    cin.tie(0);  cout.tie(0);

    int n, x;
    cin >> n >> x;

    vector <long long> coins(n), dp(x+1);

    for (auto &coin : coins) cin >> coin;
    dp[0] = 1;

    for (auto &coin : coins) {
        for (int i = 1; i <= x; ++i) {
            if (i - coin >= 0) {
                dp[i] = (dp[i] + dp[i-coin]) % MOD;
            }
        }
    }

    cout << dp[x] % MOD << '\n';
    return 0;
}