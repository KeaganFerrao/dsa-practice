package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 2, 3, 4, 5, 5}
	x := reverseArray(arr)

	fmt.Println(x)
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
	largest := -1
	secondLargest := -1
	thirdLargest := -1

	for _, v := range arr {
		if v > largest {
			thirdLargest = secondLargest
			secondLargest = largest
			largest = v
		} else if v < largest && v > secondLargest {
			thirdLargest = secondLargest
			secondLargest = largest
		} else if v < secondLargest && v > thirdLargest {
			thirdLargest = v
		}
	}

	return largest, secondLargest, thirdLargest
}

// https://www.geeksforgeeks.org/leaders-in-an-array/
// Approach 1 (Naive approach): Use a nested loop to go through each element and chech if its greater
// than all its rightmost elements, if yes then add it to the result set
// Time complexity: O(N^2)
// Space complexity: O(1)
// Approach 2 (Suffix maximum): Scan the array from right to left and keep a track of the maximum
// element till now and add to the result set. Once the maximum changes, add that value to the
// result set as well. At the end reverse the array and return it.
// In short, if the element is greater than that maximum, means its greater that all the others
// as well. Hence it should be included in the leaders. The last elemet=nt by default is a leader
// since it has no rightmost elements
// Time compexity: O(N)
// Space complexity: O(1)
func leadersInArray(arr []int) []int {
	result := make([]int, 0)
	if len(arr) == 0 {
		return result
	}

	max := arr[len(arr)-1]
	result = append(result, max)

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > max {
			result = append(result, arr[i])
			max = arr[i]
		}
	}

	// Reverse the array
	start := 0
	end := len(arr) - 1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}

	return result
}

// https://www.geeksforgeeks.org/program-check-array-sorted-not-iterative-recursive/
// Loop through the whole array and compare current and next element and check if you get
// an unsorted pair, if yes then the array is not sorted.
// Time comlexity: O(N)
// Space complexity: O(1)
func checkIfSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true
}

// https://www.geeksforgeeks.org/remove-duplicates-sorted-array/
// Approach 1: Use a hashmap to keep a track of visited elements, So if the element is in the hashmap
// that means we have already seen it, so its a duplicate and do not include it in the array.
// Time complexity: O(N)
// Space complexity: O(N), for the hashmap
// Approach 2: Since the array is sorted, we can avoid using a hasmap, since all the duplicate elements
// would be in a continuous order, so we can just compare the current element with the previous element
// to check for duplicates
// Here we loop and build the array in place from 0 to idx
// Time complexity: O(N)
// Space complexity: O(1)
func removeDuplicatesFromSortedArray(arr []int) []int {
	idx := 1
	for i := 1; i < len(arr); i++ {
		// Here, we build the array in place, such that o to idx will have unique elements
		if arr[i] != arr[i-1] {
			arr[idx] = arr[i]
			idx++
		}
	}

	// All the elements from 0 to idx(not included) will be sorted with unique values
	return arr[:idx]
}

// https://www.geeksforgeeks.org/program-to-reverse-an-array/
// Keep two pointers, one at the start and other at the end and contimuously swap start and end elements
// Time complexity: O(N)
// Space complexity: O(1)
func reverseArray(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	start := 0
	end := len(arr) - 1

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}

	return arr
}

func rotateArray(arr []int, d int) {

}
