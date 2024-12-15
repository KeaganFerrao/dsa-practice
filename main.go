package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	arr1 := []int{5, 10, 10, 15, 30}
	arr2 := []int{5, 10, 10, 15, 30}
	x := sortedArrayIntersection(arr1, arr2)

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

// https://www.geeksforgeeks.org/find-element-appears-array-every-element-appears-twice/
// Approach 1: Use through all combinations using a nested loop
// Keep a count for every element and check if the count is 1, if yes return it
// Time complexity: O(N^2)
// Space Complexity: O(1)
func singleElementAmongDoublesNestedLoop(arr []int) int {
	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++
			}
		}

		if count == 1 {
			return arr[i]
		}
	}

	return -1
}

// Approach 2: Using a Hash Map
// Looop over the array and build a hashmap with the occurences
// Then loop over the map and check for the element with only 1 occurence
// Time Complexity: O(N)
// Space Complexity: O(N), hashmap
func singleElementAmongDoublesHashMap(arr []int) int {
	hashMap := make(map[int]int)

	for _, val := range arr {
		v, ok := hashMap[val]
		if !ok {
			hashMap[val] = 1
		} else {
			hashMap[val] = v + 1
		}
	}

	for k, v := range hashMap {
		if v == 1 {
			return k
		}
	}

	return -1
}

// Approach 3: Using XOR (Most efficient)
// Using XOR properties
// XOR of a number with itself is 0
// XOR of a number with 0 is itself
// Also XOR is associative i.e
// 7 ^ 3 ^ 5 ^ 4 ^ 5 ^ 3 ^ 4 = 7 ^ (3 ^ 3) ^ (4 ^ 4) ^ (5 ^ 5) = 7
// Since XOR of number with itself will result in 0
func singleElementAmongDoublesXOR(arr []int) int {
	res := 0

	for _, v := range arr {
		// XOR operation
		res = res ^ v
	}

	return res
}

// https://www.geeksforgeeks.org/find-the-missing-number/
// Approach 1: Using Formula: n * (n + 1) / 2
// This is the formula for sum of all elements from 1 to n
// Time Complexity: O(N)
// Space Complexity: O(1)
func missingNumber(arr []int) int {
	n := len(arr) + 1

	expectedSum := n * (n + 1) / 2

	totalSum := 0
	for _, v := range arr {
		totalSum += v
	}

	return expectedSum - totalSum
}

// Approach 2: Using XOR
// Using XOR properties
// XOR of a number with itself is 0
// XOR of a number with 0 is itself
// XOR all numbers from 1 to n
// XOR all number from the array
// Then xor both these results to get the missing number
// Time Complexity: O(N)
// Space Complexity: O(1)
func missingNumberXOR(arr []int) int {
	xor1 := 0
	xor2 := 0

	n := len(arr) + 1
	for i := 1; i <= n; i++ {
		xor1 = i ^ xor1
	}

	for i := 0; i < n-1; i++ {
		xor2 = arr[i] ^ xor2
	}

	return xor1 ^ xor2
}

// https://www.geeksforgeeks.org/find-a-repeating-and-a-missing-number/
// Approach 1: Using visited array
// We keep a track of the values visited as booleans in a visited array using the index
// If that value is already visited means its repeating
// Then we loop over the visited and check which is missing
// Time complexity: O(N)
// Space complexity: O(N), visited array
func repeatingAndMissing(arr []int) (repeating, missing int) {
	repeating = -1
	missing = -1

	visited := make([]bool, len(arr))

	for _, v := range arr {
		if visited[v-1] {
			repeating = v
		} else {
			visited[v-1] = true
		}
	}

	for i, v := range visited {
		if !v {
			missing = i
		}
	}

	return
}

// https://www.geeksforgeeks.org/find-repetitive-element-1-n-1/
// Approach 1: Using math formula
// First we get the sum of all numbers from 1 to n-1, then we get the sum of all numbers in the array
// Return the difference of both sums
// Time complexity: O(N)
// Space complexity: O(1)
func repetitive(arr []int) int {
	totalSum := -1

	for i := 1; i <= len(arr)-1; i++ {
		totalSum += i
	}

	arrSum := -1
	for _, v := range arr {
		arrSum += v
	}

	return arrSum - totalSum
}

