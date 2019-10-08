package main

import (
	"fmt"
	"github.com/Zfzlv/ali-green/uuid"
	"github.com/Zfzlv/ali-green/aliyun"
	"encoding/json"
	"time"
)

const accessKeyId string = "xxx"
const accessKeySecret string = "yyyy"


func TestGreenImgAndVideo(){
	//img
	profile := aliyun.Profile{AccessKeyId:accessKeyId, AccessKeySecret:accessKeySecret}
	path := "/green/image/scan"
	clientInfo := aliyun.ClinetInfo{Ip:"127.0.0.1"}
	// 构造请求数据
	bizType := "Green"
	scenes := []string{"porn"}
	//scenes := []string{"porn","terrorism","sface"}
	task := aliyun.Task{DataId:uuid.Rand().Hex(), Url:"https://xxx"}
	//task1 := aliyun.Task{DataId:uuid.Rand().Hex(),Url:"https://yyyy"}
	tasks := []aliyun.Task{task}
	bizData := aliyun.BizData{ bizType, scenes, tasks}
	var client aliyun.IAliYunClient = aliyun.DefaultClient{Profile:profile}
	result := client.GetResponse(path, clientInfo, bizData)
	//fmt.Println(result)
	var r map[string]interface{}
	err := json.Unmarshal([]byte(result), &r)
	//result
	if err == nil{
		code,ok := r["code"]
		if ok && code.(float64) == 200{
			for _,img := range r["data"].([]interface{}){
				c,_ := img.(map[string]interface{})
				if c["code"].(float64) == 200{
					//success
					for _,scenes := range c["results"].([]interface{}){
						d,_ := scenes.(map[string]interface{})
						fmt.Println(d["scene"],d["suggestion"])
					}
				}else{
					//fail
					fmt.Println(time.Now(),"result img parse error code:",c["code"],",msg:",c["msg"])
				}
			}
		}else{
			fmt.Println(time.Now(),"scan img result bad code:",code)
		}
	}else{
	 	fmt.Println(time.Now(),"scan img result parse error:",err.Error())
	}

	//video
	/*path = "/green/video/asyncscan"
	scenes = []string{"porn"}
	task = aliyun.Task{DataId:uuid.Rand().Hex(), Url:"http://xx"}
	tasks = []aliyun.Task{task}
	bizData = aliyun.BizData{ bizType, scenes, tasks}
	result = client.GetResponse(path, clientInfo, bizData)
	fmt.Println(result)*/
	//query result 30s/time
	path = "/green/video/results"
	rqData := aliyun.VideoReq{"kwqi101"}
	result = ""
	num := 0
	for result == "" && num < 480{
		if num > 0{
			time.Sleep(30 * time.Second)
		}
		num++
		result = client.GetResponse(path, clientInfo, rqData)
	}
	var v map[string]interface{}
	err = json.Unmarshal([]byte(result), &v)
	if err == nil{
		code,ok := v["code"]
		if ok && code.(float64) == 200{
			for _,video := range v["data"].([]interface{}){
				c,_ := video.(map[string]interface{})
				if c["code"].(float64) == 200{
					//success
					for _,scenes := range c["results"].([]interface{}){
						d,_ := scenes.(map[string]interface{})
						fmt.Println(d["scene"],d["suggestion"])
					}
				}else{
					//fail
					fmt.Println(time.Now(),"result video parse error code:",c["code"],",msg:",c["msg"])
				}
			}
		}else{
			fmt.Println(time.Now(),"scan video result bad code:",code)
		}
	}else{
	 	fmt.Println(time.Now(),"scan video result parse error:",err.Error())
	}
}	

