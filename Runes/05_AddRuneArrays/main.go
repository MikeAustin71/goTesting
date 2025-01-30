package main

import "fmt"

func main() {

	nums := make([]rune, 3)
	nums[0] = '9'
	nums[1] = '9'
	nums[2] = '5'

	carry := make([]rune, 1)
	carry[0] = '0'

	fmt.Printf("Starting nums= %v\n",
		string(nums))

	nums[2] = nums[2] + 5

	for i := 2; i > -1; i-- {

		if carry[0] == '1' {
			carry[0] = '0'
			nums[i] = nums[i] + 1
		}

		if nums[i] > '9' {
			nums[i] = '0'

			carry[0] = carry[0] + 1

		}
	}

	if carry[0] == '1' {
		nums = append(carry, nums...)
	}

	fmt.Printf("Ending nums= %v\n",
		string(nums))

}
