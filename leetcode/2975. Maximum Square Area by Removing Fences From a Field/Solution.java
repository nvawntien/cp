import java.util.*;

class Solution {
    public int maximizeSquareArea(int m, int n, int[] hFences, int[] vFences) {
        long MOD = 1_000_000_007L;

        int[] h = Arrays.copyOf(hFences, hFences.length + 2);
        h[h.length - 2] = 1;
        h[h.length - 1] = m;
        
        int[] v = Arrays.copyOf(vFences, vFences.length + 2);
        v[v.length - 2] = 1;
        v[v.length - 1] = n;

        Arrays.sort(h);
        Arrays.sort(v);

        Set<Integer> hGaps = new HashSet<>();
        for (int i = 0; i < h.length; i++) {
            for (int j = 0; j < i; j++) {
                hGaps.add(h[i] - h[j]);
            }
        }

        long maxEdge = -1;

        for (int i = 0; i < v.length; i++) {
            for (int j = 0; j < i; j++) {
                int gap = v[i] - v[j];
                if (hGaps.contains(gap)) {
                    maxEdge = Math.max(maxEdge, (long) gap);
                }
            }
        }

        if (maxEdge == -1) {
            return -1;
        }

        return (int) ((maxEdge * maxEdge) % MOD);
    }
}