#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

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

int main()
{
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