#include <iostream>
#include <vector>
using namespace std;

vector<pair<int, int>> solve(int n, int from, int to, int aux) {
    vector<pair<int, int>> moves;
    if (n == 0) {
        return moves;
    }

    vector<pair<int, int>> moves1 = solve(n-1, from, aux, to);
    moves.insert(moves.end(), moves1.begin(), moves1.end());
    moves.push_back({from, to});
    vector<pair<int, int>> moves2 = solve(n-1, aux, to, from);
    moves.insert(moves.end(), moves2.begin(), moves2.end());
    return moves;
}

int main() {
    int n;
    cin >> n;
    vector<pair<int, int>> moves = solve(n, 1, 3, 2);
    cout << moves.size() << "\n";
    for (const auto& move : moves) {
        cout << move.first << " " << move.second << "\n";
    }
    return 0;
}