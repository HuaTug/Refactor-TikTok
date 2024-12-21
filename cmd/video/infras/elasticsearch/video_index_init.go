package elasticsearch

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

const videoMapping = `{
	"settings": {
	  "number_of_shards": 1,
	  "number_of_replicas": 1
	},
	"mappings": {
	  "properties": {
		"title": {
		  "type": "text",
		  "analyzer": "ik_max_word",
		  "search_analyzer": "ik_smart"
		},
		"description": {
		  "type": "text",
		  "analyzer": "ik_max_word",
		  "search_analyzer": "ik_smart"
		},
		"created_at": {
		  "type": "long"
		},
		"user_id": {
		  "type": "keyword"
		},
		"username": {
		  "type": "keyword"
		},
		"info": {
		  "enabled": false,
		  "properties": {
			"id": {
			  "type": "text"
			},
			"video_url": {
			  "type": "text"
			},
			"cover_url": {
			  "type": "text"
			},
			"visit_count": {
			  "type": "long"
			},
			"like_count": {
			  "type": "long"
			},
			"comment_count": {
			  "type": "long"
			},
			"updated_at": {
			  "type": "long"
			},
			"deleted_at": {
			  "type": "long"
			}
		  }
		}
	  }
	}
  }`

func newVideoIndex() {
	exist, err := elasticClient.IndexExists("video").Do(context.Background())
	if err != nil {
		panic(err)
	}

	if !exist {
		create, err := elasticClient.CreateIndex("video").BodyString(videoMapping).Do(context.Background())
		if err != nil {
			panic(err)
		}

		if create.Acknowledged {
			hlog.Info("Elasticsearch index[video] initialized")
		}
	}
}
