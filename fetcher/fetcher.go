package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func Fecth(url string) ([]byte, error) {
	regexp.Compile("res")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		//如果中文乱码
		e := determineEncoding(resp.Body)
		utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
		return ioutil.ReadAll(utf8Reader)
	} else {
		//fmt.Println("Error:status code", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code %d", resp.StatusCode)
	}
}
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
