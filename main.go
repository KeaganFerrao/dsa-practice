package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 12, 10}
	x, y, z := largestThreeElements(arr)

	fmt.Println(x, y, z)
}

// https://www.geeksforgeeks.org/find-second-largest-element-array
// Approach 1: (Naive approach using sorting) Sort the array in decreasing order, then the first element is
// the largest element. Loop through the array and the first element not equal to the largest
// would be the second largest.
// Time complexity: O(NlogN) for sorting + O(N) for looping to find second largest = O(NlogN)
// Space complexity: O(1), would depend on the sorting algorithm we choose
// Approach 2: (Two pass search) Loop through the array once and find the largest element
// then lop through the array again and find the next largest distinct element
// Time complexity: O(N) + O(N) for two passes = O(N)
// Space complexity: O(1)
// Approach 3(Shown Below): (One pass search) Loop through the array once and keep a track of the 2 variables
// the largest and second largest and update then according as shown below
// Time complexity: O(N), due to single pass
// Space complexity: O(1)
func secondLargestElement(arr []int) int {
	largest := -1
	secondLargest := -1

	for _, v := range arr {
		if v > largest {
			secondLargest = largest
			largest = v
		} else if v < largest && v > secondLargest {
			secondLargest = v
		}
	}

	return secondLargest
}

// https://www.geeksforgeeks.org/find-the-largest-three-elements-in-an-array/
// Approach 1 (Naive Approach): Three pass traversal, traverse the array three time to find
// the first, second and third largest element.
// Time complexity: O(N) + O(N) + O(N) = O(N), since 3 times traversal
// Space complexity: O(1)
// Arrpoach 2 (Below) One traversal: Keep 3 variable for the largest, second and third.
// Update the values based on the conditions shown below
// Time complexity: O(N), since single traversal
// Space complexity: O(1)
func largestThreeElements(arr []int) (int, int, int) {
	var largest, secondLargest, thirdLargest *int

	for _, v := range arr {
		if largest == nil || v > *largest {
			thirdLargest = secondLargest
			secondLargest = largest
			largest = &v
		} else if secondLargest == nil || (v < *largest && v > *secondLargest) {
			thirdLargest = secondLargest
			secondLargest = &v
		} else if thirdLargest == nil || (v < *secondLargest && v > *thirdLargest) {
			thirdLargest = &v
		}
	}

	return *largest, *secondLargest, *thirdLargest
}
