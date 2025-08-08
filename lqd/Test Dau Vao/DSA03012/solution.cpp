/*
    author: n.vantien
    from: University of Engineering and Technology
*/

#include <bits/stdc++.h>
using namespace std;

void solve() {
    string str;
    cin >> str;

    int n = str.length();
    map <char, int> freq;

    for (int i = 0; i < n; ++i) freq[str[i]]++;

    for (auto &it : freq) {
        if (it.second > n / 2) {
            cout << -1 << '\n';
            return;
        }
    }

    cout << 1 << '\n';
}

int main() {
    ios_base :: sync_with_stdio(0);
    cin.tie(0); cout.tie(0);


    int t; cin >> t;
    while (t--) solve();

    return 0;
}
