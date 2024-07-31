
package dsa

//import ("fmt")

func BinarySearch(target int, arr[]int) int {
	left := 0;
	right := len(arr);

	for left < right {
		mid := left + (right - left) / 2
		if target == arr[mid] {
			return arr[mid]
		}
		if left < mid {
			left = right - 1
		}else {
			right = left + 1
		}
	}
	return -1
}