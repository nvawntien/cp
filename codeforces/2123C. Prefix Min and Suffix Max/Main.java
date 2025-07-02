import java.util.Scanner;

public class Main {
    static void solve(Scanner in) {
        int n = in.nextInt();
        int arr[] = new int[n+1];

        for (int i = 1; i <= n; i++) {
            arr[i] = in.nextInt();
        }

        int pre[] = new int[n+1], suf[] = new int[n+2];
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
                System.out.print("1");
            } else {
                System.out.print("0");
            }
        }

        System.out.println();
    }

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int T = in.nextInt();

        while (T-- > 0) {
            solve(in);
        }

        in.close();
    }
}