#include <iostream>
#include <vector>
#include <queue>
#include <climits>
using namespace std;

vector <long long>  dijkstra(int start, int n, vector < vector<pair<int, int>>> &edges ) {
    priority_queue <pair<long long, int>, vector <pair<long long, int>>, greater <pair<long long, int>>> Q;
    vector <long long> dist(n+1, LLONG_MAX);
    dist[start] = 0;

    Q.push({dist[start], start});

    while (Q.size()) {
        auto edge = Q.top(); 
        Q.pop();

        int u = edge.second;
        long long distance = edge.first;

        if (distance > dist[u]) {
            continue;
        }

        for (auto &[w, v] : edges[u]) {
            if (dist[u] + w < dist[v]) {
                dist[v] = dist[u] + w;
                Q.push({dist[v], v});
            }   
        }
    }

    return dist;
}

int main() {
    ios_base ::sync_with_stdio(0);
    cin.tie(0); cout.tie(0);

    int n, m;
    cin >> n >> m;

    vector <vector<pair<int, int>>> edges(n+1);

    for (int u, v, w, i = 1; i <= m; ++i) {
        cin >> u >> v >> w;
        edges[u].push_back({w, v});
    }

    vector <long long> dist = dijkstra(1, n, edges);
    for (int i = 1; i <= n; ++i) cout << dist[i] << ' ';   
    return 0;
}