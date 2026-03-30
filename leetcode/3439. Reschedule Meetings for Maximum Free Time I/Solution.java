class Solution {
    public int maxFreeTime(int eventTime, int k, int[] startTime, int[] endTime) {
        ArrayList <Integer> gap = new ArrayList<>();
        int len = startTime.length;

        gap.add(startTime[0]);

        for (int i = 0; i < len-1; i++) {
            gap.add(startTime[i+1]-endTime[i]);
        }

        gap.add(eventTime - endTime[len-1]);

        int windowSum = 0;

        for (int i = 0; i <= k ; i++) {
            windowSum += gap.get(i);
        }

        int maxSum = windowSum;

        for (int i = k+1; i < gap.size(); i++) {
            windowSum += gap.get(i) - gap.get(i-k-1);
            maxSum = Math.max(windowSum, maxSum);
        }
        return maxSum;
    }   
}