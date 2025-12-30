import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.util.StringTokenizer;

public class Solution {
    static void solve() {
        int n = reader.nextInt();
        int[] arr = new int[n];
        
        for (int i = 0; i < n; i++) {
            arr[i] = reader.nextInt();
        }

        int[] pre = new int[n];
        int[] min = new int[n];

        min[0] = arr[0];

        for (int i = 1; i < n; i++) {
            min[i] = Math.min(min[i - 1], arr[i]);
        }

        pre[0] = min[0];

        for (int i = 1; i < n; i++) {
            pre[i] = pre[i-1] + min[i];
        }

        long ans = pre[0] + arr[1];

        for (int i = 1; i < n; i++) {
            ans = Math.min(ans, (long) pre[i]);
        }

        writer.println(ans);
    }

    public static void main(String[] args) {
        int T = reader.nextInt();
        for (int i = 0; i < T; i++) solve();
        writer.flush();
    }

    static FastReader reader = new FastReader();
    static PrintWriter writer = new PrintWriter(System.out);

    static class FastReader {
        BufferedReader br;
        StringTokenizer st;

        public FastReader() {
            br = new BufferedReader(new InputStreamReader(System.in));
        }

        String next() {
            while (st == null || !st.hasMoreTokens()) {
                try {
                    st = new StringTokenizer(br.readLine());
                } catch (IOException e){
                    e.printStackTrace();
                }
            }

            return st.nextToken();
        }

        int nextInt() {
            return Integer.parseInt(next());
        }

        long nextLong() {
            return Long.parseLong(next());
        }

        double nextDouble() {
            return Double.parseDouble(next());
        }

        String nextLine() {
            String str = "";
            try {
                str = br.readLine().trim();
            } catch (Exception e) {
                e.printStackTrace();
            }

            return str;
        }
    }
}