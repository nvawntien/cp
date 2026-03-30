class Solution {
    public int matchPlayersAndTrainers(int[] players, int[] trainers) {
        Arrays.sort(players);
        Arrays.sort(trainers);

        int n = trainers.length, m = players.length, count = 0, index = -1;
        for (int i = 0 ; i < m; i++) {
            boolean check = false;
            int low = index+1, high = n-1; 
            while (low <= high) {
                int mid = (low + high) / 2;
                if (trainers[mid] >= players[i]) {
                    check = true;
                    index = mid;
                    high = mid-1;
                }   
                else {
                    low = mid+1;
                }
            }

            if (check) {
                count++;
            }
        }

        return count;
    }
}