// Approach 2: Using XOR
// We compute the XOR of all number from 1 to n-1, then compute XOR of all numbers in array
// Then we XOR both of these results
// Time complexity: O(N)
// Space complexity: O(1)
func repetitiveXOR(arr []int) int {
	xor1 := 0
	xor2 := 0

	for i := 1; i <= len(arr)-1; i++ {
		xor1 = xor1 ^ i
	}

	for _, v := range arr {
		xor2 = xor2 ^ v
	}

	return xor1 ^ xor2
}

// https://www.geeksforgeeks.org/find-a-sorted-subsequence-of-size-3-in-linear-time/
// Approach 1: Using auxilary arrays
// The idea is that we need to find an element which has a lesser element before it and a greater element after it
// Create 2 arrays smaller and greater which will hold if for a particular index in the arr
// the index smaller than and greater than the index in the main array. If the value in the smaller
// and greater array is -1, that means there exists no values greater or smaller than arr[i]
// At the end we loop over and check for values in the smaller and greater where values is not -1
// Time Complexity: O(N)
// Space Complexity: O(N), Smaller and larger arrays
func sortedSubsequence(arr []int) [3]int {
	n := len(arr)

	// Initialize smaller and larger arrays
	smaller := make([]int, n)
	larger := make([]int, n)

	// min is the first element at the start
	min := 0
	// Smaller has the first element -1 since this does not have any lesser elements before it
	smaller[0] = -1
	for i := 1; i < n; i++ {
		// If any next element is lesser than min so far, that means that index does not have any lesser element that it before
		if arr[i] <= arr[min] {
			min = i
			smaller[i] = -1
		} else {
			smaller[i] = min
		}
	}

	// max is the last element at the start of the loop
	max := n - 1
	// Larger is the last element since it cannot have any greater element after it
	larger[n-1] = -1
	// Start looping from the second last element
	for i := n - 2; i >= 0; i-- {
		if arr[i] >= arr[max] {
			max = i
			larger[i] = -1
		} else {
			larger[i] = max
		}
	}

	res := [3]int{-1, -1, -1}
	for i := 0; i < n; i++ {
		// If we get any element in the array which has a valid element in the smaller and greater arrays then that is the result
		if smaller[i] != -1 && larger[i] != -1 {
			res[0] = arr[smaller[i]]
			res[1] = arr[i]
			res[2] = arr[larger[i]]
			break
		}
	}

	return res
}

// https://www.geeksforgeeks.org/segregate-0s-and-1s-in-an-array-by-traversing-array-once/
// Approach 1: Loop over the array and count the number of zeros
// Then loop over the array again and add that many number of zeros in the start and add remaining ones
// Time complexity: O(N)
// Space complexity: O(1)
// The issue in this solution is that it needs 2 passes and if elements are large size values this would be
// very inefficient

// Approach 2: Using 2 pointers
// We use two pointers low and high. We bring then to the write positions. If elements are not peoper
// we swap them and then continue the loop. We use only a single pass here
// Time complexity: O(N)
// Space complexity: O(1)
func segerateZerosAndOnes(arr []int) []int {
	low := 0
	high := len(arr) - 1

	for low < high {
		// Make low to a index before which all are zeros
		for arr[low] == 0 && low < high {
			low++
		}

		// Make high to a index after which all are ones
		for arr[high] == 1 && low < high {
			high--
		}

		// After bringing low and high to proper indexes before this, now low and high are not proper
		// so we swap and increment low and high then again continue the loop
		if low < high {
			arr[low], arr[high] = arr[high], arr[low]
			low++
			high--
		}
	}

	return arr
}

// https://www.geeksforgeeks.org/stable-binary-sort/
// Create a temp array, loop over the main array and first add all the event elements
// then loop again and add all the odd elements
// At the end copy all elements to the main array
func evenBeforeOddElements(arr []int) []int {
	temp := make([]int, 0, len(arr))

	// First add even elements to temp
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			temp = append(temp, arr[i])
		}
	}

	// Then add odd elements to temp
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			temp = append(temp, arr[i])
		}
	}

	// Copy to main array
	copy(arr, temp)

	return arr
}

