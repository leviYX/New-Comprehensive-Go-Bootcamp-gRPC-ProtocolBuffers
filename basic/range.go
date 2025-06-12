package basic

import "fmt"

func main() {
	message := "Hello World"
	// range 用于遍历字符串中的每个字符,并且它操作的是当前结构的副本，而不是原始的字符串。所以你改了啥，都不会影响原始字符串。
	// 也可以使用_来丢弃你不需要的值，以后遍历其余的都可以这样做，这样可以避免内存泄漏，你丢弃的会被gc处理掉
	for i, v := range message {
		fmt.Println(i, string(v))
		fmt.Printf("index: %d, Rune: %c\n", i, v)
	}
}
