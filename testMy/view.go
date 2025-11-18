package main

import (
	"fmt"
	"runtime"
)

func start() {
	username := ""
	password := ""
	for !Login(username, password) {
		fmt.Println("欢迎使用聊天系统")
		fmt.Println("请先登录")
		fmt.Println("请输入用户名")
		fmt.Scanln(&username)
		fmt.Println("请输入密码")
		fmt.Scanln(&password)
	}
	fmt.Printf("登录成功!\n欢迎迎用户%s\n", username)
	go ListenNewMessage(username)
	option := -1
	for option != 0 {
		fmt.Println("1:发送消息\n2:接收消息\n3:退出")
		fmt.Println("请输入操作数")
		fmt.Scanln(&option)
		switch option {
		case 1:
			receive, content := "", ""
			fmt.Println("请输入要消息的接收者")
			fmt.Scanln(&receive)
			fmt.Println("请输入要发送的消息")
			fmt.Scanln(&content)
			sendMessage(username, receive, content)
		case 2:
			receiveMessage(username)
		}
		runtime.Gosched()
	}

}
