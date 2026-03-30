class Solution {
public:
    long long minCost(vector<int>& basket1, vector<int>& basket2) {
        unordered_map <int,int> count;

        for (int& fruit : basket1) count[fruit]++;
        for (int& fruit : basket2) count[fruit]++;

        int minCost = INT_MAX;

        for (auto& entry : count) {
            if (entry.second & 1) return -1;
            minCost = min(minCost, entry.first);
        }

        unordered_map <int,int> count1;

        for (int& fruit : basket1) count1[fruit]++;

        vector <int> swapFruit;

        for (auto& entry : count) {
            int fruit = entry.first;
            int diff = count1[fruit] - count[fruit] / 2;
            for (int i = 0; i < abs(diff); ++i) {
                swapFruit.push_back(fruit);
            }
        } 


        sort(swapFruit.begin(), swapFruit.end());

        long long totalCost = 0;

        for (int i = 0; i < swapFruit.size() / 2; ++i) {
            totalCost += min(swapFruit[i], minCost * 2);
        }

        return totalCost;
    }
};