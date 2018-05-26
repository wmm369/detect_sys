package ming

import (
	"fmt"
	"time"
)

//TimeNow 时间函数
func TimeNow() {
	//获取毫秒
	fmt.Println(time.Microsecond)

	//获取月
	fmt.Println(time.Month(1))

	//当前时间
	fmt.Println(time.Now())
	fmt.Println(time.Now().String())
	//当前时间-小时
	fmt.Println(time.Now().Hour())

	//当前时间unix时间戳since 1970 -1- 1
	fmt.Println(time.Now().Unix())

	//当前时间unix时间戳(nanoseconds),since 1970 -1- 1,
	fmt.Println(time.Now().UnixNano())

	//当前时间加三个小时
	fmt.Println(time.Now().Add(time.Hour * 3))

	//时间戳转化成时间
	currentTime := time.Now().Unix()
	tm := time.Unix(currentTime, 0)
	fmt.Println(tm)
}
