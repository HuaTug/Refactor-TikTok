package es

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/olivere/elastic/v7"
)

type VideoOtherData struct {
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	//PublishTime   int64 `json:"publish_time"`
	Title string `json:"title"`
}
type Video struct {
	VideoId  int64          `json:"video_id'`
	AuthorId int64          `json:"author_id"`
	Info     VideoOtherData `json:"info`
}

const indexs = "videos"

const videoMapping = `{
	"mappings":{
	  "properties":{
		"video_id":{
			"type":"long"
		},
		"author_id":{
			"type":"long"
		},
		"info":{
			"properties":{
			"play_url":{
				"type":"text"
			},
			"cover_url":{
				"type":"text"
			},
			"favorite_count":{
				"type":"long"
			},
			"comment_count":{
				"type":"long"
			},
			"title":{
				"type":"text"
			}	
			}
		}
	  }
	}
}`

type VideoIndex struct {
	Index   string
	Mapping string
}

func NewVideoIndex() (*VideoIndex, error) {
	video := VideoIndex{
		Index:   indexs,
		Mapping: videoMapping,
	}
	if err := video.Create(); err != nil {
		hlog.Info("Fail to Create!")
		return nil, err
	} else {
		return &video, nil
	}
}

func (v *VideoIndex) Create() error {
	ctx := context.Background()

	exist, err := Client.IndexExists(indexs).Do(ctx)
	if err != nil {
		hlog.Info(err)
		return err
	}
	if !exist {
		create, err := Client.CreateIndex(indexs).Body(v.Mapping).Do(ctx)
		if err != nil {
			hlog.Info(err)
			return err
		}

		if create.Acknowledged {
			hlog.Info("Elasticsearch index[video] initialized")
		}
	}
	return nil
}

func (v *VideoIndex) CreateVideoDoc(data Video) error {
	video_id := strconv.FormatInt(int64(data.VideoId), 10)
	_, err := Client.Index().Index(v.Index).Id(video_id).BodyJson(data).Do(context.Background())
	if err != nil {
		hlog.Info("Create doc failed: ", err)
		return err
	}
	return nil
}

func (v *VideoIndex) DeleteVideoDoc(vid int64) error {
	videoId := strconv.FormatInt(int64(vid), 10)
	_, err := Client.Delete().Index(v.Index).Id(videoId).Do(context.Background())
	if err != nil {
		hlog.Info("Delete doc failed: ", err)
		return err
	}
	return nil
}

func (v *VideoIndex) SearchVideoDoc(vid int64) (*Video, error) {
	videoId := strconv.FormatInt(int64(vid), 10)
	resp, err := Client.Get().Index(v.Index).Id(videoId).Do(context.Background())
	if err != nil {
		hlog.Info("Search doc failed: ", err)
		return nil, err
	}
	if resp.Found {
		var video Video
		json.Unmarshal(resp.Source, &video)
		return &video, nil
	} else {
		return nil, errors.New("Data not found")
	}
}

func (v *VideoIndex) UpdateVideoDoc(video Video) error {
	videoId := strconv.FormatInt(int64(video.VideoId), 10)
	hlog.Info(videoId)
	_, err := Client.Update().Index(v.Index).Id(videoId).Doc(video).Do(context.Background())
	if err != nil {
		hlog.Info("Update doc failed: ", err)
		return err
	}
	return nil
}

func (v *VideoIndex) searchRespCovert(resp *elastic.SearchResult) ([]*Video, int64) {
	dataList := make([]*Video, 0)
	for _, item := range resp.Each(reflect.TypeOf(Video{})) {
		data, ok := item.(Video)
		if !ok {
			continue
		}
		temp := Video{
			VideoId:  data.VideoId,
			AuthorId: data.AuthorId,
			Info:     data.Info,
		}
		dataList = append(dataList, &temp)
	}
	hits := resp.TotalHits()
	return dataList, hits
}

func (v *VideoIndex) SearchVideoDocDefault(keyword string) ([]*Video, int64, error) {
	resp, err := Client.Search().Index(v.Index).
		//同时需要注意复合字段的匹配规则 形式也有所不同
		//.Should表示只需满足其中的一个查询条件即可
		Query(elastic.NewBoolQuery().Should(
			elastic.NewMultiMatchQuery(keyword, "Info.title"),
			//使用通配符进行字段数据匹配
			elastic.NewWildcardQuery("Info.title", "*"+keyword+"*"),
		)).
		//Sort("publish_time", true).
		From(0).Size(10).
		Do(context.Background())
	if err != nil {
		return nil, -1, err
	}
	data, hits := v.searchRespCovert(resp)
	return data, hits, nil
}

func (v *VideoIndex) SearchVideoDocs(keywords, fromDate, toDate string, pageNum, pageSize int64) ([]*Video, int64, error) {
	var wg sync.WaitGroup
	var (
		shouldQuery = make([]elastic.Query, 0)
		filterQuery = make([]elastic.Query, 0)
	)
	wg.Add(3)
	go func() {
		if keywords != "" {
			shouldQuery = append(shouldQuery, elastic.NewWildcardQuery("Info.title", "*"+keywords+"*"))
		}
		wg.Done()
	}()
	/* 	go func() {
		if fromDate != "" && toDate != "" {
			filterQuery = append(filterQuery, elastic.NewRangeQuery("info.publish_time").From(fromDate).To(toDate))
		}
		wg.Done()
	}() */
	go func() {
		if pageSize == 0 {
			pageSize = 5
		}
		wg.Done()
	}()
	go func() {
		if pageNum == 0 {
			pageNum = 1
		}
		wg.Done()
	}()
	wg.Wait()
	query := elastic.NewBoolQuery().
		Should(shouldQuery...).
		Filter(filterQuery...)
	resp, err := Client.Search().
		Index(v.Index).Query(query).
		From(int(pageNum-1) * int(pageSize)).Size(int(pageSize)).Do(context.Background())
	if err != nil {
		return nil, -1, err
	}
	data, hits := v.searchRespCovert(resp)
	return data, hits, nil
}

func (v *VideoIndex) UpdateTitle(newId int64, vid string) error {
	_, err := Client.Update().Index(indexs).Type("_doc").Id(vid).
		/*Doc(
				在给定的代码示例中，结构体的字段名称和字段类型（在这种情况下是 Title string）与 Elasticsearch 文档的 Mapping 中的字段名称和字段类型应该是一致的。
				此外，如果你在结构体字段上使用了标签（例如 json:"title"），这些标签可以帮助客户端库在序列化和反序列化时将 Golang 结构体与 Elasticsearch 文档正确地映射。
				struct { ... }{ ... }: 这是一个匿名结构体，用于描述要更新的文档的内容。
		
			struct {
				Info struct {
					Title string `json:"title"`
				} `json:"info"`
			}{
				Info: struct {
					Title string `json:"title"`
				}{
					Title: newTitle,
				},
			},
		)*/
		Doc(
			struct{
				AuthorId int64 `json:"author_id"`
			}{
				AuthorId: newId,
			},
		).Do(context.Background())
	return err
}
