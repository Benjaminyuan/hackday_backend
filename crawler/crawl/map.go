package crawl

import (
	_"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/axgle/mahonia"
)

const url_all = "http://www.wrd.cn/view/home/hotEvent/allProvince.action"
const curl = `/*
curl 'http://www.wrd.cn/view/home/hotEvent/allProvince.action' 
-H 'Cookie: JSESSIONID=96273CC51FCEA2D0AAF7AF0CC12F6466' 
-H 'Origin: http://www.wrd.cn' 
-H 'Accept-Encoding: gzip, deflate'
 -H 'Accept-Language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7' 
 -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36' 
 -H 'Content-Type: application/x-www-form-urlencoded' 
 -H 'Accept: application/json, text/plain, */*' 
 -H 'Referer: http://www.wrd.cn/login.shtml' 
 -H 'X-Requested-With: XMLHttpRequest' 
 -H 'Connection: keep-alive' 
 --data 'timeType=1' 
 --compressed 
*/`
const curl_AllChina = `
curl 'http://www.wrd.cn/view/home/hotEvent/allChina.action'
 -H 'Accept: application/json, text/plain, */*' 
 -H 'Referer: http://www.wrd.cn/login.shtml' 
 -H 'Origin: http://www.wrd.cn' 
 -H 'X-Requested-With: XMLHttpRequest' 
 -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36' 
 -H 'Content-Type: application/x-www-form-urlencoded'
  --data 'timeType=1&sort=2&areaType=1' 
  --compressed`
const selectChooseList = `
curl 'http://www.wrd.cn/view/home/hotEvent/selectChooseListData.action' 
-H 'Accept: application/json, text/plain, */*'
 -H 'Referer: http://www.wrd.cn/login.shtml' 
 -H 'Origin: http://www.wrd.cn' 
 -H 'X-Requested-With: XMLHttpRequest' 
 -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36' 
 -H 'Content-Type: application/x-www-form-urlencoded' 
 --data 
 'timeType=1&sort=5&labels=&province=%E5%90%89%E6%9E%97&areaType=1' --compressed
`

const typeProp = `
curl 'http://www.wrd.cn/view/home/hotEvent/typeProp.action' 
-H 'Accept: application/json, text/plain, */*' 
-H 'Referer: http://www.wrd.cn/login.shtml' 
-H 'Origin: http://www.wrd.cn' 
-H 'X-Requested-With: XMLHttpRequest' 
-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36' 
-H 'Content-Type: application/x-www-form-urlencoded' 
--data 'timeType=1&province=%E5%90%89%E6%9E%97' 
--compressed`
const selectRankData = `
curl 'http://www.wrd.cn/view/home/hotEvent/selectRankData.action'
 -H 'Accept: application/json, text/plain, */*' 
 -H 'Referer: http://www.wrd.cn/login.shtml' 
 -H 'Origin: http://www.wrd.cn' 
 -H 'X-Requested-With: XMLHttpRequest' 
 -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36' 
 -H 'Content-Type: application/x-www-form-urlencoded' 
 --data 'timeType=1&labels=&province=%E9%BB%91%E9%BE%99%E6%B1%9F' 
 --compressed`

func GetMapData(url string, postdata string) []byte {
	client := &http.Client{}
	// request, err := http.NewRequest("POST", url, strings.NewReader("timeType=1"))
	request, err := http.NewRequest("POST", url, strings.NewReader(postdata))
	if err != nil {
		log.Fatal("fail to create a request s")
	}
	request.Header.Add("Origin", "http://www.wrd.cn")
	request.Header.Add("Accept-Encoding", "gzip, deflate")
	request.Header.Add("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Accept", "application/json, text/plain, */*")
	request.Header.Add("Referer", "http://www.wrd.cn/login.shtml")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	// request.Header.Add("Connection", "keep-alive")
	// request.Header.Add("Cokie", "JSESSIONID=96273CC51FCEA2D0AAF7AF0CC12F646")

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
		log.Fatal("fail to do request ")
	}
	defer response.Body.Close()

	dec := mahonia.NewDecoder("gbk")
	rd := dec.NewReader(response.Body)
	c, _ := ioutil.ReadAll(rd)
	// fmt.Println(string(c))
	return c

}
