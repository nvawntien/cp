class Solution {
public:
    int makeTheIntegerZero(int num1, int num2) {
        for (int k = 0; k <= 60; ++k) {
            long long sub = num1 - 1ll * k * num2;
            if (sub <= 0) continue;
            if (__builtin_popcountll(sub) <= k && k <= sub) return k;
        }

        return -1;
    }
};