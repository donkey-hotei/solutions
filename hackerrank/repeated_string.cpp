#include <bits/stdc++.h>

using namespace std;

long repeatedString(string s, long n) {
    int cnt = 0;
    int rem_cnt = 0;
    int len = s.size();
    int rem = n % len;

    for (int i = 0; i < len; i++) {
        if (s[i] == 'a')
            cnt++;
        if (i < rem && s[i] == 'a')
            rem_cnt++;
    }

    return  n / len * cnt + rem_cnt;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string s;
    getline(cin, s);

    long n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    long result = repeatedString(s, n);

    fout << result << "\n";

    fout.close();

    return 0;
}

