#include "general_functions.h"

int min(int a, int b) {
	return a < b ? a : b;
}

int max(int a, int b) {
	return a > b ? a : b;
}

void swap(int arr[], int index_a, int index_b) {
	arr[index_a] = arr[index_a] ^ arr[index_b];
	arr[index_b] = arr[index_a] ^ arr[index_b];
	arr[index_a] = arr[index_a] ^ arr[index_b];
}

