/*
产生当前时刻 精确到分钟，外加产生4位随机数
*/
package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main(){
	fmt.Println("产生当前时刻 精确到分钟，外加产生4位随机数")
	var t int64 = time.Now().Unix()
	var s string = time.Unix(t, 0).Format("20060102150405")
	fmt.Println(s)
	
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))  
    vcode := fmt.Sprintf("%04v", rnd.Int31n(10000))  
    fmt.Println(vcode)  
	fmt.Println(s+vcode)
}