package main

import (
	"fmt"
	"rbmq_new/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("zwh01")
	rabbitmq.PublishSimple("hello xiao zhang")
	fmt.Println("发送成功！")
}
