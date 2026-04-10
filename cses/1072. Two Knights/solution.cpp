#include <iostream>
using namespace std;

int main() {
    int n;
    cin >> n;

    // 22 68 156 298 506 792
    // 46 88 142 208 286
    // 42 54 66 78
    // 12 12 12

    for (int i = 1; i <= n; i++) {
        cout << 1ll * i * i  * (i * i - 1) / 2 - 4ll * (i-1) * (i - 2) << '\n';
    }
}