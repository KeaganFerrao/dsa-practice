package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{100, 80, 70, 120}
	x := maxProfitAccumulate(arr)

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

// https://www.geeksforgeeks.org/complete-guide-on-array-rotations/
// We consider only right rotate here, left is just similar
// Approach 1: Rotate one by one
// We keep two loops, outer loop runs over the number of times we need to rotate
// Inner loop rotates by one, by keeping last element in a temp variable
// and shifting all elements to the left and replacing the first with temp.
// We do a d = (d % n), because for eg. If the length of the array is 5, then rotate by 6
// is the same as rotate by 1 eventually.
// Time complexity: O(N*d), d is the number of rotations
// Space complexity: O(1)
func rotateArrayOneByOne(arr []int, d int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}

	d = d % n
	for r := 0; r < d; r++ {
		temp := arr[n-1]

		for i := n - 1; i >= 1; i-- {
			arr[i] = arr[i-1]
		}
		arr[0] = temp
	}

	return arr
}

// Approach 2: Using a temp array
// In this we use a temp array to store values, we first loop from the rotation point to the end of the
// array and copy elements to the start of the temp array. In the second loop we copy the remaining
// elements to the temp array.
// Time complexity: O(N)
// Space complexity: O(N)
func rotateArrayUsingTempArray(arr []int, d int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	d = d % n

	temp := make([]int, n)

	// Copy elements from (n-d) to n
	for i := 0; i < d; i++ {
		temp[i] = arr[n-d+i]
	}

	// Copy elements from 0 to (n-d)
	for i := 0; i < (n - d); i++ {
		temp[i+d] = arr[i]
	}

	// Copy elements from temp to original arr
	for i := 0; i < n; i++ {
		arr[i] = temp[i]
	}

	return arr
}

// Approach 3: Using Juggling algorithm
// We use the GCD of the length of array and number of rotations to get the independent cycles
// Time complexity: O(N)
// Space complexity: O(1)
func rotateArrayJugglingAlgorithm(arr []int, d int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	d = d % n

	// GCD gives the number of non-overlaping cycle to move the elements
	cycles := gcd(n, d)

	// Go over the cycles and move elements untill we reach back to the start
	for i := 0; i < cycles; i++ {
		currIdx := i
		currEle := arr[currIdx]

		for {
			nextIdx := (currIdx + d) % n
			nextEle := arr[nextIdx]

			arr[nextIdx] = currEle
			currEle = nextEle
			currIdx = nextIdx

			if currIdx == i {
				break
			}
		}
	}

	return arr
}

// Using the Eudceldian algorithm to find the GCD
func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

// Approach 4: Reversal Algorithm for array rotation
// We first reverse elements from 0 to d-n
// Then we reserver elemets from d-n to n
// Then we reverse the entire array
// Time complexity: O(N)
// Space complexity: O(1)
func rotateArrayReversal(arr []int, d int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	d = d % n

	reverse(arr, 0, n-d-1)
	reverse(arr, n-d, n-1)
	reverse(arr, 0, n-1)

	return arr
}

func reverse(arr []int, start int, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

// https://www.geeksforgeeks.org/move-zeroes-end-array/
// Approach 1: Using a temp array
// Loop through the whole array and if a non-zero element is found then push it to the temp array
// After looping through the full array, then append zeros to the temp array which is the difference
// in lengths of the two arrays
// Time complexity: O(N)
// Space complexity: O(N)
// Approach 2: Two traversal (Shown below)
// We maintain the index of the last non-zero element
// We loop through the whole array and bring all non-zero element to the count index and increment the index
// Time complexity: O(N)
// Space complexity: O(1)
func moveZerosToEndOfArrayTwoTraversal(arr []int) []int {
	// count maintains the index of the last non-zero element in the array
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		// If the ith element is not zero then we replace count element with the ith and incement count
		// if the ith element is zero, we do not increment count, since it is the index of non-zero element
		if arr[i] != 0 {
			arr[count] = arr[i]
			count++
		}
	}

	// Loop from count(which is the index from which the zeros should begin) till then end and add zeros
	for j := count; j < len(arr)-1; j++ {
		arr[j] = 0
	}

	return arr
}

// Approach 3: Single Traversal (Ideal approach)
// Its similary to the above two traversal, but instead of directly assigning arr[count] = arr[i]
// We swap the two values. Doing this all the zeros are moved to the end of the array in one pass
func moveZerosToEndOfArraySingleTraversal(arr []int) []int {
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != 0 {
			// Swap values
			arr[i], arr[count] = arr[count], arr[i]
			count++
		}
	}

	return arr
}