// https://www.geeksforgeeks.org/rearrange-positive-and-negative-numbers/
// TODO

// https://www.geeksforgeeks.org/sort-array-wave-form-2/
// Approach 1: Sort the array then swap adjacent elements
// Time complexity: O(N*LogN)
// Space conplexity: O(1)
func sortInWave(arr []int) []int {
	slices.Sort(arr)

	for i := 0; i < len(arr); i += 2 {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}

	return arr
}

// Approach 2: Make sure all even positioned elements are greater than the prev and next odd elements
// Traverse all the even positions (0,2,4,..) and if its lesser than the prev or next odd index elements
// then we swap both
// Time complexity: O(N)
// Space complexity: O(1)
func sortInWave2(arr []int) []int {
	// Loop over all even indexed elements
	for i := 0; i < len(arr); i += 2 {
		// Swap if lesser than prev odd index
		if i > 0 && arr[i] < arr[i-1] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
		}

		// Swap of lesser than next odd index
		if i < len(arr)-1 && arr[i] < arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	return arr
}

// https://www.geeksforgeeks.org/largest-sum-contiguous-subarray/
// Approach 1: Iterate over all sub arrays and find the max sum
// Time complexity: O(N^2)
// Space complexity: O(1)
func maxSubarraySum(arr []int) int {
	maxSum := arr[0]

	for i := 0; i < len(arr); i++ {
		currSum := 0
		for j := i; j < len(arr); j++ {
			currSum += arr[j]
		}
		maxSum = max(currSum, maxSum)
	}

	return maxSum
}

// Approach 2: Using Kadanes Algorithm to find max subarray sum
// The simple logic behind this is that, the max subarray can either be the current element or the previous
// max subarray plus the current element. Which is denoted by `maxEnding = max(maxEnding+arr[i], arr[i])` in code
// GlobalSum keeps track of the max subarray sum seen so far. If the maxEnding which is the max subarray sum ending at
// that position, we update the global sum
// Time complexity: O(N)
// Space complexity: O(1)
func maxSubarraySumKadanes(arr []int) int {
	// Max sum seen so far globally
	globalSum := arr[0]
	// Max subarray sum ending at the index or max sum seen locally
	maxEnding := arr[0]

	for i := 1; i < len(arr); i++ {
		// Max subarray sum can either be the current element or the previous max sum + current element
		maxEnding = max(maxEnding+arr[i], arr[i])
		// If the max sumarray sum locally is greater than the one seen so far then update the global sum
		if maxEnding > globalSum {
			globalSum = maxEnding
		}
	}

	return globalSum
}

// https://www.geeksforgeeks.org/maximum-subarray-sum-array-created-repeated-concatenation/
// Approach 1: Create an array of size n*k and run Kadanes algorithm on the whole array as above
// Time complexity: O(N*k)
// Space complexity: O(N*K)
// Approach 2: Use module operator to get the index
// Instead of creating a larger array by replicating k times, we just loop over the same array n*k elements
// and get the index using arr[i%n]
// Then we just use normal kadanes algorithm the same way
// Time complexity: O(N*k)
// Space complexity: O(1)
func maxSubArraySumRepeated(arr []int, k int) int {
	n := len(arr)

	globalSum := arr[0]
	maxEnding := arr[0]

	for i := 1; i < n*k; i++ {
		// Get actual array index by modulo
		ele := arr[i%n]
		maxEnding = max(ele, ele+maxEnding)
		if maxEnding > globalSum {
			globalSum = maxEnding
		}
	}

	return globalSum
}

// https://www.geeksforgeeks.org/maximum-product-subarray/
// Approach 1: Iterate over all sub arrays and find the max product (as shown above in max sum problem)
// Time complexity: O(N^2)
// Space complexity: O(1)

