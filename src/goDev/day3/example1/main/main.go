package main

import(
	"strings"
	"fmt"
)

func httpSupplement(url string) string {
	if (strings.HasPrefix(url, "http://")) {
		return url
	}
	return "http://" + url
}

func main() {
	url1 := "http://www.baidu.com"
	url2 := "www.sina.com"

	fmt.Printf("%s\n", httpSupplement(url1))
	fmt.Printf("%s\n", httpSupplement(url2))
}