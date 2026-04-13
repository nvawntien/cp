class Solution {
    public int countHillValley(int[] nums) {
        int n = nums.length;

        int count = 0;
        ArrayList <Integer> arr = new ArrayList<>();

        arr.add(nums[0]);

        for  (int i = 1; i < n; i++) {
            if (nums[i] == nums[i-1]) continue;
            arr.add(nums[i]);
        }

        for (int i = 1; i < arr.size()-1; i++) {
            if (arr.get(i) > arr.get(i-1) && arr.get(i) > arr.get(i+1)) count++;
            if (arr.get(i) < arr.get(i-1) && arr.get(i) < arr.get(i+1)) count++;
        }

        return count;
    }
}