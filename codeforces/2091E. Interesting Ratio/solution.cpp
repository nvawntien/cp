/*
    Author: hoang.ngzz
    From : HaNoi Open University - FITHOU
    Created at: 26/05/2025 -- 22:28
*/

#include <bits/stdc++.h>
using namespace std;
#define int long long
#define F first
#define S second
#define pb push_back
#define vi vector<int>
#define vvi vector<vector<int>>
#define di deque<int>
#define endl '\n'
#define sz(x) ((int)x.size())
#define all(p) p.begin(), p.end()
#define print(a)          \
    for (auto x : a)      \
        cout << x << " "; \
    cout << endl
#define print1(a)    \
    for (auto x : a) \
    cout << x.F << " " << x.S << endl
#define input(a, x, y)          \
    for (int i = x; i < y; i++) \
        cin >> a[i] " ";        \
    cout << endl
#define no "NO\n"
#define yes "YES\n"

vector<bool> generatePrimes(int maxLimit)
{
    vector<bool> isPrime(maxLimit + 1, true);
    isPrime[0] = isPrime[1] = false;
    for (int i = 2; i * i <= maxLimit; i++)
    {
        if (isPrime[i])
        {
            for (int j = i * i; j <= maxLimit; j += i)
            {
                isPrime[j] = false;
            }
        }
    }
    return isPrime;
}

int countInterestingRatios(int n, const vector<bool> &isPrime)
{
    int count = 0;
    for (int a = 1; a <= n; a++)
    {
        for (int b = 2 * a; b <= n; b += a)
        {
            int gcd = a;
            int lcm = b;
            int ratio = lcm / gcd;
            if (isPrime[ratio])
            {
                count++;
            }
        }
    }
    return count;
}

int32_t main()
{
    ios_base::sync_with_stdio(0);
    cin.tie(0);
    cout.tie(0);
    int t;
    cin >> t;

    vector<int> testCases(t);
    int maxN = 0;
    for (int i = 0; i < t; i++)
    {
        cin >> testCases[i];
        maxN = max(maxN, testCases[i]);
    }

    vector<bool> isPrime = generatePrimes(10000000);

    for (int n : testCases)
    {
        cout << countInterestingRatios(n, isPrime) << endl;
    }

    return 0;
}