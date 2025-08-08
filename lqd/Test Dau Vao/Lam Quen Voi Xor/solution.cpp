/*
    author: n.vantien
    from: University of Engineering and Technology
*/

#include <bits/stdc++.h>
using namespace std;

int main() {
    ios_base :: sync_with_stdio(0);
    cin.tie(0); cout.tie(0);

    int n;
    cin >> n;

    vector <long long> arr(n + 1);

    for (int i = 1; i <= n; i++) cin >> arr[i];

    vector <long long> Xor(n + 1, 0);

    for (int i = 1; i <= n; i++) {
        Xor[i] = Xor[i - 1] ^ arr[i];
    }

    long long ans = 0;

    for (int i = 1; i <= n; i++) {
        for (int len = 1; len <= n - i + 1; len++) {
            int l = i, r = i + len - 1;
            ans = max(ans, Xor[r] ^ Xor[l - 1]);
        }
    }

    cout << ans;
    return 0;
}
