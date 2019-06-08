package main

import (
	"benji/hackday/backend/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()

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

	s.Run(":8008")
}
