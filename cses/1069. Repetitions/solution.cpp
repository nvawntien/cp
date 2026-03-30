#include <bits/stdc++.h>
using namespace std;

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    cout.tie(nullptr);

    string str;
    cin >> str;

    int A = 0, C = 0, G = 0, T = 0;
    int ans = 0;
    for (char c : str) {
        if (c == 'A') {
            A++;
            ans = max(ans, max({A, C, G, T}));
            C = G = T = 0;
        } else if (c == 'C') {
            C++;
            ans = max(ans, max({A, C, G, T}));
            A = G = T = 0;
        } else if (c == 'G') {
            G++;
            ans = max(ans, max({A, C, G, T}));
            A = C = T = 0;
        } else if (c == 'T') {
            T++;
            ans = max(ans, max({A, C, G, T}));
            A = C = G = 0;
        }
    }

    cout << ans << '\n';
    return 0;
}