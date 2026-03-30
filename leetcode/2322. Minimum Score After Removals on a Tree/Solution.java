class Solution {
    int[] xor;
    int[] depth;
    int[][] up;
    final int LOG = 11;
    ArrayList<ArrayList<Integer>> tree = new ArrayList<>();

    public int minimumScore(int[] nums, int[][] edges) {
        int n = nums.length;
        xor = new int[n];
        depth = new int[n];
        up = new int[n][LOG];

        for (int i = 0; i < n; i++) {
            tree.add(new ArrayList<>());
        }

        for (int[] edge : edges) {
            tree.get(edge[0]).add(edge[1]);
            tree.get(edge[1]).add(edge[0]);
        }

        for (int i = 0; i < n; i++) xor[i] = nums[i];
        depth[0] = 0;
        dfs(0);

        // swap để đảm bảo edges[i][1] là con của edges[i][0]
        for (int i = 0; i < edges.length; i++) {
            if (depth[edges[i][0]] > depth[edges[i][1]]) swap(edges[i], 0, 1);
        }

        int total = xor[0];
        int ans = Integer.MAX_VALUE;

        for (int i = 0; i < edges.length; i++) {
            int u = edges[i][1];
            for (int j = i + 1; j < edges.length; j++) {
                int v = edges[j][1];
                int par = lca(u, v);
                int a, b, c;

                if (par == u) {
                    a = xor[v];
                    b = xor[u] ^ xor[v];
                    c = total ^ xor[u];
                } else if (par == v) {
                    a = xor[u];
                    b = xor[v] ^ xor[u];
                    c = total ^ xor[v];
                } else {
                    a = xor[u];
                    b = xor[v];
                    c = total ^ xor[u] ^ xor[v];
                }

                int max = Math.max(a, Math.max(b, c));
                int min = Math.min(a, Math.min(b, c));
                ans = Math.min(ans, max - min);
            }
        }

        return ans;
    }

    private void dfs(int u) {
        for (int v : tree.get(u)) {
            if (v != up[u][0]) {
                depth[v] = depth[u] + 1;
                up[v][0] = u;
                for (int j = 1; j < LOG; j++) {
                    up[v][j] = up[up[v][j - 1]][j - 1];
                }
                dfs(v);
                xor[u] ^= xor[v];
            }
        }
    }

    private int lca(int u, int v) {
        if (depth[u] < depth[v]) {
            int temp = u;
            u = v;
            v = temp;
        }

        for (int j = LOG - 1; j >= 0; j--) {
            if (depth[u] - (1 << j) >= depth[v]) {
                u = up[u][j];
            }
        }

        if (u == v) return u;

        for (int j = LOG - 1; j >= 0; j--) {
            if (up[u][j] != up[v][j]) {
                u = up[u][j];
                v = up[v][j];
            }
        }

        return up[u][0];
    }

    private void swap(int[] arr, int i, int j) {
        int tmp = arr[i];
        arr[i] = arr[j];
        arr[j] = tmp;
    }
}

