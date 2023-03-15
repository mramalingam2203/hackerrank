// C++ program to find all common
// divisors of N numbers
#include <bits/stdc++.h>
using namespace std;

// Function to calculate gcd of
// two numbers
int gcd(int a, int b)
{
	if (a == 0)
		return b;
	return gcd(b % a, a);
}

// Function to print all the
// common divisors
void printAllDivisors(int arr[], int N,int first_array[])
{
	// Variable to find the gcd
	// of N numbers
	int g = arr[0];

	// Set to store all the
	// common divisors
	set<int> divisors;

	// Finding GCD of the given
	// N numbers
	for (int i = 1; i < N; i++) {
		g = gcd(arr[i], g);
	}

	// Finding divisors of the
	// HCF of n numbers
	for (int i = 1; i * i <= g; i++) {
		if (g % i == 0) {
			divisors.insert(i);
			if (g / i != i)
				divisors.insert(g / i);
		}
	}

	// Print all the divisors
	
	int divs_array[divisors.size()];
	
	int i = 0;
	for (auto& it : divisors){
		divs_array[i] =it;
		i++;
	}

	int count = 0;
	int sec_array_size = sizeof(arr) / sizeof(int);
	for (i=0; i < divisors.size(); i++)
		for (int j =0; j < sec_array_size; j++)
			if (first_array[j]%divs_array[i] == 0)
				count++;
	
	cout << count ;



}

// Driver's Code
int main()
{
	int first_arr[] = {2,6};
	int sec_arr[] = { 24,36};

	int n = sizeof(sec_arr) / sizeof(sec_arr[0]);

	// Function to print all the
	// common divisors
	printAllDivisors(sec_arr, n, first_arr);

	return 0;
}
