class Solution {
    public long largestSquareArea(int[][] bottomLeft, int[][] topRight) {
        long maxArea = 0;
        int n = bottomLeft.length;

        for (int i = 0; i < n; i++) {
            for (int j = i+1; j < n; j++) {
                int interBLX = Math.max(bottomLeft[i][0], bottomLeft[j][0]);
                int interBLY = Math.max(bottomLeft[i][1], bottomLeft[j][1]);
                int interTRX = Math.min(topRight[i][0], topRight[j][0]);
                int interTRY = Math.min(topRight[i][1], topRight[j][1]);

                int width = interTRX - interBLX;
                int height = interTRY - interBLY;

                if (width > 0 && height > 0) {
                    long side = Math.min(width, height);
                    long area = side * side;
                    if (area > maxArea) {
                        maxArea = area;
                    }
                }
            }
        }

        return maxArea;
    }
}