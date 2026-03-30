import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.util.StringTokenizer;

public class Main {
    static void solve() {
        int n = reader.nextInt();
        int arr[] = new int[n+1];

        for (int i = 1; i <= n; i++) {
            arr[i] = reader.nextInt();
        }

        int pre[] = new int[n+1];
        int suf[] = new int[n+2];

        pre[0] = arr[1];
        suf[n+1] = arr[n];

        for (int i = 1; i <= n; i++) {
            pre[i] = Math.min(pre[i-1], arr[i]);
        }

        for (int i = n; i >= 1; i--) {
            suf[i] = Math.max(suf[i+1], arr[i]);
        }

        for (int i = 1; i <= n; i++) {
            if (arr[i] == pre[i] || arr[i] == suf[i]) {
                writer.print("1");
            } 
            else {
                writer.print("0");
            }
        }

        writer.println();
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