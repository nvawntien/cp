class Solution {
    public int maximumLength(int[] nums) {
        int n = nums.length;
        int count = 1, bit = nums[0] & 1, odd = bit;

        for (int i = 1; i < n; i++) {
            if (((nums[i] & 1) ^ bit) == 1) {
                count++;
                bit = nums[i] & 1;
            }  

            odd += nums[i] & 1;          
        }
        
        return Math.max(Math.max(odd, n-odd), count);
    }
}

