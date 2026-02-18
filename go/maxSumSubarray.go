package main

import (
	"strings"
	"strconv"
	"fmt"
	"os"
)

func main() {
	
	fmt.Println("Max sum of subarray!")
	
	fmt.Println(maxSumSubarray([]int{2, 1, 5, 1, 3, 2}, 3)) 
	
	// 1,2,3,4,5,6,7,8,9,0,9,8,7,6,6,5,4,4,3,3,2,1 : 3
	numList, size := getNumListNSizeFromUser()
	if numList == nil || size == 0 {
		fmt.Println("Exiting program due to invalid input.")
		os.Exit(1)
	}
	
	fmt.Printf("Max sum of subarray of size %d is %d\n", size, maxSumSubarray(numList, size))

}

/*
 * Given an array of integers and a number k, find the maximum sum of a subarray of size k.
 * For example, if the input array is [1, 2, 3, 4, 5] and k is 2, the maximum sum of a subarray of size 2 is 9 (4 + 5).
 * If the input array is [1, 2, 3, 4, 5] and k is 3, the maximum sum of a subarray of size 3 is 12 (3 + 4 + 5).
 */

func maxSumSubarray(list []int, size int) int {

    if size < 0 || size > len(list) {
        return -1
    }
	sum := 0
	for _, i := range(list[:size]) {
		sum += i
	}
	max_sum := sum

    for i := size; i < len(list); i++ {
        sum += list[i] - list[i-size]
		if sum > max_sum {
			max_sum = sum
		}
    }
	return max_sum
}

/*
 * Function gets input from user in the form of
 *    - an array of integers 
 *    - a number, size of the subarray 
 *   returns the array of integers & size of a subarray
 */
func getNumListNSizeFromUser() ([]int, int) {
	var numList []int
	var str string
	var strList string

	// Get input from user
	fmt.Print("Enter int array list comma seperated : ")
	fmt.Scanln(&strList)
	fmt.Print("Enter subarray size : ")
	fmt.Scan(&str)

	strs := strings.Split(strList, ",")
	size, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		fmt.Printf("Error converting %s to int: %v\n", str, err)
		return nil, 0
	}
	
	for _, i := range strs {
		num, err := strconv.Atoi(strings.TrimSpace(i))
		if err != nil {
			fmt.Printf("Error converting %s to int: %v\n", i, err)
			continue
		}
		numList = append(numList, num)
	}

	if len(numList) == 0 {
		fmt.Println("No valid integers entered.")
		return numList, 0
	}
	return numList, size
}
