/**
 * 爬虫获取拉钩数据
 */
package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	Url  = "http://www.lagou.com/jobs/positionAjax.json?px=default&city={city}&needAddtionalResult=false"
	Post = "first=false&pn={pn}&kd={word}"
)

func main() {
	city := []string{"北京"}      //, "上海", "广州", "深圳", "杭州"}
	language := []string{"php"} //, "ios", "android", "java"}

	for i := 0; i < len(city); i++ {
		for j := 0; j < len(language); j++ {
			url := getUrl(city[i])
			post := getPostData(language[j], j)
			fmt.Println(post)
			httpPost(url, language[j], j)
		}

	}
}

func getUrl(city string) string {
	return strings.Replace(Url, "{city}", city, -1)
}

func getPostData(word string, pageNumber int) string {
	pn := strconv.Itoa(pageNumber)
	str := Post
	str = strings.Replace(str, "{word}", word, -1)
	str = strings.Replace(str, "{pn}", pn, -1)
	return str
}

func httpPost(url string, word string, pageNumber int) {
	pn := strconv.Itoa(pageNumber + 1)
	postData := "first=false&kd=" + word + "&pn=" + pn
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	js, err := simplejson.NewJson(body)

	success, _ := js.Get("success").Bool()
	code, _ := js.Get("code").Int()
	if success == false {
		fmt.Println("json marshal fail success:false")
	} else if code != 0 {
		fmt.Println(code, "json marshal fail code:!0")
	} else {
		/**
		 * 数组array信息获取
		 * interface to string
		 * interface.(string)
		 */
		contents, _ := js.Get("content").Get("positionResult").Get("result").Array()
		for _, info := range contents {
			newInfo, _ := info.(map[string]interface{})
			positionType := newInfo["positionType"].(string)
			positionName := newInfo["positionName"].(string)
			workYear := newInfo["workYear"].(string)
			salary := newInfo["salary"].(string)
			language := word
			city := newInfo["city"].(string)
			fmt.Println(positionType, positionName, workYear, salary, city)
			err := InsertLagouUser(positionType, positionName, workYear, salary, city, language)
			if err != nil {
				fmt.Println("数据插入失败！", err)
			} else {
				fmt.Println("数据插入成功！")
			}
		}
	}
}
