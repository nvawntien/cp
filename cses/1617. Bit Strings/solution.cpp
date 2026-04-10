#include <iostream>
using namespace std;

int main() {
    int n;
    // use binary pow to calculate 2^n
    cin >> n;
    const int mod = 1e9 + 7;
    long long ans = 1, pow = 2;

    while (n) {
        if (n & 1) {
            ans = (ans * pow) % mod;
        }
        pow = (pow * pow) % mod;
        n >>= 1;
    }

    cout << ans << '\n';
}