// https://www.geeksforgeeks.org/minimum-increment-k-operations-make-elements-equal/
// Here we get the max element and consider that we need to make all elements equal to
// the max element. Hence we can use simple divisibility checks
// Time complexity: O(N)
// Space Complexity: O(1)
func minIncrementOps(arr []int, k int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	// Find max element
	maxElement := arr[0]
	for i := 0; i < n; i++ {
		if arr[i] > maxElement {
			maxElement = arr[i]
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		// If the max element and the array element are not divisible, then its not possible to make equal
		if (maxElement-arr[i])%k != 0 {
			return -1
		}

		// If its possible then we would need max - element / k operations to make the element
		// equal to max element
		res += (maxElement - arr[i]) / k
	}

	return res
}

// https://www.geeksforgeeks.org/minimum-cost-make-array-size-1-removing-larger-pairs/\
// Here the min cose woul always be min element * (n - 1)
// Since we could always choose the min element and pair it with every other element
// And just remove the other element, making the cose always the min element
// Time complexity: O(N)
// Space complexity: O(1)
func minCost(arr []int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	min := arr[0]
	for i := 0; i < n; i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}

	return min * (n - 1)
}

// https://www.geeksforgeeks.org/print-distinct-elements-given-integer-array/
// Approach 1: We can sort the array, the every duplicate element will be in consecutive order,
// So we can just loop over the array and check if its the same as previous, if yes then do not add
// Time complexity: O(NlogN), for sorting
// Space complexity: O(1)
func printUniqueUsingSorting(arr []int) []int {
	slices.Sort(arr)

	res := []int{arr[0]}
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			continue
		}
		res = append(res, arr[i])
	}

	return res
}

// Approach 2: Using a hashmap, we can loop over the array and only if not exists in the hashmap,
// we can add it to the array.
// Time complexity: O(N)
// Space complexity: O(N), due to hashmap
func printUniqueUsingHashing(arr []int) []int {
	hashMap := make(map[int]int)
	res := make([]int, 0)

	for i := 0; i < len(arr); i++ {
		_, ok := hashMap[arr[i]]
		if !ok {
			res = append(res, arr[i])
			hashMap[arr[i]] = arr[i]
		}
	}

	return res
}

