package main

import (
	"fmt"
	"io/ioutil"
	"learn_go/crawler/case1/handle"
	"net/http"
	"regexp"
)

var (
	// \d是数字
	reQQEmail = `(\d+)@qq.com`

	// s?有或者没有s
	// +代表出1次或多次
	//\s\S各种字符
	// +?代表贪婪模式
	reLinke  = `href="(https?://[\s\S]+?)"`
	rePhone  = `1[3456789]\d\s?\d{4}\s?\d{4}`
	reIdcard = `[123456789]\d{5}((19\d{2})|(20[01]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`
	reImg    = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

// 爬取
func main() {
	//caseEmail()
	//caseCallPhone()

}

func caseCallPhone() {
	str := getBodyStr("http://www.zhaohaowang.com/")
	results := getRes(str, rePhone)
	for _, result := range results {
		fmt.Println(result[0])
	}
}

func caseEmail() {
	str := getBodyStr("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	results := getRes(str, reQQEmail)
	if results != nil {
		m := make(map[string]string)
		for _, result := range results {
			email := result[0]
			qq := result[1]
			_, ok := m[email]
			if !ok {
				fmt.Printf("email:[%s] qq:[%s]\n", email, qq)
				m[email] = qq
			}
		}
	}
}

func getRes(bodyStr string, regular string) [][]string {
	if bodyStr == "" || regular == "" {
		return nil
	}
	rp := regexp.MustCompile(regular)
	results := rp.FindAllStringSubmatch(bodyStr, -1)
	return results
}

func getBodyStr(url string) (bodyStr string) {
	resp, err := http.Get(url)
	if err != nil {
		handle.HandleError("request failed. err:", err)
		return ""
	}
	defer resp.Body.Close()
	byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		handle.HandleError("read body failed. err:", err)
		return ""
	}
	str := string(byte)
	return str
}
