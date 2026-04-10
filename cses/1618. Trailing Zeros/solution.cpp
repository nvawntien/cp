#include <iostream>
using namespace std;

int main() {
    int n;
    cin >> n;

    // use legendre's formula to calculate the number of trailing zeros in n!
    long long ans = 0;
    for (int i = 5; i <= n; i *= 5) {
        ans += n / i;
    }

    cout << ans << '\n';
}