// C++ program for the above approach
#include <bits/stdc++.h>
using namespace std;

// Function to assign rank to
// array elements
void changeArr(int input[], int N)
{

	// Copy input array into newArray
	int newArray[N];
	copy(input, input + N, newArray);

	// Sort newArray[] in ascending order
	sort(newArray, newArray + N);
	int i;

	// Map to store the rank of
	// the array element
	map<int, int> ranks;
	int rank = 1;

	for(int index = 0; index < N; index++)
	{
		int element = newArray[index];
		// Update rank of element
		if (ranks[element] == 0)
		{
			ranks[element] = rank;
			rank++;
		}
	}

	// Assign ranks to elements
	for(int index = 0; index < N; index++)
	{
		int element = input[index];
		input[index] = ranks[input[index]];
		printf("%d\ ",input[index] 	 );
	}
}

// Driver code
int main()
{

	// Given array arr[]
	int arr[] = { 100, 100, 50, 40, 40, 20, 10};

	int N = sizeof(arr) / sizeof(arr[0]);

	// Function call
	changeArr(arr, N);

	// Print the array elements
	cout << "[";
	for(int i = 0; i < N - 1; i++)
	{
		cout << arr[i] << ", ";
	}
	cout << arr[N - 1] << "]";
	return 0;
}

// This code is contributed by divyeshrabadiya07