// Approach 2: Using Kadanes algorithm (A little modified to handle negatives for multiplication)
// It is similar to actual kadanes, just that with maxEnding in Kadanes, we also keep a minEnding value
// Time complexity: O(N)
// Space complexity: O(1)
func maxSubArrayProduct(arr []int) int {
	globalProduct := arr[0]
	maxEnding := arr[0]
	minEnding := arr[0] // Track minimum product for handling negatives

	for i := 1; i < len(arr); i++ {
		if arr[i] < 0 {
			// Swap maxEnding and minEnding when encountering a negative number
			// because multiplying a negative will will make greater smaller and smaller greater
			// So this is required
			maxEnding, minEnding = minEnding, maxEnding
		}

		// Update maxEnding and minEnding
		maxEnding = max(arr[i], maxEnding*arr[i])
		minEnding = min(arr[i], minEnding*arr[i])

		// Update global maximum product
		if maxEnding > globalProduct {
			globalProduct = maxEnding
		}
	}

	return globalProduct
}

// https://www.geeksforgeeks.org/equilibrium-index-of-an-array/
// Approach 1: Using nested loop and find teh sum of all elements upto the start and all elements upto the end
// Time complexity: O(N^2)
// Space complexity: O(1)
func equilibriumIndexNested(arr []int) int {
	for i := 0; i < len(arr); i++ {
		sumStart := 0
		sumEnd := 0

		// Loop from index to end of array
		for j := i + 1; j < len(arr); j++ {
			sumEnd += arr[j]
		}

		// Loop from index to start of array
		for k := i - 1; k >= 0; k-- {
			sumStart += arr[k]
		}

		// If both sums equal, we found the equilibrium
		if sumStart == sumEnd {
			return i + 1
		}
	}

	return -1
}

// Approach 2: Using prefix and suffix sum
// We compute the prefix and suffix sum and then loop over both and check which index is equal
// Time complexity: O(N)
// Space complexity: O(N)
func equilibriumIndexPrefixSuffixSum(arr []int) int {
	n := len(arr)

	prefixSum := make([]int, n)
	suffixSum := make([]int, n)

	prefixSum[0] = arr[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i]
	}

	suffixSum[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + arr[i]
	}

	for i := 0; i < n; i++ {
		if prefixSum[i] == suffixSum[i] {
			return i + 1
		}
	}

	return -1
}

// Approach 3: Optimized prefix and suffix sum
// We compute the prefix and suffix sum here as well but using a different approach
// We compute the entire array sum starting from 1st element
// Then we get the prefix sum by adding normall the elements
// But we get the suffix sum by subtracting that element from the array sum computed
// In this way we would not need extra space like the approach before this
// In basic words, we are just increasing the left portion (By adding the element)
// and decreasing the right portion (By subtrating the element from total) and at each step
// we are checking if both sums are equal
// Time complexity: O(N)
// Space complexity: O(1)
func equilibriumIndexPrefixSuffixSumOptimized(arr []int) int {
	left := 0
	right := 0
	pivot := 0

	// Compute the array sum
	for i := 1; i < len(arr); i++ {
		right += arr[i]
	}

	for pivot < len(arr)-1 && left != right {
		pivot++
		// Right side sum by just subtracting the current element from the pre computed sum
		right -= arr[pivot]
		// left sum is computed by just adding the element normally
		left += arr[pivot-1]
	}

	if left == right {
		return pivot + 1
	}

	return -1
}

// https://www.geeksforgeeks.org/check-if-pair-with-given-sum-exists-in-array/
// Approach 1: Check all possible pairs
// Time complexity: O(N^2)
// Space complexity: O(1)
func pairWithSum(arr []int, target int) ([2]int, bool) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return [2]int{arr[i], arr[j]}, true
			}
		}
	}

	return [2]int{-1, -1}, false
}

// Approach 2: Sorting and 2 pointers
// Sort the array first, then using 2 pointers in the start and end, check the sum of both elements
// Move the pointers based on the sum being less or greater than the target
// Time complexity: O(N*logN), for sorting
// Space complexity: O(1)
func pairWithSumSortingWith2Pointer(arr []int, target int) ([2]int, bool) {
	slices.Sort(arr)

	start := 0
	end := len(arr) - 1

	for start < end {
		sum := arr[start] + arr[end]
		if sum > target {
			end--
		} else if sum < target {
			start++
		} else {
			return [2]int{arr[start], arr[end]}, true
		}
	}

	return [2]int{-1, -1}, false
}

