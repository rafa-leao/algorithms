package main

//	array I	-	q1. concatenation of array
// 	[1, 2, 3]
//	[1, 2, 3, 1, 2, 3]
//	final complexity : O(n)
func getConcatenation(nums []int) []int {
	ans := nums
	for _, v := range nums {
		ans = append(ans, v)
	}
	return ans
}

//	array I	-	q2. shuffle the array
//	2n elements	[x1, x2, xn, y1, y2, yn]
//		 		[x1, y1, x2, y2, xn, yn]
//	final complexity :
func shuffle(nums []int, n int) []int {
	var shuffled []int
	for i := range nums {
		if i == n {
			break
		}
		shuffled = append(shuffled, nums[i])
		if i == 0 {
			shuffled = append(shuffled, nums[n])
			continue
		}
		shuffled = append(shuffled, nums[n+i])
	}
	return shuffled
}

//	array I	-	q3. max consecutive ones
// [1,1,0,1,1,1]
// 3
// [1,0,1,1,0,1]
// 2
// final complexity : O(n)
func findMaxConsecutiveOnes(nums []int) int {
	var max int
	var maxes []int
	for k, v := range nums {
		if v == 1 {
			max += 1
			if k == len(nums)-1 {
				maxes = append(maxes, max)
			}
			continue
		}
		maxes = append(maxes, max)
		max = 0
		continue
	}

	max = 0
	for _, v := range maxes {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {}
