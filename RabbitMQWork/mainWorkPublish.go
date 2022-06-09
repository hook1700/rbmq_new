package main

import (
	"fmt"
	"rbmq_new/RabbitMQ"

	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("zwh02")

	for i := 0; i < 50; i++ {
		rabbitmq.PublishSimple("Hello zhang!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
