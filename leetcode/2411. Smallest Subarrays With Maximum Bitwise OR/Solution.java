class Solution {
    class SegmentTree {
        private int[] tree;
        private int n;

        public SegmentTree(int[] nums) {
            this.n = nums.length;
            tree = new int[4 * n];
            build(nums, 1, 0, n-1);
        }

        private void build(int[] nums,  int id,  int l, int r) {
            if (l == r) {
                tree[id] = nums[l];
                return;
            }

            int mid = (l + r) >> 1;          
            build(nums, id<<1, l, mid);
            build(nums, id<<1|1, mid+1, r);

            tree[id] = tree[id<<1] | tree[id<<1|1];
        }

        private int get(int id, int l, int r, int u, int v) {
            if (r < u || v < l) return 0;
            if (u <= l && r <= v) return tree[id];

            int mid = (l + r) >> 1;
            return get(id<<1, l, mid, u, v) | get(id<<1|1, mid+1, r, u, v);
        }
    }

    public int[] smallestSubarrays(int[] nums) {
        SegmentTree st = new SegmentTree(nums);
        int n = nums.length;
        int[] ans = new int[n];

        for (int i = 0; i < n; i++) {
            int low = 1, high = n-i, len = n-i;
            int value = st.get(1, 0, n-1, i, n-1);
            while (low <= high) {
                int mid = (low + high) >> 1;
                if (st.get(1, 0, n-1, i, i+mid-1) == value) {
                    len = mid;
                    high = mid-1;
                }
                else {
                    low = mid+1;
                }
            }

            ans[i] = len;
        }
        
        return ans;
    }
}