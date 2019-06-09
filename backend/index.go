package main

import (
	"benji/hackday/backend/model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"strconv"

	"io/ioutil"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()
	const pagesize = 20
	s.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		c.Set("content-type", "application/json")
		c.Next()
	})
	s.GET("/china", func(c *gin.Context) {
		client, err := redis.Dial("tcp", "127.0.0.1:6379")
		defer client.Close()
		if err != nil {
			fmt.Println(err)
		}
		data, err := redis.Bytes(client.Do("GET", "china_data"))
		if err != nil {
			c.String(400, "查询失败，请重试")
			return
		}
		jsonData := &model.ChinaAction{}
		json.Unmarshal(data, jsonData)
		c.JSON(http.StatusOK, gin.H{
			"data": jsonData,
		})
	})
	s.GET("/province", func(c *gin.Context) {
		client, err := redis.Dial("tcp", "127.0.0.1:6379")
		defer client.Close()
		if err != nil {
			fmt.Println(err)
		}
		data, err := redis.Bytes(client.Do("GET", "all_province_data"))
		if err != nil {
			fmt.Println(err)
			c.String(400, "查询失败，请重试")
			return
		}
		jsonData := &model.ProvinceAction{}
		json.Unmarshal(data, jsonData)
		c.JSON(http.StatusOK, gin.H{
			"data": jsonData,
		})
	})
	s.GET("/provinces/change", func(c *gin.Context) {
		client, err := redis.Dial("tcp", "127.0.0.1:6379")
		defer client.Close()
		if err != nil {
			fmt.Println(err)
		}
		data, err := redis.Values(client.Do("LRANGE", "province_hot_list", 0, 32))
		if err != nil {
			fmt.Println(err)
			c.String(400, "查询失败，请重试")
			return
		}
		var resData []*model.SelectChooseListAction
		for _, v := range data {
			temp := &model.SelectChooseListAction{}
			json.Unmarshal(v.([]byte), temp)
			resData = append(resData, temp)
		}
		c.JSON(http.StatusOK, gin.H{
			"data": resData,
		})
	})
	s.GET("/provinces/top", func(c *gin.Context) {
		client, err := redis.Dial("tcp", "127.0.0.1:6379")
		defer client.Close()
		if err != nil {
			fmt.Println(err)
		}
		data, err := redis.Values(client.Do("LRANGE", "province_top_list", 0, 32))
		if err != nil {
			fmt.Println(err)
			c.String(400, "查询失败，请重试")
			return
		}
		var resData []*model.SelectChooseListAction
		for _, v := range data {
			temp := &model.SelectChooseListAction{}
			json.Unmarshal(v.([]byte), temp)
			resData = append(resData, temp)
		}
		c.JSON(http.StatusOK, gin.H{
			"data": resData,
		})
	})
	s.GET("/typeprop", func(c *gin.Context) {
		client, err := redis.Dial("tcp", "127.0.0.1:6379")
		defer client.Close()
		if err != nil {
			fmt.Println(err)
		}
		data, err := redis.Values(client.Do("LRANGE", "type_prop", 0, 32))
		if err != nil {
			fmt.Println(err)
			c.String(400, "查询失败，请重试")
			return
		}
		var resData []*model.TypeProp
		for _, v := range data {
			temp := &model.TypeProp{}
			json.Unmarshal(v.([]byte), temp)
			resData = append(resData, temp)
		}
		c.JSON(http.StatusOK, gin.H{
			"data": resData,
		})
	})
	s.GET("/rank", func(c *gin.Context) {
		client, err := redis.Dial("tcp", "127.0.0.1:6379")
		defer client.Close()
		if err != nil {
			fmt.Println(err)
		}
		data, err := redis.Values(client.Do("LRANGE", "rank", 0, 32))
		if err != nil {
			fmt.Println(err)
			c.String(400, "查询失败，请重试")
			return
		}
		var resData []*model.RankData
		for _, v := range data {

			temp := &model.RankData{}
			fmt.Println(string(v.([]byte)))
			json.Unmarshal(v.([]byte), temp)
			fmt.Println(temp)
			resData = append(resData, temp)
		}
		c.JSON(http.StatusOK, gin.H{
			"data": resData,
		})
	})
	s.GET("/news/:type", func(c *gin.Context) {

		strPage := c.Query("page")
		page, _ := strconv.Atoi(strPage)
		path := "./news/" + c.Param("type") + "/"
		var strData []string
		for i := page * pagesize; i < (page+1)*pagesize; i++ {
			file, err := os.Open(path + strconv.Itoa(i) + ".txt")
			if err != nil {
				fmt.Println(err)
				break
			}
			b, _ := ioutil.ReadAll(file)

			strData = append(strData, string(b))
		}
		if len(strData) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"info": "fail to get data",
			})
			return 
		}
		c.JSON(http.StatusOK, gin.H{
			"news": strData,
		})

	})
	s.GET("/blog",func(c *gin.Context){
		strPage := c.Query("page")
		page,_ := strconv.Atoi(strPage)
		const filePath = "/Users/mac/Documents/project/hackday/backend/original-microblog/"
		var newData []*model.News
		for i:= page*pagesize ;i<(page+1)*pagesize;i++{
			file ,err  := os.Open(filePath+strconv.Itoa(i)+".json")
			if err != nil {
				fmt.Println(err)
				break 
			}
			b,_:= ioutil.ReadAll(file)
			news := &model.News{}
			json.Unmarshal(b,news)
			newData = append(newData,news)	
		}
		if len(newData) == 0{
			c.JSON(http.StatusNotFound,gin.H{
				"info":"blog not found",
			})
			return 
		}
		c.JSON(http.StatusOK,gin.H{
			"blog":newData,
		})
	})
	// dir, _ := ioutil.ReadDir(filePath)
	// for index, file := range dir {
	// 	err := os.Rename(filePath+file.Name(), strconv.Itoa(index)+".json")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	s.Run(":8008")
}
