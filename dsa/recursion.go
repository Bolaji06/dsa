package dsa

import "fmt"

// this file contain some list of recursion examples
func factorial(n int) int {
	// base case
	if (n < 2){
		return 1
	}else {
		return factorial(n - 1) + factorial(n - 2) 
	}
}
func Recursion(){
	fact := factorial(5);
	fmt.Println(fact);

}