class Solution {    
    final int MAX = Integer.MAX_VALUE;

    class SegmentTree {
        int[] tree;
        int n;

        public SegmentTree(int[] arr) {
            n = arr.length;
            tree = new int[4 * n];
            build(1, 0, n-1, arr);
        }

        private void build(int id, int l, int r, int[] arr) {
            if (l == r) {
                tree[id] = arr[l];
                return;
            }

            int mid = (l + r) >> 1;
            build(id<<1, l, mid, arr);
            build(id<<1|1, mid+1, r, arr);
            tree[id] = Math.min(tree[id<<1], tree[id<<1|1]);
        }

        private void update(int id, int l, int r, int pos, int value) {
            if (pos < l || r < pos) return;
            
            if (l == r) {
                tree[id] = value;
                return;
            }

            int mid = (l+r) >>1;
            update(id<<1, l, mid, pos, value);
            update(id<<1|1, mid+1, r, pos, value);

            tree[id] = Math.min(tree[id<<1], tree[id<<1|1]);
        }

        private int get(int id, int l, int r, int u, int v) {
            if (r < u || v < l) return MAX;

            if (u <= l && r <= v) return tree[id];

            int mid = (l + r) >> 1;
            return Math.min(get(id<<1, l, mid, u, v), get(id<<1|1, mid+1, r, u, v));
        }
    }

    public int numOfUnplacedFruits(int[] fruits, int[] baskets) {
        int n = fruits.length;
        int[][] sortedBaskets = new int[n][2];

        for (int i = 0; i < n; i++) {
            sortedBaskets[i][0] = baskets[i];
            sortedBaskets[i][1] = i;
        }  

        Arrays.sort(sortedBaskets, (a, b) -> Integer.compare(a[0], b[0]));

        int[] pos = new int[n];
        int[] index = new int[n];

        for (int i = 0; i < n; i++) {
            pos[i] = sortedBaskets[i][1];
            index[pos[i]] = i;
        }

        SegmentTree st = new SegmentTree(pos);

        int count = 0;

        for (int fruit : fruits) {
            int low = 0, high = n-1, ans = -1;
            while (low <= high) {
                int mid = (low + high) >> 1;
                if (sortedBaskets[mid][0] >= fruit) {
                    ans = mid;
                    high = mid-1;
                }
                else {
                    low = mid+1;
                }
            }

            if (ans == -1) {
                count++;
                continue;
            }

            int value = st.get(1, 0, n-1, ans, n-1);
            if (value == MAX) count++;
            else {
                st.update(1, 0, n-1, index[value], MAX);
            }
        }

        return count;
    }
}