/*
    author: n.vantien
    from: University of Engineering and Technology
*/

#include <bits/stdc++.h>
using namespace std;

void solve() {
    string number;
    cin >> number;

    vector <char> odd, even;

    for (int i = 0; i < number.length(); ++i) {
        if ((number[i] - '0') & 1) odd.push_back(number[i]);
        else even.push_back(number[i]);
    }

    int n = odd.size(), m = even.size(), i = 0, j = 0;
    string answer = "";

    while (i < n && j < m) {
        if (odd[i] > even[j]) {
            answer += even[j];
            j++;
        }
        else {
            answer += odd[i];
            i++;
        }
    }

    while (i < n) answer += odd[i++];
    while (j < m) answer += even[j++];
    
    cout << answer << '\n';
}

int main() {
    ios_base :: sync_with_stdio(0);
    cin.tie(0); cout.tie(0);

    int T;
    cin >> T;
    while (T--) solve();
    return 0;
}
