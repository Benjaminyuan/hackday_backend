package main

import (
	"benji/hackday/crawl"
	"fmt"

	// "sync"
	"encoding/json"

	"github.com/garyburd/redigo/redis"
)

const ALLCHINA = "http://www.wrd.cn/view/home/hotEvent/allChina.action"
const ALlCHINADATA = "timeType=1&sort=2&areaType=1"

const ALLPROVINCE = "http://www.wrd.cn/view/home/hotEvent/allProvince.action"
const ALLPROVINCEDATA = "timeType=1"

const SELECTCHOOSELIST = "http://www.wrd.cn/view/home/hotEvent/selectChooseListData.action"
const SELECTCHOOSELISTDATA = "timeType=1&sort=5&labels=&areaType=1&province=%E5%90%89%E6%9E%97"

const TYPEPROP = "http://www.wrd.cn/view/home/hotEvent/typeProp.action"
const TYPEPROPDATA = "timeType=1&province=%E5%90%89%E6%9E%97"

const SELECTRANKDATA = "http://www.wrd.cn/view/home/hotEvent/selectRankData.action"
const SELECTRANKDATADATA = "timeType=1&labels=&province=%E9%BB%91%E9%BE%99%E6%B1%9F"

var provinces = []string{"北京", "天津", "上海", "重庆", "河北", "山西省", "辽宁", "吉林", "黑龙江",
	"江苏", "浙江", "安徽", "福建", "江西", "山东", "河南", "湖北", "湖南", "广东",
	"海南", "四川", "贵州", "云南", "陕西", "甘肃", "青海", "台湾", "西藏", "新疆", "内蒙古", "宁夏", "香港", "澳门"}

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Print("fail to connect redis ")
	}
	var wait_group sync.WaitGroup
	lock := make(chan int)
	defer c.Close()
	// allChina := crawl.GetMapData(ALLCHINA, ALlCHINADATA)
	// allChinaData := crawl.ParseAllChinaData(allChina)
	// fmt.Println(allChinaData)
	// temp, err := json.Marshal(allChinaData)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// c.Do("SET", "china_data", string(temp))
	// fmt.Print(allChinaData)
	// fmt.Print("------------allchina--------------\n\n\n")
	/*--------------*/
	wait_group.Add(1)
	go func(){
		defer wait_group.Done()
		allChina := crawl.GetMapData(ALLCHINA, ALlCHINADATA)
		allChinaData := crawl.ParseAllChinaData(allChina)
		temp , err := json.Marshal(allChinaData)
		if err != nil {
			fmt.Println(err)
		}
		c.Do("SET","china_data",string(temp))
		fmt.Print(allChinaData)
		fmt.Print("------------allchina--------------\n\n\n")
	}()

	/*--------------*/

	// allProvince := crawl.GetMapData(ALLPROVINCE, ALLPROVINCEDATA)
	// allProvinceData := crawl.ParseAllProvinceData(allProvince)
	// fmt.Println(allProvinceData)
	// provinceTemp, _ := json.Marshal(allProvinceData)
	// c.Do("SET", "all_province_data", string(provinceTemp))
	// fmt.Print("---------------province-------------------\n\n\n")
	wait_group.Add(1)
	go func(){
		defer wait_group.Done()
		allProvince := crawl.GetMapData(ALLPROVINCE, ALLPROVINCEDATA)
		allProvinceData := crawl.ParseAllProvinceData(allProvince)
		fmt.Println(allProvinceData)
		temp ,_:= json.Marshal(allProvinceData)
		fmt.Println(temp)
		c.Do("SET","all_province_data",string(temp))
		fmt.Print("---------------province-------------------\n\n\n")
	}()

	/*--------------*/

	// for _, province := range provinces {
	// 	hot := crawl.GetMapData(SELECTCHOOSELIST, crawl.GetSelectChoose(province, "2"))
	// 	top := crawl.GetMapData(SELECTCHOOSELIST, crawl.GetSelectChoose(province, "5"))
	// 	typeProp := crawl.GetMapData(TYPEPROP, crawl.GetTypeProp(province))
	// 	Rank := crawl.GetMapData(SELECTRANKDATA, crawl.GetRank(province))

	// 	hotData := crawl.ParseSelectChoseList(hot)

	// 	topData := crawl.ParseSelectChoseList(top)
	// 	typePropData := crawl.ParseTypeProp(typeProp)
	// 	RankData := crawl.ParseRankData(Rank)

	// 	fmt.Println("hot:  ", hotData)
	// 	fmt.Print("---------------\n\n\n")
	// 	fmt.Println("hot:   ", typePropData)
	// 	fmt.Print("---------------\n\n\n")
	// 	fmt.Println(string(Rank))
	// 	fmt.Println("rank:  ", RankData)
	// 	fmt.Print("---------------\n\n\n")
	// 	hotTemp, _ := json.Marshal(hotData)
	// 	topTemp, _ := json.Marshal(topData)
	// 	typeTemp, _ := json.Marshal(typePropData)
	// 	rankTemp, _ := json.Marshal(RankData)

	// 	c.Do("LPUSH", "province_hot_list", string(hotTemp))
	// 	c.Do("LPUSH", "province_top_list", string(topTemp))
	// 	c.Do("LPUSH", "type_prop", string(typeTemp))
	// 	c.Do("LPUSH", "rank", string(rankTemp))

	// }
	wait_group.Add(1)
	go func(){
		defer wait_group.Done()
		for _, province := range provinces{
			hot:= crawl.GetMapData(SELECTCHOOSELIST, crawl.GetSelectChoose(province,"2"))
			top := crawl.GetMapData(SELECTCHOOSELIST,crawl.GetSelectChoose(province,"5"))
			typeProp := crawl.GetMapData(TYPEPROP,crawl.GetTypeProp(province))
			Rank := crawl.GetMapData(SELECTRANKDATA, crawl.GetRank(province))

			hotData := crawl.ParseSelectChoseList(hot)
			topData := crawl.ParseSelectChoseList(top)
			typePropData := crawl.ParseTypeProp(typeProp)
			RankData := crawl.ParseRankData(Rank)

			hotTemp, _ := json.Marshal(hotData)
			topTemp,_:= json.Marshal(topData)
			typeTemp,_:= json.Marshal(typePropData)
			rankTemp,_ :=json.Marshal(RankData)

			lock <- 1
			c.Do("LPUSH","province_hot_list",string(hotTemp))
			c.Do("LPUSH","province_top_list", string(topTemp))
			c.Do("LPUSH","type_prop",string(typeTemp))
			c.Do("LPUSH","rank",string(rankTemp))
			<-lock
		}

	}()

	wait_group.Wait()

}
