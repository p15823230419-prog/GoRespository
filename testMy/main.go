package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"time"
)

// 访问的网站地址
var url = "http://jszddj.top:8080"

// 定义根结构体
type MessageList struct {
	Messages []struct {
		ID        int       `json:"id"`
		Sender    string    `json:"sender"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"消息列表"`
}

type sendUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type reqUser struct {
	Message bool `json:"message"`
	User    struct {
		Username string `json:"username"`
	} `json:"user"`
}

func Login(username string, password string) (isLogin bool) {
	data := sendUser{
		Username: username,
		Password: password,
	}
	jsonBytes, _ := json.Marshal(data)
	resp, err := http.Post(url+"/login", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	req := reqUser{}
	errr := json.Unmarshal(body, &req)
	if errr != nil {
		fmt.Println(errr)
		return false
	}
	fmt.Println(string(body))
	fmt.Println(req)
	return req.Message
}

// 发送消息函数
func sendMessage(sender, receiver, content string) {
	data := map[string]string{
		"sender":   sender,
		"receiver": receiver,
		"content":  content,
	}
	jsonBytes, _ := json.Marshal(data)
	resp, err := http.Post(url+"/send", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	//fmt.Println("发送消息成功: 状态码:", resp.StatusCode)
}

// 接受信息函数
func receiveMessage(receiver string) (messages MessageList) {
	resp, err := http.Get(url + "/messages?user=" + receiver + "&after_id=0")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	messageList := MessageList{}
	err = json.Unmarshal(body, &messageList)
	//for _, message := range messageList.Messages {
	//	fmt.Printf("%-v\t\t\t\t\t\t\t%v\t%v\n", message.Content, message.Sender, message.CreatedAt)
	//}
	return messageList
}

func main() {
	//start()
	//sendMessage("zhangsan", "lisi", "你好用户lisi22")
	//Login("zhangsan", "123456")
	//获取消息
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	start()

}

// 监听消息
func ListenNewMessage(receiver string) {

	var lastMessageTime int64
	messages := receiveMessage(receiver).Messages
	for _, message := range messages {
		fmt.Println(message.Sender, message.Content, message.CreatedAt)
	}
	lastMessageTime = messages[len(messages)-1].CreatedAt.Unix()
	for {
		messages := receiveMessage(receiver).Messages
		lastMessage := messages[len(messages)-1]
		if lastMessage.CreatedAt.Unix() > lastMessageTime {
			fmt.Println("收到新消息")
			lastMessageTime = messages[len(messages)-1].CreatedAt.Unix()
			fmt.Println(lastMessage.Sender, lastMessage.Content, lastMessage.CreatedAt)
		}
		runtime.Gosched()
	}
}