// https://www.geeksforgeeks.org/check-given-array-contains-duplicate-elements-within-k-distance/
// Approach 1: Nexted Looping through each element and checking the window k
// Time complexity: O(N*k), k is the distance
// Space complexity: O(1)
func duplicatesWithinKDistance(arr []int, k int) bool {
	// Loop through each element
	for i := 0; i < len(arr); i++ {
		// Loop through next k elements
		for c := 1; c <= k && (i+c) < len(arr); c++ {
			j := i + c
			// Check if the element is duplicated within the window k
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

// Approach 2: Using a Hashmap
// Time Complexity: O(N)
// Space Complexity: O(N)
func duplicatesWithinKDistanceUsingHashSet(arr []int, k int) bool {
	set := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		// Check if exists in hash map
		_, ok := set[arr[i]]
		if ok {
			return true
		}

		// Add to hash map
		set[arr[i]] = arr[i]

		// Delete from hash map when it goes beyond the k window
		if i >= k {
			delete(set, arr[i-k])
		}
	}
	return false
}

// https://www.geeksforgeeks.org/rearrange-array-such-that-even-positioned-are-greater-than-odd/
// We consider 1 based array indexing for position
// Approach 1: Using sorting, and picking elements one by one from both ends of the array using 2 pointers
// Time Complexity: O(NlogN), for sorting
// Space Complexity: O(N), for res array
func rearrangeUsingSorting(arr []int) []int {
	// Sort the array
	slices.Sort(arr)

	n := len(arr)
	res := make([]int, n)

	start := 0
	end := n - 1

	// Since even positions have to be greater than prev and the array is sorted, we start putting
	// element from the end into even positions and from the start into odd positions
	for i := 0; i < n; i++ {
		if (i+1)%2 == 0 {
			// Even position based on 1-based indexing
			res[i] = arr[end]
			end--
		} else {
			// Odd position
			res[i] = arr[start]
			start++
		}
	}

	return res
}

// Approach 2: Using Swapping
// We loop through the whole array, and swap where the condition is not satisfied
// Time complexity: O(N)
// Space compelexity: O(1)
func rearrangeUsingSwapping(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		if (i+1)%2 == 0 {
			// Even position based on 1-based array indexing
			if arr[i] < arr[i-1] {
				// If the condition is not satisfied, i.e. even position greater than prev element
				// Then swap with prev element to satisfy condition
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		} else {
			// Odd position
			if arr[i] > arr[i-1] {
				// If the condition is not satisfied, i.e. odd position less than prev element
				// Then swap with prev element to satisfy condition
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
	}

	return arr
}

// https://www.geeksforgeeks.org/sum-of-all-subarrays/
// Approach 1: Generate all subarrays, and accumulate a sum
// Time Complexity: O(N^2)
// Space Complexity: O(1)
func sumAllSubarrays(arr []int) int {
	total := 0

	// i is the start index for a subarray
	for i := 0; i < len(arr); i++ {
		temp := 0
		// j is the end index for a subarray
		// So the subarray is from arr[i:j+1]
		for j := i; j < len(arr); j++ {
			/*
				For example, arr = [1,2,3]
				Iteration 1:
				(i,j) => (0,0), (0,1), (0,2)
				temp =>  1, 1+2, 1+2+3
				total => 1 + (1+2) + (1+2+3)

				Iteration 2:
				(i,j) => (1,1), (1,2)
				temp =>  2, 2+3
				total => 1 + (1+2) + (1+2+3) + 2 + (2+3)

				Iteration 3:
				(i,j) => (2,2)
				temp =>  3
				total => 1 + (1+2) + (1+2+3) + 2 + (2+3) + 3

				Final Total = 20

				This is a good optimization to accumulate total here itself rather than
				computing all subarrays and then looping again over all to sum
			*/
			temp += arr[j]
			total += temp
		}
	}

	return total
}

// Approach 2: Using Formula
// Each element arr[i] contribution in the total subarray sum is arr[i]*(i+1)*(n-i)
// Hence we can just loop once over each element and calculate its total
// Time Complexity: O(N)
// Space Complexity: O(1)
func sumAllSubarraysFormula(arr []int) int {
	total := 0

	n := len(arr)
	for i := 0; i < n; i++ {
		total += arr[i] * (i + 1) * (n - i)
	}

	return total
}

// https://www.geeksforgeeks.org/stock-buy-sell/
// Appraoch 1: Using recursion and consider all valid profit combinations (Inefficient approach, but need to understand the pattern)
// We start with the whole array of prices
// i represents a buy day, j represents a sell day
// We check for price[j] > price[i], this means that we could peform a valid buy-sell for profit
// After we get one combination of i,j. We recursively check for all the days before and all the days
// after the range i,j to find the max profits of those ranges as well. We keep a track of the max by storing and returning
// from the recursive calls and in the loop.
// So inshort we check every possibility for buying and selling for each day.
// Time complexity: O(2^N), since the recursion further breaks the problem into 2 more subproblems.
// Space complexity: O(N), due to recursion
func maxProfit(prices []int, start int, end int) int {
	// Keep a track of max profit
	res := 0

	for i := start; i < end; i++ {
		for j := i + 1; j <= end; j++ {
			// Check if it could be a profitable transaction
			if prices[j] > prices[i] {
				// Calculate the current profit => (prices[j] - prices[i])
				// Recursively calculate the left and right days profits if any
				curr := (prices[j] - prices[i]) + maxProfit(prices, start, i-1) + maxProfit(prices, j+1, end)
				// Store the max profit seen so far, since for different i, j combinations there can be multiple
				// profitable days and we need just the maximum profit
				res = max(res, curr)
			}
		}
	}
	return res
}

// Approach 2: Using local minima and local maxima
// Loop throught the whole list and check where there are increases and decreases in prices
// If there is an increase in price, we can buy at the previous value, if decrease we can sell
// Time complexity: O(N)
// Space complexity: O(1)
func maxProfitLocalMinMax(prices []int) int {
	n := len(prices)
	lMin := prices[0]
	lMax := prices[0]
	res := 0

	i := 0
	for i < n-1 {
		// Find the local minima, if the next element is greater than the current, then current is the local minima
		// We can buy at this value
		// For conditions where the next value could be lesser than the currently selected value
		// eg. [100, 40, 30]
		// Here we choose 40 as the local minima, then we choose 40 as the local maxima also in the loop below
		// Hence that res becomes zero, then again 30 will be selected as the local minma, so this condition is handled
		for i < n-1 && prices[i] >= prices[i+1] {
			i++
		}
		lMin = prices[i]

		// Find the local maxima, if the next element is less than the current, then current is the local maxima
		// We can sell at this value
		for i < n-1 && prices[i] <= prices[i+1] {
			i++
		}
		lMax = prices[i]

		// Profit
		res += (lMax - lMin)
	}

	return res
}

// Approach 3: Accumulate profit (Expected solution)
// Here we loop through the whole array, and if the current price is greater than the previous we consider
// buying at i-1 and selling at i. We do that for the while array and accumulate the total costs(profit)
// i.e for a buy we subract the price and for sell we add the price of stock.
// So at the end we will have a accumulate maximum profit
// Time complexity: O(N)
// Space complexity: O(1)
func maxProfitAccumulate(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		// If its resulting in a profit we accumulate the profit
		if prices[i] > prices[i-1] {
			res += (prices[i] - prices[i-1])
		}
	}

	return res
}
