#include <iostream>
#include <vector>
#include <climits>
using namespace std;

int main() {
    ios_base :: sync_with_stdio(0);
    cin.tie(0); 

    int n, x;
    cin >> n >> x;

    vector <int> coins(n);
    for (auto &coin : coins) cin >> coin;

    vector <long long> dp(x+1, INT_MAX);
    dp[0] = 0;

    for (int i = 1; i <= x; ++i) {
        for (auto &coin : coins) {
            if (i >= coin) {
                dp[i] = min(dp[i], dp[i-coin] + 1);
            }
        }
    }

    if (dp[x] == INT_MAX) cout << -1;
    else cout << dp[x];
    return 0;
}