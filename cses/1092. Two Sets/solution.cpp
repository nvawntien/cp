#include <iostream>
#include <map>
#include <vector>
using namespace std;

int main() {
    int n;
    cin >> n;

    long long sum = 1ll * n * (n + 1) / 2;
    map <int, int> freq;

    for (int i = 1; i <= n; i++) {
        freq[i]++;
    }

    if (sum & 1) {
        cout << "NO\n";
    } else {
        cout << "YES\n";
        long long curr = 0;
        sum /= 2;
        vector <int> ans;

        for (int i = n; i >= 1; i--) {
            if (curr + i <= sum) {
                ans.push_back(i);
                curr += i;
                freq[i]--;
            }
        }

        cout << ans.size() << '\n';
        for (int i : ans) {
            cout << i << ' ';
        }

        cout << '\n';
        vector <int> ans2;
        for (int i = 1; i <= n; i++) {
            if (freq[i] > 0) {
                ans2.push_back(i);
                freq[i]--;
            }
        }
        cout << ans2.size() << '\n';
        for (int i : ans2) {
            cout << i << ' ';
        }
        cout << '\n';
    }
}