// Iterative binary search
// Time complexity: O(logN)
// Space complexity: O(1)
func binarySearch(arr []int, val int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2
		if val < arr[mid] {
			end = mid - 1
		} else if val > arr[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

// Approach 3: Using sorting and binary search
// Here we sort the array and use binary search to find a pair in the array
// i.e for every arr[i] we binary search (target - arr[i])
// Time complexity: O(N*logN), since we loop over each element and do binary search
// Space complexity: O(1)
func pairWithSumSortingAndBinarySearch(arr []int, target int) ([2]int, bool) {
	slices.Sort(arr)

	for i := 0; i < len(arr); i++ {
		index := binarySearch(arr, target-arr[i])
		if index != -1 && index != i {
			return [2]int{arr[i], target - arr[i]}, true
		}
	}

	return [2]int{-1, -1}, false
}

// Approach 4: Using a hash map
// We maintain a hashmap for (target - arr[i]), and then check in O(1) time if a pair exists
// if yes we return it
// Time complexity: O(N)
// Space complexity: O(N), due to hash map
func pairWithSumHashMap(arr []int, target int) ([2]int, bool) {
	hashMap := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		index, ok := hashMap[arr[i]]
		if ok {
			return [2]int{arr[index], arr[i]}, true
		}

		hashMap[target-arr[i]] = i
	}

	return [2]int{-1, -1}, false
}

// https://www.geeksforgeeks.org/two-elements-whose-sum-is-closest-to-zero/
// A varient of this problem can also be closes sum to a given target!
// In short this means that we need to find a pair which has the absolute sum the least
// Approach 1: Loop over all possible pairs
// Time complexity: O(N^2)
// Space complexity: O(1)
func pairWithSumClosestToZero(arr []int) [2]int {
	res := arr[0] + arr[1]
	pair := [2]int{-1, -1}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			sum := arr[i] + arr[j]
			if math.Abs(float64(sum)) < math.Abs(float64(res)) {
				pair = [2]int{arr[i], arr[j]}
				res = sum
			} else if math.Abs(float64(sum)) == math.Abs(float64(res)) {
				// If we get the same absoulte sum, then we check for the max as needed in the problem
				if sum > res {
					pair = [2]int{arr[i], arr[j]}
					res = sum
				}
			}
		}
	}

	return pair
}

