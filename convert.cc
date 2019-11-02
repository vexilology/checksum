#include <iostream>

using namespace std;

int main()
{
	int a[10], n, i;
	cout << "Number to convert: ";
	cin >> n;
	for(i=0; n>0; i++)
	{
		a[i] = n%2;
		n = n/2;
	}
	cout << "New binary number: ";
	for(i=i-1; i>=0; i--)
	{
		cout << a[i];
	}
	system("pause >> NUL");
}