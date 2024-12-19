package main

import (
	"fmt"
	"slices"
)

func main(){
	//dynamic
	
	//uninitialzed slice is nil
	// var nums []int
	// fmt.Println(nums,len(nums),cap(nums),nums==nil)

	//nums := []int{}//short handle

	//we can use index also to append elements like array

	// var nums = make([]int,2,9)//type,length,capacity
	// fmt.Println(nums,len(nums),cap(nums),nums==nil)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 5)
	// fmt.Println(nums,len(nums),cap(nums),nums==nil)//if the number os elements exceeds the capacity it automatically doubles the capacity{if odd then it doubles its succesor even number as capacity}
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 5)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 5)
	// fmt.Println(nums,len(nums),cap(nums),nums==nil)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 5)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 5)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 5)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 5)
	// fmt.Println(nums,len(nums),cap(nums),nums==nil)

	//copy function
	var nums = make([]int,2,5)
	nums= append(nums, 2)//len=3 now
	var nums2 = make([]int,len(nums))
	copy(nums2,nums)//destiny,source
	fmt.Println(nums,nums2)

	//slice operator
	fmt.Println(nums[0:2])//0 to less than 2
	fmt.Println(nums[:])//all elements//slicing is similar to python only

	//slice methods
	fmt.Println(slices.Equal(nums,nums2))//returns boolean


}
