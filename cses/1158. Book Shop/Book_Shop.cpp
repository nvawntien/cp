/*
	author: n.vantien
	from: University of Engineering and Technology
*/

#include <iostream>
#include <vector>
using namespace std;

int main() {
	ios_base :: sync_with_stdio(0);
	cin.tie(0); cout.tie(0);

	int n, x;
	cin >> n >> x;

	vector <int> price(n+1), pages(n+1);

	for (int i = 1; i <= n; ++i) cin >> price[i];
	for (int i = 1; i <= n; ++i) cin >> pages[i];

	vector <vector <int>> dp(n+1, vector <int> (x+1, 0));

	for (int i = 1; i <= n; ++i) {
		for (int j = 1; j <= x; ++j) {
			dp[i][j] = dp[i-1][j];
			if (j >= price[i]) {
				dp[i][j] = max(dp[i][j], dp[i-1][j-price[i]] + pages[i]);
			}
		}
	}

	cout << dp[n][x] << '\n';
	return 0;
}