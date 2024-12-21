package main

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Video represents a video in the recommendation result.
type Video struct {
	VideoID     int    `json:"video_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	LabelNames  string `json:"label_names"`
	Category    string `json:"category"`
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"recommend_queue", // name
		false,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var recommendedVideos []Video
			err := json.Unmarshal(d.Body, &recommendedVideos)
			if err != nil {
				log.Fatalf("Failed to unmarshal JSON: %v", err)
			}

			// 处理推荐结果
			for _, video := range recommendedVideos {
				fmt.Printf("Video ID: %d, Title: %s, Description: %s, Label Names: %s, Category: %s\n",
					video.VideoID, video.Title, video.Description, video.LabelNames, video.Category)
			}
		}
	}()

	fmt.Println("Waiting for messages. To exit press CTRL+C")
	<-forever
}
