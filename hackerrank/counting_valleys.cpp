#include <bits/stdc++.h>

using namespace std;

// Complete the countingValleys function below.
int countingValleys(int n, string s) {
    int elevation = 0;
    int valley_count = 0;

    for (char const &step: s) {
        if (step == 'U') {
            elevation++;

            if (elevation == 0) {
                valley_count++;
            }
        } else if (step == 'D') {
            elevation--;
        }
    }

    return valley_count;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    string s;
    getline(cin, s);

    int result = countingValleys(n, s);

    fout << result << "\n";

    fout.close();

    return 0;
}

