#include <iostream>
#include <vector>
using namespace std;

int main() {
    int t;
    cin >> t;
    while (t--) {
        int x, y;
        cin >> x >> y;
        int val = max(x, y);
        long long ans;
        
        if (val & 1) {
            if (x >= y) {
                ans = 1ll * (val - 1) * (val - 1) + y;
            } else {
                ans = 1ll * val * val - x + 1;
            }
        } else {
            if (x >= y) {
                ans = 1ll * val * val - y + 1;
            } else {
                ans = 1ll * (val - 1) * (val - 1) + x;
            }
        }
        cout << ans << endl;
    }
}