package main

import "rbmq_new/RabbitMQ"

func main() {
	imoocOne := RabbitMQ.NewRabbitMQRouting("exImooc", "imooc_two")
	imoocOne.RecieveRouting()
}
