package basic

import (
	"fmt"
)

func main() {
	//score := rand.Intn(100)
	//switch {
	//case score > 80:
	//	fmt.Println("优秀")
	//	// fallthrough
	//case score > 60:
	//	fmt.Println("良好")
	//default:
	//	fmt.Println("不及格")
	//}
	checkType(1)
	checkType("hello")
}

/**
 * 类型断言,参数接收一个interface{}类型，然后根据x的实际类型进行断言
 * 之所以interface{}能表示任意类型，然后进行实际判断，因为interface{}类型本身就是一个空接口，不包含任何方法。
 * 空接口可以持有任意类型的值，因为任何类型都实现了至少一个方法。
 */
func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("unknown")
	}
}
