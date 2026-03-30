class Solution {
    public int maxTotalFruits(int[][] fruits, int startPos, int k) {
        int len = 200005;
        int[] pre = new int[len];

        for (int[] it : fruits) {
            pre[it[0]] = it[1];            
        }

        for (int i = 1; i < len; i++) {
            pre[i] = pre[i-1] + pre[i];
        }

        int ans = 0;

        for (int i = startPos; i <= Math.min(startPos+k , len-1);i++){
            int r = i;
            int x = i-startPos;
            int l = startPos - (k-2*x);
            l = Math.min(l , startPos);
            int sum = pre[r];
            if (l > 0){
                sum -= pre[l-1];
            }
            ans = Math.max(ans , sum);
        }

        for (int i = startPos; i >= Math.max(startPos-k , 0);i--){
            int l = i;
            int x = startPos-i;
            int r = startPos + (k-2*x);
            r = Math.max(Math.min(len-1, r) , startPos);
            int sum = pre[r];
            if (l > 0){
                sum -= pre[l-1];
            }
            ans = Math.max(ans , sum);
        }

        return ans;
    }
}