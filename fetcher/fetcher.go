package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/text/encoding/unicode"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"

	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("wrong http request: %s", err.Error())
		return nil, fmt.Errorf("wrong http request: %s", err.Error())
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status code %d", resp.StatusCode)
	}

	newReader := bufio.NewReader(resp.Body)
	e := determinEncoding(newReader)
	reader := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(reader)

}

// 获取编码 解决gbk 等其他charset的问题 返回&{UTF-8}等
func determinEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, e := r.Peek(1024)
	if e != nil {
		log.Printf("Fetcher %s", e)
		return unicode.UTF8
	}
	e2, _, _ := charset.DetermineEncoding(bytes, "")
	return e2
}
