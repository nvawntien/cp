class Solution {
public:
    string makeFancyString(string s) {
        string str = "";
        int n = 0;
        for (int i = 0; i < s.length(); i++) {
            if (i >= 2 && s[i] == str[n-1] && str[n-1] == str[n-2]) continue;
            str += s[i];
            n++;
        }

        return str;
    }
};