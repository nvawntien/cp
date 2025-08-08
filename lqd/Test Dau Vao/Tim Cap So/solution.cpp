/*
    author: n.vantien
    from: University of Engineering and Technology
*/

#include <bits/stdc++.h>
using namespace std;

int main() {
    ios_base :: sync_with_stdio(0);
    cin.tie(0); cout.tie(0);

    int n, k;
    cin >> n >> k;

    vector <int> arr(n+1);

    for (int i = 1; i <= n; ++i) cin >> arr[i];

    map <int,int> index;

    for (int i = 1; i <= n; ++i) {
        if (index[k - arr[i]]) {
            cout << index[k - arr[i]] << ' ' << i << '\n';
            return 0;
        }

        index[arr[i]] = i;
    }

    cout << "No solution";

    return 0;
}
