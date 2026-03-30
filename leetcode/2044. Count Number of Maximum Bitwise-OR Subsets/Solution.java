class Solution {
    public int countMaxOrSubsets(int[] nums) {
        int n = nums.length;

        int orMax = 0, count = 0;

        for (int mask = 0; mask < (1 << n); mask++) {
            int or = 0;
            for (int i = 0; i < n; i++) {
                if (((mask >> i) & 1) == 1) {
                    or |= nums[i];
                }
            }

            if (or > orMax) {
                count = 1;
                orMax = or;
            }
            else if (or == orMax) {
                count++;
            }
        }

        return count;
    }
}