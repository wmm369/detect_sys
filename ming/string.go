package ming

import (
	"fmt"
	"strings"
)

//StrFunc 字符串说明
func StrFunc() {
	var testStr string
	testStr = "seetatech"
	testStrTrim := "   seeta   tech    "

	checkStr := "tech"
	fmt.Println("------:" + testStr + ":-----")

	//验证是否包含
	fmt.Println(strings.Contains(testStr, checkStr))

	//计算字符串包含某字符数量
	fmt.Println(strings.Count(testStr, checkStr))

	//计算字符串长度
	fmt.Println(len(testStr))

	// 字符串分割
	splitStr := strings.Split(testStr, "t")
	fmt.Println(splitStr)
	// 字符串连接
	fmt.Println(strings.Join(splitStr, "t"))
	//字符串替换
	fmt.Println(strings.Replace(testStr, "seeta", "firevison", 1))

	//全部小写
	fmt.Println(strings.ToLower(testStr))
	//全部大写
	fmt.Println(strings.ToUpper(testStr))

	//去除左右所要过滤的字符
	fmt.Println(strings.Trim(testStrTrim, " "))
	//去除左所要过滤的字符
	fmt.Println(strings.TrimLeft(testStrTrim, " "))
	//去除右所要过滤的字符
	fmt.Println(strings.TrimRight(testStrTrim, " "))

	//字符出现的位置
	fmt.Println(strings.Index("go gopher", "go"))
	// 最后一次出现的位置
	fmt.Println(strings.LastIndex("go gopher", "go"))
	//没有定位字符时，显示-1
	fmt.Println(strings.LastIndex("go gopher", "rodent"))
}
