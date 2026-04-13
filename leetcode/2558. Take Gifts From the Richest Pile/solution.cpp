class Solution {
public:
    long long pickGifts(vector<int>& gifts, int k) {
        priority_queue <int> Q;
        for (auto &a : gifts) Q.push(a);
        while (k--) {
            int value = Q.top(); Q.pop();
            Q.push(int(sqrt(value)));
        }

        long long sum = 0;

        while (Q.size()) {
            sum += Q.top(); Q.pop();
        }

        return sum;
    }
};