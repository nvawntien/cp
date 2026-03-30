#include <vector>
#include <algorithm>
#include <unordered_set>

using namespace std;

class Solution {
public:
    int maximizeSquareArea(int m, int n, vector<int>& hFences, vector<int>& vFences) {
        long long MOD = 1e9 + 7;

        hFences.push_back(1);
        hFences.push_back(m);
        vFences.push_back(1);
        vFences.push_back(n);

        sort(hFences.begin(), hFences.end());
        sort(vFences.begin(), vFences.end());

        unordered_set<int> hGaps;
        int N = hFences.size();
        for (int i = 0; i < N; ++i) {
            for (int j = 0; j < i; ++j) {
                hGaps.insert(hFences[i] - hFences[j]);
            }
        }

        long long edgeMax = -1;
        int M = vFences.size();

        for (int i = 0; i < M; ++i) {
            for (int j = 0; j < i; ++j) {
                int gap = vFences[i] - vFences[j];
                if (hGaps.count(gap)) {
                    edgeMax = max(edgeMax, (long long)gap);
                }
            }
        }

        if (edgeMax == -1) {
            return -1;
        }

        return (edgeMax * edgeMax) % MOD;
    }
};