// Approach 2: Sorting and binary search
// First we sort the array
// Then for each element in the array at index i, we perform a binary search in i...n-1
// and update the sum with the mid element
// Then update the left and right based on the current sum value
// Time complexity: O(NlogN), for sorting
// Space complexity: O(1)
func pairWithSumClosestToZeroBinarySearch(arr []int) [2]int {
	slices.Sort(arr)

	sum := math.MaxInt
	pair := [2]int{-1, -1}

	for i := 0; i < len(arr); i++ {
		x := arr[i]

		// For each element perform a binary search in i...n-1
		left := i + 1
		right := len(arr) - 1

		for left <= right {
			mid := (left + right) / 2
			currSum := arr[mid] + x

			// If the curr sum is zero, then we found a perfect match
			if currSum == 0 {
				return [2]int{x, arr[mid]}
			}

			// Update sum based on lesser abs value
			if math.Abs(float64(sum)) > math.Abs(float64(currSum)) {
				sum = currSum
				pair = [2]int{x, arr[mid]}
			} else if math.Abs(float64(sum)) == math.Abs(float64(currSum)) {
				// If same sum, update to be the max sum
				if currSum > sum {
					sum = currSum
					pair = [2]int{x, arr[mid]}
				}
			}

			// Update left and right based on the curr sum
			if currSum < 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return pair
}

// Approach 3: Sorting and 2 pointers
// Time complexity: O(NlogN), for sorting
// Space complexity: O(1)
func pairWithSumClosestToZeroTwoPointer(arr []int) [2]int {
	slices.Sort(arr)
	pair := [2]int{-1, -1}

	i := 0
	j := len(arr) - 1

	// Keep a track of the sum and the absolute value of it
	// initialize it to the sum of first and last element
	sum := arr[i] + arr[j]
	diff := math.Abs(float64(sum))

	for i < j {
		currSum := arr[i] + arr[j]

		// If we find a sum to be equal to 0, that means we found a perfect pair
		if currSum == 0 {
			return [2]int{arr[i], arr[j]}
		}

		// If the abs value of the current sum is lesser then the abs sum so far, then we found a new pair closer to 0
		if math.Abs(float64(currSum)) < diff {
			sum = currSum
			diff = math.Abs(float64(currSum))
			pair = [2]int{arr[i], arr[j]}
		} else if math.Abs(float64(currSum)) == diff {
			// If two value have the same abs sum, then we choose the larger sum value, as needed in the problem
			if currSum > sum {
				sum = currSum
				diff = math.Abs(float64(currSum))
				pair = [2]int{arr[i], arr[j]}
			}
		}

		// Update pointers based on current sum value
		if currSum > 0 {
			j--
		} else {
			i++
		}
	}

	return pair
}

// https://www.geeksforgeeks.org/chocolate-distribution-problem/
// We use a sliding window of size m here, and calculate the difference between the start and end of the window
// Time complexity: O(NlogN), for sorting
// Space complexity: O(1)
func chocolateDistribution(arr []int, m int) int {
	slices.Sort(arr)

	minDiff := math.MaxInt
	for i := 0; i+m-1 < len(arr); i++ {
		if arr[i+m-1]-arr[i] < minDiff {
			minDiff = arr[i+m-1] - arr[i]
		}
	}

	return minDiff
}

// https://www.geeksforgeeks.org/union-of-two-arrays/
// Approach 1: We can use a nested loop one to loop over all elements and inner loop to check if the element exists in array
// Time complexity: O(n+m)^2
// Space complexity: O(1)

// Approach 2: Using a hashmap
// Time compleity: O(n+m)
// Space compexity: O(n+m)
func arrayUnion(arr1 []int, arr2 []int) []int {
	hashMap := make(map[int]bool)

	// Add elements from array 1 to hashmap if not already there
	for i := 0; i < len(arr1); i++ {
		_, ok := hashMap[arr1[i]]
		if !ok {
			hashMap[arr1[i]] = true
		}
	}

	// Add elements from array 2 to hashmap if not already there
	for i := 0; i < len(arr2); i++ {
		_, ok := hashMap[arr2[i]]
		if !ok {
			hashMap[arr2[i]] = true
		}
	}

	// Add to result array
	union := make([]int, 0)
	for k := range hashMap {
		union = append(union, k)
	}

	return union
}

// https://www.geeksforgeeks.org/intersection-of-two-arrays/
// Approach 1: Using 2 Hash maps
// Time complexity: O(n+m)
// Space complexity: O(n)
func arrayIntersection(arr1 []int, arr2 []int) []int {
	// This hashmap to keep track of all elements in arr1
	hashMap := make(map[int]bool)
	// This is a result hashmap
	resultMap := make(map[int]bool)
	result := make([]int, 0)

	// We loop over all elements of arr1 and if not present in hash map we add it
	for i := 0; i < len(arr1); i++ {
		if _, ok := hashMap[arr1[i]]; !ok {
			hashMap[arr1[i]] = true
		}
	}

	// We loop over all elements of arr2 and if present in hashmap and not present in result map we add to result set and the map
	for i := 0; i < len(arr2); i++ {
		if _, ok := hashMap[arr2[i]]; ok {
			if _, ok := resultMap[arr2[i]]; !ok {
				resultMap[arr2[i]] = true
				result = append(result, arr2[i])
			}
		}
	}

	return result
}

// Approach 2: Using 1 Hash map
// This is an optimized version of the 2 hashmap version
// In this we just remove the result map, and just after adding to the result set we delete it from the arr1 hashmap
// Time complexity: O(n+m)
// Space complexity: O(n)
func arrayIntersectionOneHashMap(arr1 []int, arr2 []int) []int {
	hashMap := make(map[int]bool)
	result := make([]int, 0)

	for i := 0; i < len(arr1); i++ {
		if _, ok := hashMap[arr1[i]]; !ok {
			hashMap[arr1[i]] = true
		}
	}

	for i := 0; i < len(arr2); i++ {
		if _, ok := hashMap[arr2[i]]; ok {
			result = append(result, arr2[i])
			delete(hashMap, arr2[i])
		}
	}

	return result
}

// https://www.geeksforgeeks.org/union-of-two-sorted-arrays/
// Approach 1: Using a Hash map, same as above array union
// The only difference is that here we need the result set to be sorted as well
// Time complexity: O((n+m)*log(n+m)), due to sorting the result set
// Space complexity: O(n+m), due to hash map
func sortedArrayUnion(arr1 []int, arr2 []int) []int {
	hashMap := make(map[int]bool)
	result := make([]int, 0)

	for i := 0; i < len(arr1); i++ {
		if _, ok := hashMap[arr1[i]]; !ok {
			hashMap[arr1[i]] = true
		}
	}

	for i := 0; i < len(arr2); i++ {
		if _, ok := hashMap[arr2[i]]; !ok {
			hashMap[arr2[i]] = true
		}
	}

	for k := range hashMap {
		result = append(result, k)
	}

	slices.Sort(result)

	return result
}

// Approach 2: Using merge step of merge sort
// The idea is to keep to pointers to the 2 arrays and check if the a value is equal to the previous value
// Then do not add to result set and move that pointer ahead, then based on which is lesser element between two array we
// add them to the result set
// Time complexity: O(n+m)
// Space complexity: O(1)
func mergeLikeArrayUnion(arr1 []int, arr2 []int) []int {
	n := len(arr1)
	m := len(arr2)

	// Two pointer for the 2 sorted arrays
	i := 0
	j := 0

	result := make([]int, 0)

	for i < n && j < m {
		// If same as previous element in arr1
		if i > 0 && arr1[i-1] == arr1[i] {
			i++
			continue
		}

		// If same as previous element in arr2
		if j > 0 && arr2[j-1] == arr2[j] {
			j++
			continue
		}

		if arr1[i] < arr2[j] {
			// If arr1 element is lesser we add that to result set and move ahead
			result = append(result, arr1[i])
			i++
		} else if arr1[i] > arr2[j] {
			// If arr2 element is lesser we add to resukt set and move ahead
			result = append(result, arr2[j])
			j++
		} else {
			// If both are equal, we add any one to the result set and move both pointers ahead
			result = append(result, arr1[i])
			i++
			j++
		}
	}

	// Check for remaining element in arr1 and add them if required
	for i < n {
		if i > 0 && arr1[i-1] == arr1[i] {
			i++
			continue
		}
		result = append(result, arr1[i])
		i++
	}

	// Check for remaining element of arr2 and add them if required
	for j < m {
		if j > 0 && arr2[j-1] == arr2[j] {
			j++
			continue
		}
		result = append(result, arr2[j])
		j++
	}

	return result
}

// https://www.geeksforgeeks.org/intersection-of-two-sorted-arrays/
// Approach 1: We can use the same solution as for an unsorted array using 1 hash map
// Time complexity: O(n+m)
// Space complexity: O(n)

// Approach 2: Using merge step of merge sort
// Same as above for sorted array union, only here we just add if both elements are equal
// Time complexity: O(n+m)
// Space complexity: O(1)
func sortedArrayIntersection(arr1 []int, arr2 []int) []int {
	n := len(arr1)
	m := len(arr2)

	i := 0
	j := 0

	result := make([]int, 0)

	for i < n && j < m {
		if i > 0 && arr1[i-1] == arr1[i] {
			i++
			continue
		}

		if j > 0 && arr2[j-1] == arr2[j] {
			j++
			continue
		}

		if arr1[i] < arr2[j] {
			i++
		} else if arr1[i] > arr2[j] {
			j++
		} else {
			result = append(result, arr1[i])
			i++
			j++
		}
	}

	return result
}
