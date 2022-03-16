package main

import (
    "net/http"	//请求
	"time"		//计时器
	"io/ioutil"	//IO
	"strings"	//字符串
	"math/rand" //伪随机数
	"strconv"   //Int转String
	"fmt"		//基本输入输出
)

var uuid string
var legitMode,debug bool
var requestTimes,waitTime,likes,tmp,i int

func initSettings(){
/*
	legitMode: 合法模式开关
	requestTimes: 请求计数器
	waitTime: 等待时间
	debug: DEBUG模式开关
*/
	legitMode = false
	requestTimes = 10000
	waitTime = 10
	debug = false
	tmp=100
}

func motd(){
	fmt.Println("dingtalk like helper v0.1 by ZZPeng")
	fmt.Println("https://blog.zzpeng.com")
}

func getInfo(){
	fmt.Print("UUID:")
	fmt.Scan(&uuid)
}

func main(){
	motd()
	initSettings()
	getInfo()
	fmt.Print("Testing Network...")
	resp, err := http.Get("https://ipinfo.io")
	if err != nil {
		fmt.Println("network error", err)
		return
	}
	fmt.Print("OK\nStarting thread...")
	time.Sleep(2*time.Second)
	fmt.Println("Working")
	rand.Seed(time.Now().UnixNano())
	fmt.Println("URL: https://lv.dingtalk.com/interaction/createLike?uuid="+uuid+"&count="+strconv.Itoa(tmp))
	for {
		//模式判断(暴力,合法)
		if legitMode {
			//合法模式,随机请求数
			tmp = rand.Intn(100)
		}
		i ++
		//请求
		resp, err = http.Get("https://lv.dingtalk.com/interaction/createLike?uuid="+uuid+"&count="+strconv.Itoa(tmp))
		resp.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
		//处理http错误
		if err != nil {
			fmt.Println("http get error", err)
			break
		}

		//处理解析错误
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("read error", err)
			break
		}
		
		//检查是否API调用错误
		if strings.Index(string(body),"punish") >0  {
			fmt.Println("API error")
			break
		}
		likes += tmp
		fmt.Println("[Info] 第 "+strconv.Itoa(i)+" 次点赞,本次点赞: "+strconv.Itoa(tmp)+" ,累计点赞: "+strconv.Itoa(likes)+",下次点赞还需等待"+strconv.Itoa(waitTime)+"秒")
		//DEBUG模式下输出响应信息
		if debug {
			fmt.Println(string(body))
			fmt.Println(resp)
		}

		if i ==requestTimes {break} //计数器
		time.Sleep(time.Duration(waitTime)*time.Second)
	}
}