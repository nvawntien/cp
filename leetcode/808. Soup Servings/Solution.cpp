class Solution {
public:
    unordered_map <int, unordered_map<int, double>> memo;

    double P(int a, int b) {
        if (a <= 0 && b > 0) return 1;
        if (a <= 0 && b <= 0) return 0.5;
        if (a > 0 && b <= 0) return 0;

        if (memo.count(a) && memo[a].count(b)) return memo[a][b];   
        double res = 0.25 * (P(a - 100, b) + P(a - 75, b - 25) + P(a - 50, b - 50) + P(a - 25, b - 75));
        memo[a][b] = res;
        return memo[a][b];
    }
    double soupServings(int n) {
        if (n > 4500) return 1;
        return P(n, n);
    }
};