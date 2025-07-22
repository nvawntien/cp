class Solution {
    public int maximumUniqueSubarray(int[] nums) {
        int n = nums.length;
        int sum = 0, ans = 0;
        boolean[] check = new boolean[10001];

        for (int l = 0, r = 0; r < n; r++) {
            while (check[nums[r]]) {
                check[nums[l]] = false;
                sum -= nums[l];
                l++;
            }

            check[nums[r]] = true;
            sum += nums[r];
            ans = Math.max(ans, sum);
        }
        return ans;
    }
}