#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

void backtracking(const string &s, int idx, vector<bool> &used, vector<string> &res, string &curr) {
    if (idx == s.size()) {
        res.push_back(curr);
        return;
    }

    for (int i = 0; i < s.size(); i++) {
        if (used[i] || (i > 0 && s[i] == s[i-1] && !used[i-1])) {
            continue;
        }

        used[i] = true;
        curr.push_back(s[i]);
        backtracking(s, idx + 1, used, res, curr);
        curr.pop_back();
        used[i] = false;
    }
}

int main() {
    ios_base::sync_with_stdio(false);
    cout.tie(nullptr);

    string s;
    cin >> s;

    sort(s.begin(), s.end());
    vector <bool> used(s.size(), false);
    vector <string> res;
    string curr = "";
    backtracking(s, 0, used, res, curr);
    
    cout << res.size() << "\n";
    for (const auto& str : res) {
        cout << str << "\n";
    }
}