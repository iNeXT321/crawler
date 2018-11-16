package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//给一个链接， 从网上那下一个text
func Fetch(url string) ([]byte, error) {
	//开一个客户端,模仿浏览器访问，解决403 Forbidden问题
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	return  ioutil.ReadAll(resp.Body)

}
