package main

import (
	"bytes"
	"detect_sys/ming"
	"fmt"
	"os"
	"os/exec"
	/* 	"strconv" */
	"strings"
)

// OsEnv 系统环境变量获取
type OsEnv struct {
	EnvUser string
	EnvHome string
	Pid     int
}

func main() {
	fmt.Println("start")
	//fmt.Println(GetEnvFunc())

	//获取服务器配置
	GetEnvSvr()

	//时间time，学习
	ming.TimeNow()

	//字符串操作
	//ming.StrFunc()

	//键盘输入字段
	fmt.Println("输入开始")
	//ming.Input()

	fmt.Println("end")

	//获取yaml文件配置
	ming.GetYaml("./config/first.yaml")
}

//GetEnvFunc 获取环境变量函数
func GetEnvFunc() (res OsEnv) {
	res.EnvUser = os.Getenv("USER")
	res.EnvHome = os.Getenv("HOME")
	res.Pid = os.Getpid()
	return res
}

// GetEnvSvr 1.获取cpu核数
func GetEnvSvr() int {

	cmd := exec.Command("/bin/bash", "-c", "lscpu |grep 'CPU(s)' |grep -v 'li'|grep -v '-'|awk  '{print $1}'")
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("error")
		return -1
	}
	// 去除空格
	str := strings.Replace(out.String(), " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	//number, _ := strconv.Atoi(str)
	fmt.Println("CPU核数：" + str)
	return 1
}
