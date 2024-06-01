package recursion

import "fmt"

func Recursion(num int) int {

	// base case
	if num == 1 {
		return num
	}
	// pre-recursive case
	log := num + Recursion(num-1)
	// post-recursive case
	fmt.Println(log)
	return log
}
