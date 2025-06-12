package basic

import "fmt"

func main() {
	var numbers [5]int
	fmt.Println(numbers)
	numbers[0] = 1
	fmt.Println(numbers)

	var strs = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(strs)

	// 数组赋值给另一个数组，其实是拷贝了一份新的数组，两者互不影响
	var originArray = [5]int{1, 2, 3, 4, 5}
	fmt.Println(originArray)
	var newArray = originArray
	fmt.Println(newArray)

	newArray[0] = 10
	fmt.Println(newArray)
	fmt.Println(originArray)

	// 遍历数组1
	for i := 0; i < len(originArray); i++ {
		fmt.Println("index is :", i, " value is :", originArray[i])
	}
	// 遍历数组2
	for index, value := range originArray {
		fmt.Println("index is :", index, " value is :", value)
	}
}
