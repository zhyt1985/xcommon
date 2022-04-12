package mq

import (
	"fmt"
	"testing"
	"time"
)
func TestMq(t *testing.T) {

	err := SetupRMQ("amqp://wangxinbo:wangxinbo@172.17.31.185:5672/hotmap") // amqp://用户名:密码@地址:端口号/host
	defer Close()
	if err != nil {
		fmt.Println("err01 : ", err.Error())
	}

	err = Ping()

	if err != nil {
		fmt.Println("err02 : ", err.Error())
	}

	fmt.Println("receive message")

	err = Receive("first", "second", func(msg []byte) {
		fmt.Printf("receve msg is :%s\n", string(msg))
	})

	if err != nil {
		fmt.Println("err04 : ", err.Error())
	}

	fmt.Println("1 - end")

	fmt.Println("send message")

	for i := 0; i < 10; i++ {
		err = Publish("first", "当前时间："+time.Now().String())
		if err != nil {
			fmt.Println("err03 : ", err.Error())
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("2 - end")



}
