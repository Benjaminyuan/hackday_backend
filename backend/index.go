package main
import (
	"github.com/gin-gonic/gin"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"net/http"
)
func main(){
	s :=gin.Default()
	client,err:=redis.Dial("tcp","127.0.0.1:6379")
	defer client.Close()
	if err!= nil{
		fmt.Println(err)
	}
	s.GET("/china",func(c *gin.Context){
		data , err := redis.Bytes(client.Do("GET","china_data"))
		if err != nil {
			c.String(400,"查询失败，请重试")
			return 
		}
		c.String(http.StatusOK,string(data))
	})
	s.GET("/province",func(c *gin.Context){
		data, err := redis.Bytes(client.Do("GET","all_province_data"))
		if err!= nil {
			fmt.Println(err)
			c.String(400,"查询失败，请重试")
			return 
		}
		c.String(http.StatusOK,string(data))
	})
	s.GET("/provinces/hot",func(c *gin.Context){
		data , err := redis.Values(client.Do("LRANGE","province_hot_list",0,32))
		if err != nil {
			fmt.Println(err)
			c.String(400,"查询失败，请重试")
			return 
		}	
		var resData []string
		for _, v := range data {
			resData = append(resData,string(v.([]byte)))
		}
		c.JSON(http.StatusOK,gin.H{
			"data":resData,
		})
	})
	s.GET("/provinces/top",func(c *gin.Context){
		data , err := redis.Values(client.Do("LRANGE","province_top_list",0,32))
		if err != nil {
			fmt.Println(err)
			c.String(400,"查询失败，请重试")
			return 
		}
		var resData []string
		for _, v := range data {
			resData = append(resData,string(v.([]byte)))
		}
		c.JSON(http.StatusOK,gin.H{
			"data":resData,
		})
	})
	s.GET("/typeprop",func(c *gin.Context){
		data , err := redis.Values(client.Do("LRANGE","type_prop",0,32))
		if err != nil {
			fmt.Println(err)
			c.String(400,"查询失败，请重试")
			return
		}
		var resData []string
		for _, v := range data {
			resData = append(resData,string(v.([]byte)))
		}
		c.JSON(http.StatusOK,gin.H{
			"data":resData,
		})
	})
	s.GET("/rank",func(c *gin.Context){
		data , err := redis.Values(client.Do("LRANGE","rank",0,32))
		if err != nil {
			fmt.Println(err)
			c.String(400,"查询失败，请重试")
			return 
		}
		var resData []string
		for _, v := range data {
			resData = append(resData,string(v.([]byte)))
		}
		c.JSON(http.StatusOK,gin.H{
			"data":resData,
		})
	})
	s.Run(":8008")
}
