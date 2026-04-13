class Solution {
public:
    int maximum69Number (int num) {
        int ans = 0;
        vector <int> digits;    
        while (num) {
            digits.push_back(num%10);
            num /= 10;
        }

        reverse(digits.begin(), digits.end());

        for (int &x : digits) if (x == 6)  {
            x = 9;
            break;
        }

        for (int &x : digits) ans = ans * 10 + x;

        return ans;
    }
};