package main

import (
	"os"
	"fmt"
)

// OsEnv 系统环境变量获取
type OsEnv struct{
	EnvUser string
	EnvHome string
	Pid  int
}

func main()  {
	fmt.Println("start")
	fmt.Println(GetEnvFunc())
	fmt.Println("end")
}

//GetEnvFunc 获取环境变量函数
func GetEnvFunc() (res OsEnv){
    res.EnvUser =os.Getenv("USER")
	res.EnvHome =os.Getenv("HOME")
	res.Pid =os.Getpid()
	return res
}
