class Solution {
    public int maxFreeTime(int eventTime, int[] startTime, int[] endTime) {
        int n = startTime.length;   
        int last = 0;
        int[] left = new int[n+1];
        int[] right = new int[n+1];
        int[] gaps = new int[n+1];


        for (int i = 0; i < n; i++) {
            gaps[i] = startTime[i] - last;
            last = endTime[i];
        }

        gaps[n] = eventTime - last;

        for (int i = 0; i <= n; i++) {
            left[i] = Math.max(left[Math.max(i-1,0)], gaps[i]);
        }

        for (int i = n; i >= 0; i--) {
            right[i] = Math.max(right[Math.min(i+1, n)], gaps[i]);
        }

        int ans = left[n];

        for (int i = 1; i <= n; i++) {
            int busyTime = endTime[i-1] - startTime[i-1];
            if (i == 1) {
                if (right[i+1] >= busyTime) {
                    ans = Math.max(ans, startTime[i]);
                } 
                else {
                    ans = Math.max(ans, startTime[i] - busyTime);
                }
            }
            else if (i == n) {
                if (left[i-2] >= busyTime) {
                    ans = Math.max(ans, eventTime - endTime[i-2]);
                }
                else {
                    ans = Math.max(ans, eventTime - endTime[i-2] - busyTime);
                }
            }
            else {
                if (left[i-2] >= busyTime || right[i+1] >= busyTime) {
                    ans = Math.max(ans, startTime[i] - endTime[i-2]);
                }
                else {
                    ans = Math.max(ans, startTime[i] - endTime[i-2] - busyTime);
                }
            }
        }

        return ans;
    }
}