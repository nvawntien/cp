#include <iostream>
#include <vector>
using namespace std;

const int LOG = 30;

int binary_lifting(int u, int k, vector <vector<int>>& ancestor) {
    for (int i = 0; (1 << i) <= k; ++i) {
        if (k >> i & 1) u = ancestor[u][i];
    }
    return u;
}

int main() {
    ios_base :: sync_with_stdio(0);
    cin.tie(0); cout.tie(0);

    int n, m;
    cin >> n >> m;

    vector <vector <int>> ancestor(n+1, vector <int> (LOG));

    for (int i = 1; i <= n; ++i) cin >> ancestor[i][0];

    for (int k = 1; k < LOG; ++k) {
        for (int u = 1; u <= n; ++u) {
            ancestor[u][k] = ancestor[ancestor[u][k-1]][k-1];
        }
    } 

    for (int u, k, i = 1; i <= m; ++i) {
        cin >> u >> k;
        cout << binary_lifting(u, k, ancestor) << '\n';
    }

    return 0;
}