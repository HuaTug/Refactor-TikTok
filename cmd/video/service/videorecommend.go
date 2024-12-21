package service

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"HuaTug.com/cmd/video/dal/db"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/videos"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Video struct {
	VideoID     int    `json:"video_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	LabelNames  string `json:"label_names"`
	Category    string `json:"category"`
}

type RecommendVideoService struct {
	ctx context.Context
}

func NewRecommendVideoService(ctx context.Context) *RecommendVideoService {
	return &RecommendVideoService{ctx: ctx}
}

func (service *RecommendVideoService) RecommendVideo(req *videos.RecommendVideoRequest) ([]*base.Video, error) {
	var rev []int64
	var videos []*base.Video
	var mu sync.Mutex

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
		true,              // durable
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

	done := make(chan bool)

	go func() {
		hlog.Info("Start consuming messages from RabbitMQ")
		timeout := time.After(5 * time.Second)
		for {
			select {
			case d := <-msgs:
				{
					var recommendedVideos []Video
					if err := json.Unmarshal(d.Body, &recommendedVideos); err != nil {
						log.Fatalf("Failed to unmarshal JSON: %v", err)
					}
					hlog.Info("Consume messages from RabbitMQ successfully")

					mu.Lock()
					for _, video := range recommendedVideos {
						rev = append(rev, int64(video.VideoID))
						hlog.Infof("Video ID: %d, Title: %s, Description: %s, Label Names: %s, Category: %s",
							video.VideoID, video.Title, video.Description, video.LabelNames, video.Category)
					}
					mu.Unlock()
				}
			case <-timeout:
				{
					hlog.Info("Timeout reached, stopping consumption")
					done <- true
					return
				}
			}
		}
	}()

	<-done

	if len(rev) > 0 {
		var err error
		videos, err = db.GetVideoByVideoId(service.ctx, rev)
		if err != nil {
			return nil, err
		}
	}

	return videos, nil
}
