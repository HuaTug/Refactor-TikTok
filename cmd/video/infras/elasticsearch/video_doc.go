package elasticsearch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"sync"

	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/pkg/constants"
	utils "HuaTug.com/pkg/utils"
	"github.com/olivere/elastic/v7"
)

type VideoOtherdata struct {
	VideoId      int64  `json:"id"`
	VideoUrl     string `json:"video_url"`
	CoverUrl     string `json:"cover_url"`
	VisitCount   int64  `json:"visit_count"`
	LikeCount    int64  `json:"like_count"`
	CommentCount int64  `json:"comment_count"`
	UpdatedAt    int64  `json:"updated_at"`
	DeletedAt    int64  `json:"deleted_at"`
}

type Video struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Username    string         `json:"username"`
	UserId      int64          `json:"user_id"`
	CreatedAt   int64          `json:"created_at"`
	Info        VideoOtherdata `json:"info"`
}

func CreateVideoDoc(data *Video) error {
	_, err := elasticClient.Index().
		Index("video").
		Type("_doc").
		Id(fmt.Sprint(data.Info.VideoId)).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func DeleteVideoDoc(vid string) error {
	_, err := elasticClient.Delete().
		Index("video").
		Type("_doc").
		Id(vid).
		Do(context.Background())
	return err
}

func searchRespCovert(resp *elastic.SearchResult) ([]*base.Video, int64) {
	dataList := make([]*base.Video, 0)
	for _, item := range resp.Each(reflect.TypeOf(Video{})) {
		d, ok := item.(Video)
		if !ok {
			continue
		}
		temp := base.Video{
			VideoId:      d.Info.VideoId,
			UserId:       d.UserId,
			VideoUrl:     d.Info.VideoUrl,
			CoverUrl:     d.Info.CoverUrl,
			Title:        d.Title,
			Description:  d.Description,
			VisitCount:   d.Info.VisitCount,
			LikesCount:   d.Info.LikeCount,
			CommentCount: d.Info.CommentCount,
			CreatedAt:    utils.ConvertTimestampToString(d.CreatedAt),
			UpdatedAt:    utils.ConvertTimestampToString(d.Info.UpdatedAt),
			DeletedAt:    utils.ConvertTimestampToString(d.Info.DeletedAt),
		}
		dataList = append(dataList, &temp)
	}
	hits := resp.TotalHits()
	return dataList, hits
}

func SearchVideoDocDefault(keywords string) ([]*base.Video, int64, error) {
	resp, err := elasticClient.Search().
		Index("video").
		Type("_doc").
		Query(elastic.NewBoolQuery().Should(
			elastic.NewMultiMatchQuery(keywords, "title", "description"),
		)).
		From(0).Size(constants.DefaultPageSize).
		Do(context.Background())

	if err != nil {
		return nil, -1, err
	}
	data, hits := searchRespCovert(resp)
	return data, hits, nil
}

func RandomVideoDoc(fromDate int64) ([]*base.Video, int64, error) {
	resp, err := elasticClient.Search().
		Index("video").
		Type("_doc").
		Query(elastic.NewFunctionScoreQuery().
			AddScoreFunc(elastic.NewRandomFunction()).
			Query(elastic.NewRangeQuery("created_at").From(fromDate)),
		).
		Size(constants.DefaultPageSize).
		Do(context.Background())
	if err != nil {
		return nil, -1, err
	}
	data, hits := searchRespCovert(resp)
	return data, hits, nil
}

func SearchVideoDoc(keywords, username string, pageSize, pageNum, fromDate, toDate int64) ([]*base.Video, int64, error) {
	var (
		mustQuery    = make([]elastic.Query, 0)
		mustnotQuery = make([]elastic.Query, 0)
		shouldQuery  = make([]elastic.Query, 0)
		filterQuery  = make([]elastic.Query, 0)
	)
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		if keywords != constants.ESNoKeywordsFlag {
			shouldQuery = append(mustQuery, elastic.NewMultiMatchQuery(keywords, "title", "description"))
		}
		wg.Done()
	}()
	go func() {
		if username != constants.ESNoUsernameFilterFlag {
			filterQuery = append(shouldQuery, elastic.NewTermQuery("username", username))
		}
		wg.Done()
	}()
	go func() {
		if fromDate != constants.ESNoTimeFilterFlag && toDate != constants.ESNoTimeFilterFlag {
			filterQuery = append(shouldQuery, elastic.NewRangeQuery("created_at").From(fromDate).To(toDate))
		} else if fromDate != constants.ESNoTimeFilterFlag && toDate == constants.ESNoTimeFilterFlag {
			filterQuery = append(shouldQuery, elastic.NewRangeQuery("created_at").From(fromDate))
		} else if toDate != constants.ESNoTimeFilterFlag && fromDate == constants.ESNoTimeFilterFlag {
			filterQuery = append(filterQuery, elastic.NewRangeQuery("created_at").To(toDate))
		}
		wg.Done()
	}()
	go func() {
		if pageNum == constants.ESNoPageParamFlag {
			pageNum = 1
		}
		wg.Done()
	}()
	go func() {
		if pageSize == constants.ESNoPageParamFlag {
			pageSize = constants.DefaultPageSize
		}
		wg.Done()
	}()
	wg.Wait()
	query := elastic.NewBoolQuery().
		Should(shouldQuery...).
		Must(mustQuery...).
		MustNot(mustnotQuery...).
		Filter(filterQuery...)
	resp, err := elasticClient.Search().
		Index("video").
		Type("_doc").
		Query(query).
		From(int(pageNum-1) * int(pageSize)).Size(int(pageSize)).
		Do(context.Background())

	if err != nil {
		return nil, -1, err
	}

	data, hits := searchRespCovert(resp)
	return data, hits, nil
}

