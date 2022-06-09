package main

import "rbmq_new/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("zwh01")
	rabbitmq.ConsumeSimple()
}