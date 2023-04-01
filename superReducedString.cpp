#include <cstdio>
#include <iostream>
#include <vector>
#include <algorithm>
#include <cmath>
using namespace std;

char s[107];

bool redukuj() {
	int n = 0;
	for(; s[n]; ++n);
	for(int i = 0; s[i + 1]; ++i) {
		if(s[i] == s[i + 1]) {
			for(int j = i; s[j + 2]; ++j) {
				s[j] = s[j + 2];
			}
			s[n - 1] = s[n - 2] = 0;
			return true;
		}
	}
	return false;
}

int main() {
	ios_base::sync_with_stdio(0);
	cin >> s;
	while(redukuj());
	if(!s[0]) cout << "Empty String\n";
	else cout << s << endl;
	return 0;
}
