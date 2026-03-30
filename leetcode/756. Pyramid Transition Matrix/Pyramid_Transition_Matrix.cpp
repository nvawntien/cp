#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    unordered_map<string, vector<char>> mp;
    unordered_map<string, bool> memo;

    bool dfs(const string& row) {
        if (row.size() == 1) return true;

        if (memo.count(row)) return memo[row];

        int n = row.size();
        vector<vector<char>> choices;

        for (int i = 0; i < n - 1; i++) {
            string pair = row.substr(i, 2);
            if (!mp.count(pair)) {
                memo[row] = false;
                return false;
            }
            choices.push_back(mp[pair]);
        }

        string nextRow(n - 1, ' ');
        function<bool(int)> backtrack = [&](int idx) {
            if (idx == n - 1) {
                return dfs(nextRow);
            }

            for (char c : choices[idx]) {
                nextRow[idx] = c;
                if (backtrack(idx + 1))
                    return true;
            }
            return false;
        };

        bool res = backtrack(0);
        memo[row] = res;
        return res;
    }

    bool pyramidTransition(string bottom, vector<string>& allowed) {
        for (auto& s : allowed) {
            mp[s.substr(0, 2)].push_back(s[2]);
        }
        return dfs(bottom);
    }
};