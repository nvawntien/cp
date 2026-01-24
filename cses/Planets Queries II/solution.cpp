#include <iostream>
#include <vector>
#include <queue>
#include <unordered_map>
using namespace std;

const int LOG = 30;

void findNodesInCycle(int n, vector <int> &to, vector <int> &indeg, vector <bool> &inCycle) {
    queue <int> qnodes;

    for (int i = 1; i <= n; i++) {
        if (indeg[i] == 0) {
            qnodes.push(i);
        }
    }

    while (!qnodes.empty()) {
        int u = qnodes.front();
        qnodes.pop();
        if (--indeg[to[u]] == 0) {
            qnodes.push(to[u]);
        }
    }

    for (int i = 1; i <= n; i++) {
        if (indeg[i] > 0) {
            inCycle[i] = true;
        }
    }
}

void dfs(int u, vector <vector<int>> &adj, vector <int> &depth, vector <int> &root, vector <bool> &inCycle) {
    for (int v : adj[u]) {
        if (!inCycle[v]) {
            depth[v] = depth[u] + 1;
            root[v] = root[u];
            dfs(v, adj, depth, root, inCycle);
        }
    }
}

void computeRootAndDepth(int n, vector <vector<int>> &adj, vector <int> &to, vector <bool> &inCycle, vector <int> &root, vector <int> &depth, vector <int> &pos, vector <int> &cycleSize) {
    vector<bool> vis(n+1, false);
    
    for (int i = 1; i <= n; i++) {
        if (inCycle[i] && !vis[i]) {
            int curr = i, len = 0;
            do {
                vis[curr] = true;
                root[curr] = i;
                pos[curr] = len++;
                curr = to[curr];
            } while (curr != i);
            cycleSize[i] = len;
        }
    }

    for (int i = 1; i <= n; i++) {
        if (inCycle[i]) {
            depth[i] = 0;
            dfs(i, adj, depth, root, inCycle);
        }
    }
}

int jump(int u, int k, vector <vector<int>> &ancestor) {
    for (int j = 0; j < LOG; j++) {
        if (k & (1 << j)) {
            u = ancestor[u][j];
        }
    }
    return u;
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);cout.tie(nullptr);

    int n, q;
    cin >> n >> q;

    vector <vector<int>> ancestor(n+1, vector<int>(LOG));
    vector <int> indeg(n+1, 0);
    vector <int> to(n+1);
    vector <bool> inCycle(n+1, false);
    vector <int> depth(n+1, 0);
    vector <int> root(n+1, 0);
    vector <vector<int>> adj(n+1);
    vector <int> pos(n+1, -1);
    vector <int> cycleSize(n+1, 0);

    for (int i = 1; i <= n; i++) {
        cin >> to[i];
        indeg[to[i]]++;    
        ancestor[i][0] = to[i];
        adj[to[i]].push_back(i);    
    }

    for (int j = 1; j < LOG; j++) {
        for (int i = 1; i <= n; i++) {
            ancestor[i][j] = ancestor[ancestor[i][j-1]][j-1];
        }
    }

    findNodesInCycle(n, to, indeg, inCycle);
    computeRootAndDepth(n, adj, to, inCycle, root, depth, pos, cycleSize);

    for (int _ = 0; _ < q; _++) {
        int u, v;
        cin >> u >> v;

        if (root[u] != root[v]) {
            cout << "-1\n";
            continue;
        }

        if (depth[v] > depth[u]) {
            cout << "-1\n";
            continue;
        }

        int dist = depth[u] - depth[v];
        int ancestor_of_u = jump(u, dist, ancestor);

        if (depth[v]) {
            if (ancestor_of_u == v) {
                cout << dist << "\n";
            } else {
                cout << "-1\n";
            }
        } else {
            cout << (pos[v] - pos[ancestor_of_u] + cycleSize[root[v]]) % cycleSize[root[v]] + dist << "\n";
        }
    }

    return 0;
}