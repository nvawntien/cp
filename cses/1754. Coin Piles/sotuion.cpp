#include <iostream>
using namespace std;

void solve() {
    int a, b;
    cin >> a >> b;
    // x + 2y = a
    // 2x + y = b
    // x + 2(b - 2x) = a
    // -3x + 2b = a
    // x = (2b - a) / 3
    // y = (2a - b) / 3
    // 2b - a >= 0 && 2a - b >= 0

    if ((a + b) % 3 == 0 && min(a, b) * 2 >= max(a, b)) {
        cout << "YES\n";
    } else {
        cout << "NO\n";
    }
    
}

int main() {
    int t;
    cin >> t;
    while (t--) {
        solve();
    }
}