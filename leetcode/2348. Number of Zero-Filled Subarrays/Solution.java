class Solution {
    public long zeroFilledSubarray(int[] nums) {
        long ans = 0;    
        for (int l = 0, r = 0; r < nums.length; r++) {
            if (nums[r] != 0) {
                l = r;
                continue;
            }

            while (nums[l] != 0 && l < r) {
                l++;
            }

            ans += r - l + 1;
        }

        return ans;
    } 
}