func SearchVideoDocByUserId(uid string, pageNum, pageSize int64) ([]*base.Video, int64, error) {
	var (
		mustQuery    = make([]elastic.Query, 0)
		mustnotQuery = make([]elastic.Query, 0)
		shouldQuery  = make([]elastic.Query, 0)
		filterQuery  = make([]elastic.Query, 0)
	)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		filterQuery = append(filterQuery, elastic.NewTermQuery("user_id", uid))
		wg.Done()
	}()
	go func() {
		if pageNum == constants.ESNoPageParamFlag {
			pageNum = 1
		}
		wg.Done()
	}()
	go func() {
		if pageSize == constants.ESNoPageParamFlag {
			pageSize = constants.DefaultPageSize
		}
		wg.Done()
	}()
	wg.Wait()
	query := elastic.NewBoolQuery().
		Should(shouldQuery...).
		Must(mustQuery...).
		MustNot(mustnotQuery...).
		Filter(filterQuery...)

	resp, err := elasticClient.Search().
		Index("video").
		Type("_doc").
		Query(query).
		From(int(pageNum-1) * int(pageSize)).Size(int(pageSize)).
		Do(context.Background())
	if err != nil {
		return nil, -1, err
	}
	data, hits := searchRespCovert(resp)
	return data, hits, nil
}

// 覆盖更新
func UpdateVideoDoc(info *Video) error {
	return CreateVideoDoc(info)
}

func UpdateVideoVisitCount(vid, visitCount string) error {
	bulk := elasticClient.Bulk()
	var (
		newVisitCount, _ = strconv.ParseInt(visitCount, 10, 64)
	)
	vRequest := elastic.NewBulkUpdateRequest().Index("video").Type("_doc").Id(vid).
		Script(elastic.NewScript(`"ctx._source.info.visit_count=params.new_visit_count"`).Param("new_visit_count", newVisitCount))
	bulk.Add(vRequest)
	if _, err := bulk.Do(context.Background()); err != nil {
		return err
	}
	return nil
}

func OnlyUpdateUsername(newUsername, userId string) error {
	_, err := elasticClient.UpdateByQuery().
		Index("video").
		Type("_doc").
		Query(elastic.NewBoolQuery().Filter(
			elastic.NewMatchQuery("user_id", userId),
		)).Script(elastic.NewScript(`"ctx._source.username=params.new_username"`).Param("new_username", newUsername)).
		Do(context.Background())
	return err
}

func OnlyUpdateTitle(newTitle, vid string) error {
	_, err := elasticClient.Update().Index("video").Type("_doc").Id(vid).
		Doc(
			struct {
				Title string `json:"title"`
			}{
				Title: newTitle,
			},
		).
		Do(context.Background())
	return err
}

func OnlyUpdateDescription(newDescription, vid string) error {
	_, err := elasticClient.Update().Index("video").Type("_doc").Id(vid).
		Doc(
			struct {
				Description string `json:"description"`
			}{
				Description: newDescription,
			},
		).Do(context.Background())
	return err
}

func UpdateTitleAndDescription(newTitle, newDescription, vid string) error {
	_, err := elasticClient.Update().Index("video").Type("_doc").Id(vid).
		Doc(
			struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			}{
				Title:       newTitle,
				Description: newDescription,
			},
		).Do(context.Background())
	return err
}

func GetVideoDoc(id string) (*base.Video, error) {
	resp, err := elasticClient.Get().
		Index("video").
		Type("_doc").
		Id(id).
		Do(context.Background())
	if err != nil {
		return nil, err
	}
	if !resp.Found {
		return nil, errors.New("video not found")
	}
	data, _ := resp.Source.MarshalJSON()
	var d Video
	json.Unmarshal(data, &d)
	return &base.Video{
		VideoId:      d.Info.VideoId,
		UserId:       d.UserId,
		VideoUrl:     d.Info.VideoUrl,
		CoverUrl:     d.Info.CoverUrl,
		Title:        d.Title,
		Description:  d.Description,
		VisitCount:   d.Info.VisitCount,
		LikesCount:   d.Info.LikeCount,
		CommentCount: d.Info.CommentCount,
		CreatedAt:    utils.ConvertTimestampToString(d.CreatedAt),
		UpdatedAt:    utils.ConvertTimestampToString(d.Info.UpdatedAt),
		DeletedAt:    utils.ConvertTimestampToString(d.Info.DeletedAt),
	}, nil
}
