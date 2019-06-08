package main

import (
	"benji/hackday/crawl"
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

func main() {
	crawl.GetMapData(SELECTRANKDATA, SELECTRANKDATADATA)
}
