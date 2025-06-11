package basic

import (
	"fmt"
	myhttp "net/http"
)

const (
	baiduURL  = "https://www.baidu.com"
	googleURL = "https://www.google.com"
)

func main1() {
	resp, err := myhttp.Get(baiduURL)
	if err != nil {
		fmt.Println("http get error:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
}
