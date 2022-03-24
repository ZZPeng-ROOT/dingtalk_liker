package main

import (
    "net/http"	//请求
	"io/ioutil"	//IO
	// "strings"	//字符串
	// "math/rand" //伪随机数
	// "strconv"   //Int转String
	"fmt"		//基本输入输出
)

func main(){
	resp, err := http.Get("https://lv.dingtalk.com/interaction/createLike?uuid=1e780da5-11ad-420e-ad52-12eafa64d918&count=1")
	//处理http错误
	if err != nil {
		fmt.Println("http get error", err)
		return
	}

	//处理解析错误
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	fmt.Println(body)
	fmt.Println("OK")
}