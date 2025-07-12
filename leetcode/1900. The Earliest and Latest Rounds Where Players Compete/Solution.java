class Solution {
    private int earliest = Integer.MAX_VALUE;
    private int latest = Integer.MIN_VALUE;

    public int[] earliestAndLatest(int n, int firstPlayer, int secondPlayer) {
        int[] players = new int[n];
        for (int i = 0; i < n; i++) players[i] = i+1;
        backtrack(players, firstPlayer, secondPlayer, 1);
        return new int[] {earliest, latest};
    }

    private void backtrack(int[] players, int firstPlayer, int secondPlayer, int round) {
        int n = players.length;
        int[] nextPlayers = new int[n/2 + n%2];
        if (n < 2) return;

        loop: for (int mask = 0; mask < (1 << (n/2)); mask++) {
            int bit = 1;
            for (int i = 0; i < n/2; i++) {
                if (players[i] == firstPlayer && players[n-i-1] == secondPlayer) {
                    earliest = Math.min(earliest, round);
                    latest = Math.max(latest, round);
                    break loop;
                }
                else if ((players[i] == firstPlayer || players[i] == secondPlayer) && (mask & bit) == 0) {
                    continue loop;
                }
                else if ((players[n-i-1] == firstPlayer || players[n-i-1] == secondPlayer) && (mask & bit) != 0) {
                    continue loop;
                }
                
                nextPlayers[i] = (mask & bit) != 0 ? players[i] : players[n-i-1];
                bit = bit << 1;
            }

            if (n%2 != 0) nextPlayers[n/2] = players[n/2];
            Arrays.sort(nextPlayers);
            backtrack(nextPlayers, firstPlayer, secondPlayer, round+1);
        }
    }
}