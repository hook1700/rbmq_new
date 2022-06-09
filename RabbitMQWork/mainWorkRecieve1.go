package main

import "rbmq_new/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("zwh02")
	rabbitmq.ConsumeSimple()
}
