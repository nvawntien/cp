class Solution {
    public int mostBooked(int n, int[][] meetings) {
        Arrays.sort(meetings, Comparator.comparingInt(a -> a[0]));

        PriorityQueue <long[]> busy = new PriorityQueue<> ((a,b) -> {
            if (a[0] != b[0]) return Long.compare(a[0], b[0]);
            return Long.compare(a[1], b[1]);
        });

        PriorityQueue <Integer> free = new PriorityQueue<>();

        for (int i = 0; i < n; i++) free.add(i);

        int[] count = new int[n];

        for (int[] meeting : meetings) {
            int start = meeting[0], end = meeting[1];

            while (!busy.isEmpty() && busy.peek()[0] <= start) {
                free.add((int) busy.poll()[1]);
            }

            if (!free.isEmpty()) {
                int room = free.poll();
                count[room]++;
                busy.add(new long[] {end, room});
            }
            else {
                long[] arr = busy.poll();
                count[(int)arr[1]]++;
                busy.add(new long[] {arr[0] + end - start, arr[1]});
            }
        }

        int room = 0, meeting = 0;

        for (int i = 0; i < n; i++) {
            if (count[i] > meeting) {
                room = i;
                meeting = count[i];
            }
        }

        return room;
    }
}