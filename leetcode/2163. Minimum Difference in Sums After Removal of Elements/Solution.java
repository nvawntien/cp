class Solution {
    public long minimumDifference(int[] nums) {
        int n = nums.length;
        int k = n / 3;
        PriorityQueue <Integer> maxLeft = new PriorityQueue<> ((a, b) -> b - a);
        PriorityQueue <Integer> minRight = new PriorityQueue<>();

        long[] left = new long[n];
        long[] right = new long[n];
        long leftSum = 0, rightSum = 0;

        for (int i = 0; i < k; i++) {
            leftSum += nums[i];
            maxLeft.add(nums[i]);
        }

        left[k-1] = leftSum;

        for (int i = k; i < n - k; i++) {
            if (nums[i] < maxLeft.peek()) {
                leftSum += nums[i] - maxLeft.poll();
                maxLeft.add(nums[i]);
            }
            left[i] = leftSum;
        }

        for (int i = n-1; i >= n-k; i--) {
            rightSum += nums[i];
            minRight.add(nums[i]);
        }
        
        right[n-k] = rightSum;

        for (int i = n-k-1; i >= k; i--) {
            if (nums[i] > minRight.peek()) {
                rightSum += nums[i] - minRight.poll();
                minRight.add(nums[i]);
            }

            right[i] = rightSum;
        }

        long minDiff = Long.MAX_VALUE;

        for (int i = k-1; i <= n-k-1; i++) {
            minDiff = Math.min(minDiff, left[i] - right[i+1]);
        }

        return minDiff; 
    }
}