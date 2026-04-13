#include <iostream>
#include <vector>

using namespace std;

vector<string> gray(int n) {
    if (n == 1) {
        return {"0", "1"};
    }

    vector <string> prev = gray(n-1);
    vector <string> rev(prev.rbegin(), prev.rend());
    
    for (auto &code : prev) {
        code = "0" + code;
    }

    for (auto &code : rev) {
        code = "1" + code;
    }

    prev.insert(prev.end(), rev.begin(), rev.end());
    return prev;
}

int main() {
    ios_base::sync_with_stdio(false);
    cout.tie(nullptr);

    int n;
    cin >> n;
    vector<string> res = gray(n); 
    
    for (auto &code : res) {
        cout << code << "\n";
    }
}