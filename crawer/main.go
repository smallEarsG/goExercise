package main

import (
	"fmt"
	"regexp"
)

func main() {
	//regexp.Compile("res")
	//resp, err := http.Get("http://localhost:8080/mock/www.zhenai.com/zhenghun")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//if resp.StatusCode == http.StatusOK {
	//	//如果中文乱码
	//	e := determineEncodeing(resp.Body)
	//	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	//	all, err := ioutil.ReadAll(utf8Reader)
	//	if err != nil {
	//		panic(err)
	//	}
	//	//fmt.Printf("%s\n", all)
	//	printCityList(all)
	//} else {
	//	fmt.Println("Error:status code", resp.StatusCode)
	//}

}

// []byte 就相当于一个string
func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"+[^>]*>(.*[^<])</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City:%s ,URL:%s\n", m[2], m[1])
		//for _, subMatch := range m {
		//	fmt.Printf("City:%s ,URL:%s\n", subMatch)
		//}
		//fmt.Println()
	}
	//return matches
}
