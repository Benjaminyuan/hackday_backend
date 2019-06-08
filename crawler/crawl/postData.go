package crawl 
import (
	"net/url"
)
func GetTypeProp(province string) string {
	data := &url.Values{}
	data.Add("timeType","1")
	data.Add("province",province)
	return data.Encode()
}
func GetSelectChoose(province string,sort string) string {
	data:= &url.Values{}
	data.Add("timeType","1")
	/*
	  sort = 2 -> hot top 
   	  sort = 5 -> change top 
   */
	data.Add("sort",sort)
	data.Add("labels","")
	data.Add("areaType","1")
	data.Add("province",province)
	return data.Encode()
}
func GetRank(province string) string {
	data:= &url.Values{}
	data.Add("timeType","1")
	data.Add("labels","")
	data.Add("province",province)
	return data.Encode()
}