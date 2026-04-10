#include <iostream>
#include <map>
#include <algorithm>
using namespace std;

int main() {
    string str;
    cin >> str;
    map<char, int> freq;

    for (char &c : str) {
        freq[c]++;
    }

    string ans = "";

    int check = 0;

    for (auto &p : freq) {
        if (p.second & 1) {
            check++;
        }

        for (int i = 0; i < p.second / 2; i++) {
            ans += p.first;
        }
    }

    if (check > 1 || (check == 1 && str.size() % 2 == 0)) {
        cout << "NO SOLUTION" << endl;
        return 0;
    }
    
    string rev = ans;
    reverse(rev.begin(), rev.end());

    for (auto &p : freq) {
        if (p.second & 1) {
            ans += p.first;
            break;
        }
    }

    ans += rev;
    cout << ans << endl;
}

