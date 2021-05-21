package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

// 单任务爬虫
// 创建人: huangjinmu
// 创建时间: 2021/5/21 下午3:02
func main() {
	//resp, err := http.Get("http://www.183fls.com/post/13423.html")
	// class = my_nav
	// /photo_list.aspx?category_id=64 精品
	//url := "http://www.528pcat.com/"
	//all := sendReq(url + "/index.aspx")
	//str := 167055
	for {
		rand.Seed(time.Now().UnixNano())
		str := rand.Int31n(1000000)
		if str < 100000 {
			continue
		}
		url := "http://www.528pcat.com/upload/202105/18/202105180039" + string(str) + ".jpg"
		all := sendReq(url)
		if all != nil {
			fmt.Println(url)
		}

	}

	/*match := `<li><a href="([/a-zA-Z_?.]+=[1-9.]+)">(.+)</a></li>`
	uri := printTab(match, all)
	fmt.Println(uri)
	for _, str := range uri {
		if str == ""{
			continue
		}
		newUrl := url + str
		println("当前爬取的页面地址 = " + newUrl)
		all := sendReq(newUrl)
		//fmt.Printf("%s\n", all)
		//http://www.528pcat.com/photo_show.aspx?id=6175
		match := `href="([/a-zA-Z_?.]+=[1-9.]+)"\s+class="preview-video-container"`
		uri := printTab(match, all)
		//fmt.Printf("%s\n", uri)
		for _, str := range uri {
			if str == ""{
				continue
			}
			newUrl = url + str
			println("当前爬取的页面地址 = " + newUrl)

			all := sendReq(newUrl)
			fmt.Printf("%s\n", all)

			break
		}

		break
	}*/
	/*allString := re.FindAllStringSubmatch(text, -1)
	for _,m := range allString {
		fmt.Println(m)
	}*/
}

func sendReq(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // 延迟执行,关闭连接

	// 获得页面内容
	if resp.StatusCode != http.StatusOK {
		//fmt.Println("Error : status code ", resp.StatusCode)
		return nil
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return all
}

func printTab(match string, all []byte) []string {
	re := regexp.MustCompile(match)
	findAll := re.FindAllSubmatch(all, -1)

	host := make([]string, 1)
	for _, m := range findAll {
		host = append(host, string(m[1]))
	}
	return host
}
