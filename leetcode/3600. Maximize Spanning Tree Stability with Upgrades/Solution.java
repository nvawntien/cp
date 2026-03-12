class Solution {
    static class DSU {
        int[] parent;

        public DSU(int n) {
            parent = new int[n];
            for (int i = 0; i < n; i++) {
                parent[i] = i;
            }
        }

        public int find(int v) {
            if (v == parent[v]) {
                return v;
            }
            parent[v] = find(parent[v]);
            return parent[v];
        }

        public boolean union(int u, int v) {
            int rootU = find(u);
            int rootV = find(v);

            if (rootU != rootV) {
                parent[rootV] = rootU;
                return true;
            }
            return false;
        }
    }

    public int maxStability(int n, int[][] edges, int k) {
        DSU dsuTest = new DSU(n);
        for (int[] edge : edges) {
            int u = edge[0], v = edge[1], must = edge[3];
            if (must == 1) {
                if (!dsuTest.union(u, v)) {
                    return -1;
                }
            }
        }

        DSU dsuAll = new DSU(n);
        int need = n;
        for (int[] edge : edges) {
            int u = edge[0], v = edge[1];
            if (dsuAll.union(u, v)) {
                need--;
            }
        }

        if (need > 1) {
            return -1;
        }

        int low = 1;
        int high = 200000;
        int ans = -1;

        while (low <= high) {
            int mid = (low + high) >> 1;
            if (check(n, edges, k, mid)) {
                ans = mid;
                low = mid + 1;
            } else {
                high = mid - 1;
            }
        }

        return ans;
    }

    private boolean check(int n, int[][] edges, int k, int x) {
        int used = 0;
        int upgraded = 0;
        DSU dsu = new DSU(n);

        for (int[] edge : edges) {
            int u = edge[0], v = edge[1], s = edge[2], must = edge[3];
            if (must == 1) {
                if (s < x) {
                    return false;
                }
                if (dsu.union(u, v)) {
                    used++;
                }
            }
        }

        for (int[] edge : edges) {
            int u = edge[0], v = edge[1], s = edge[2], must = edge[3];
            if (must == 0 && s >= x) {
                if (dsu.union(u, v)) {
                    used++;
                }
            }
        }

        for (int[] edge : edges) {
            int u = edge[0], v = edge[1], s = edge[2], must = edge[3];
            if (must == 0 && s * 2 >= x && upgraded < k) {
                if (dsu.union(u, v)) {
                    used++;
                    upgraded++;
                }
            }
        }

        return used == n - 1;
    }
}