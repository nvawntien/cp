#include <algorithm>
#include <iostream>
#include <vector>
#include <climits>
using namespace std;

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);

    int n;
    cin >> n;
    vector <int> arr(n);
    long long sum = 0, ans = LLONG_MAX;
    for (auto &a : arr) cin >> a, sum += a;

    for (int mask = 0; mask < (1 << n); mask++) {
        long long curr = 0;
        for (int i = 0; i < n; i++) {
            if (mask >> i & 1) {
                curr += arr[i];
            }
        }

        ans = min(ans, abs(sum - 2 * curr));
    }

    cout << ans << '\n';
    return 0;
}