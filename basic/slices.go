package basic

import (
	"fmt"
	"slices"
)

func main() {
	//var nums1 []int
	//var nums2 = []int{1, 2, 3}
	//nums3 := []int{4, 5, 6}
	//slice := make([]int, 5)

	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:4]
	fmt.Println(slice1)
	slice1 = append(slice1, 6, 7, 8)
	fmt.Println(slice1)
	fmt.Println("*****************************")

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("sliceCopy is:", sliceCopy)

	// 遍历切片
	for i, v := range sliceCopy {
		fmt.Println(i, v)
	}

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slices are equal")
	}

}
