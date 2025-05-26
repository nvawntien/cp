#include <iostream>
#include <vector>
#include <queue>
#include <climits>
using namespace std;
 
int main() {
    ios_base ::sync_with_stdio(0);
    cin.tie(0); cout.tie(0);
 
    int n, m, q;
    cin >> n >> m >> q;
    
    vector <vector<long long>> dist(n+1, vector <long long> (n+1, LLONG_MAX));

    for (int i = 1; i <= n; ++i) dist[i][i] = 0;

    for (int u, v, w, i = 1; i <= m; ++i) {
        cin >> u >> v >> w;
        dist[u][v] = min(dist[u][v], 1ll*w), dist[v][u] = min(dist[v][u], 1ll*w);
    }

    for (int k = 1; k <= n; ++k) {
        for (int i = 1; i <= n; ++i) {
            for (int j = 1; j <= n; ++j) {
                if (dist[i][k] == LLONG_MAX || dist[k][j] == LLONG_MAX) continue;
                dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j]);
            }
        }
    }

    for (int u, v, i = 1; i <= q; ++i) {
        cin >> u >> v;
        if (dist[u][v] == LLONG_MAX) cout << -1 << '\n';
        else cout << dist[u][v] << '\n';
    }

    return 0;
}