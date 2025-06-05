#include <iostream>
#include <vector>
using namespace std;

int main() {
    ios_base :: sync_with_stdio(0);
    cin.tie(0); cout.tie(0);

    int n;
    cin >> n;

    vector <vector <int>> dp(n+1, vector <int> (10, 0));

    for (int i = 1; i <= n; ++i) {
        int k = i;
        while (k) {
            dp[i][k%10]++;
            k /= 10;
        }
    }

    int count = 0;

    while (n) {
        for (int i = 9; i >= 1; --i) {
            if (dp[n][i]) {
                n -= i;
                count++;
                break;
            } 
        }
    }

    cout << count << '\n';
    return 0;
}