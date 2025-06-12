package array_code

//704
func BinarySearch704(nums []int, target int) int {
	index := 0
	numsLength := len(nums) - 1
	for index <= numsLength {
		mid := index + (numsLength-index)/2
		if nums[mid] > target {
			numsLength = mid - 1
		} else if nums[mid] < target {
			index = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func BinarySearchVersionTwo704(nums []int, target int) int {
	index := 0
	numsLength := len(nums) //len(nums) - 1
	for index < numsLength {
		mid := index + (numsLength-index)/2
		if nums[mid] > target {
			numsLength = mid
		} else if nums[mid] < target {
			index = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

//func search(nums []int, target int) int {
//	index := 0
//	length := len(nums) - 1
//	for index <= length {
//		mid := (index + length) / 2
//		if nums[mid] == target {
//			return mid
//		} else if nums[mid] > target {
//			length = mid - 1
//		} else {
//			index = mid + 1
//		}
//	}
//	return -1
//}

func Demo(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
