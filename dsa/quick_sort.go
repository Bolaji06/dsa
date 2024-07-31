
package dsa


func QuickSort(arr[]int) []int {

	if len(arr) < 2 {
		return arr;
	}
	// select the pivot
	pivot := arr[0];

	// storages
	var left, middle, right [] int

	for i := 0; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else if arr[i] == pivot {
			middle = append(middle, arr[i])
		}else {
			right = append(right, arr[i])
		}
	}
	return append(append(QuickSort(left), middle...), QuickSort(right)...)

}