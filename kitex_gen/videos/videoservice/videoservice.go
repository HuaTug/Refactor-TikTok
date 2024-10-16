// Code generated by Kitex v0.10.3. DO NOT EDIT.

package videoservice

import (
	videos "HuaTug.com/kitex_gen/videos"
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"FeedService": kitex.NewMethodInfo(
		feedServiceHandler,
		newVideoServiceFeedServiceArgs,
		newVideoServiceFeedServiceResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"VideoPublishStart": kitex.NewMethodInfo(
		videoPublishStartHandler,
		newVideoServiceVideoPublishStartArgs,
		newVideoServiceVideoPublishStartResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"VideoPublishUploading": kitex.NewMethodInfo(
		videoPublishUploadingHandler,
		newVideoServiceVideoPublishUploadingArgs,
		newVideoServiceVideoPublishUploadingResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"VideoPublishComplete": kitex.NewMethodInfo(
		videoPublishCompleteHandler,
		newVideoServiceVideoPublishCompleteArgs,
		newVideoServiceVideoPublishCompleteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"VideoPublishCancle": kitex.NewMethodInfo(
		videoPublishCancleHandler,
		newVideoServiceVideoPublishCancleArgs,
		newVideoServiceVideoPublishCancleResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"VideoDelete": kitex.NewMethodInfo(
		videoDeleteHandler,
		newVideoServiceVideoDeleteArgs,
		newVideoServiceVideoDeleteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"VideoIdList": kitex.NewMethodInfo(
		videoIdListHandler,
		newVideoServiceVideoIdListArgs,
		newVideoServiceVideoIdListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"VideoFeedList": kitex.NewMethodInfo(
		videoFeedListHandler,
		newVideoServiceVideoFeedListArgs,
		newVideoServiceVideoFeedListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"VideoSearch": kitex.NewMethodInfo(
		videoSearchHandler,
		newVideoServiceVideoSearchArgs,
		newVideoServiceVideoSearchResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"VideoPopular": kitex.NewMethodInfo(
		videoPopularHandler,
		newVideoServiceVideoPopularArgs,
		newVideoServiceVideoPopularResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"VideoInfo": kitex.NewMethodInfo(
		videoInfoHandler,
		newVideoServiceVideoInfoArgs,
		newVideoServiceVideoInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"VideoVisit": kitex.NewMethodInfo(
		videoVisitHandler,
		newVideoServiceVideoVisitArgs,
		newVideoServiceVideoVisitResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UpdateVisitCount": kitex.NewMethodInfo(
		updateVisitCountHandler,
		newVideoServiceUpdateVisitCountArgs,
		newVideoServiceUpdateVisitCountResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetVideoVisitCount": kitex.NewMethodInfo(
		getVideoVisitCountHandler,
		newVideoServiceGetVideoVisitCountArgs,
		newVideoServiceGetVideoVisitCountResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetVideoVisitCountInRedis": kitex.NewMethodInfo(
		getVideoVisitCountInRedisHandler,
		newVideoServiceGetVideoVisitCountInRedisArgs,
		newVideoServiceGetVideoVisitCountInRedisResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	videoServiceServiceInfo                = NewServiceInfo()
	videoServiceServiceInfoForClient       = NewServiceInfoForClient()
	videoServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return videoServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return videoServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*videos.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "videos",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.10.3",
		Extra:           extra,
	}
	return svcInfo
}

func feedServiceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videos.VideoServiceFeedServiceArgs)
	realResult := result.(*videos.VideoServiceFeedServiceResult)
	success, err := handler.(videos.VideoService).FeedService(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFeedServiceArgs() interface{} {
	return videos.NewVideoServiceFeedServiceArgs()
}

func newVideoServiceFeedServiceResult() interface{} {
	return videos.NewVideoServiceFeedServiceResult()
}

func videoPublishStartHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videos.VideoServiceVideoPublishStartArgs)
	realResult := result.(*videos.VideoServiceVideoPublishStartResult)
	success, err := handler.(videos.VideoService).VideoPublishStart(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceVideoPublishStartArgs() interface{} {
	return videos.NewVideoServiceVideoPublishStartArgs()
}

func newVideoServiceVideoPublishStartResult() interface{} {
	return videos.NewVideoServiceVideoPublishStartResult()
}

func videoPublishUploadingHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videos.VideoServiceVideoPublishUploadingArgs)
	realResult := result.(*videos.VideoServiceVideoPublishUploadingResult)
	success, err := handler.(videos.VideoService).VideoPublishUploading(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceVideoPublishUploadingArgs() interface{} {
	return videos.NewVideoServiceVideoPublishUploadingArgs()
}

func newVideoServiceVideoPublishUploadingResult() interface{} {
	return videos.NewVideoServiceVideoPublishUploadingResult()
}

func videoPublishCompleteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videos.VideoServiceVideoPublishCompleteArgs)
	realResult := result.(*videos.VideoServiceVideoPublishCompleteResult)
	success, err := handler.(videos.VideoService).VideoPublishComplete(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceVideoPublishCompleteArgs() interface{} {
	return videos.NewVideoServiceVideoPublishCompleteArgs()
}

func newVideoServiceVideoPublishCompleteResult() interface{} {
	return videos.NewVideoServiceVideoPublishCompleteResult()
}

func videoPublishCancleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videos.VideoServiceVideoPublishCancleArgs)
	realResult := result.(*videos.VideoServiceVideoPublishCancleResult)
	success, err := handler.(videos.VideoService).VideoPublishCancle(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceVideoPublishCancleArgs() interface{} {
	return videos.NewVideoServiceVideoPublishCancleArgs()
}

func newVideoServiceVideoPublishCancleResult() interface{} {
	return videos.NewVideoServiceVideoPublishCancleResult()
}

func videoDeleteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videos.VideoServiceVideoDeleteArgs)
	realResult := result.(*videos.VideoServiceVideoDeleteResult)
	success, err := handler.(videos.VideoService).VideoDelete(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceVideoDeleteArgs() interface{} {
	return videos.NewVideoServiceVideoDeleteArgs()
}

func newVideoServiceVideoDeleteResult() interface{} {
	return videos.NewVideoServiceVideoDeleteResult()
}

func videoIdListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videos.VideoServiceVideoIdListArgs)
	realResult := result.(*videos.VideoServiceVideoIdListResult)
	success, err := handler.(videos.VideoService).VideoIdList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceVideoIdListArgs() interface{} {
	return videos.NewVideoServiceVideoIdListArgs()
}

func newVideoServiceVideoIdListResult() interface{} {
	return videos.NewVideoServiceVideoIdListResult()
}

func videoFeedListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videos.VideoServiceVideoFeedListArgs)
	realResult := result.(*videos.VideoServiceVideoFeedListResult)
	success, err := handler.(videos.VideoService).VideoFeedList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceVideoFeedListArgs() interface{} {
	return videos.NewVideoServiceVideoFeedListArgs()
}

func newVideoServiceVideoFeedListResult() interface{} {
	return videos.NewVideoServiceVideoFeedListResult()
}

func videoSearchHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videos.VideoServiceVideoSearchArgs)
	realResult := result.(*videos.VideoServiceVideoSearchResult)
	success, err := handler.(videos.VideoService).VideoSearch(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceVideoSearchArgs() interface{} {
	return videos.NewVideoServiceVideoSearchArgs()
}

func newVideoServiceVideoSearchResult() interface{} {
	return videos.NewVideoServiceVideoSearchResult()
}

func videoPopularHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videos.VideoServiceVideoPopularArgs)
	realResult := result.(*videos.VideoServiceVideoPopularResult)
	success, err := handler.(videos.VideoService).VideoPopular(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceVideoPopularArgs() interface{} {
	return videos.NewVideoServiceVideoPopularArgs()
}

func newVideoServiceVideoPopularResult() interface{} {
	return videos.NewVideoServiceVideoPopularResult()
}

func videoInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videos.VideoServiceVideoInfoArgs)
	realResult := result.(*videos.VideoServiceVideoInfoResult)
	success, err := handler.(videos.VideoService).VideoInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceVideoInfoArgs() interface{} {
	return videos.NewVideoServiceVideoInfoArgs()
}

func newVideoServiceVideoInfoResult() interface{} {
	return videos.NewVideoServiceVideoInfoResult()
}

func videoVisitHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videos.VideoServiceVideoVisitArgs)
	realResult := result.(*videos.VideoServiceVideoVisitResult)
	success, err := handler.(videos.VideoService).VideoVisit(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceVideoVisitArgs() interface{} {
	return videos.NewVideoServiceVideoVisitArgs()
}

func newVideoServiceVideoVisitResult() interface{} {
	return videos.NewVideoServiceVideoVisitResult()
}

func updateVisitCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videos.VideoServiceUpdateVisitCountArgs)
	realResult := result.(*videos.VideoServiceUpdateVisitCountResult)
	success, err := handler.(videos.VideoService).UpdateVisitCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceUpdateVisitCountArgs() interface{} {
	return videos.NewVideoServiceUpdateVisitCountArgs()
}

func newVideoServiceUpdateVisitCountResult() interface{} {
	return videos.NewVideoServiceUpdateVisitCountResult()
}

func getVideoVisitCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videos.VideoServiceGetVideoVisitCountArgs)
	realResult := result.(*videos.VideoServiceGetVideoVisitCountResult)
	success, err := handler.(videos.VideoService).GetVideoVisitCount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceGetVideoVisitCountArgs() interface{} {
	return videos.NewVideoServiceGetVideoVisitCountArgs()
}

func newVideoServiceGetVideoVisitCountResult() interface{} {
	return videos.NewVideoServiceGetVideoVisitCountResult()
}

func getVideoVisitCountInRedisHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*videos.VideoServiceGetVideoVisitCountInRedisArgs)
	realResult := result.(*videos.VideoServiceGetVideoVisitCountInRedisResult)
	success, err := handler.(videos.VideoService).GetVideoVisitCountInRedis(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceGetVideoVisitCountInRedisArgs() interface{} {
	return videos.NewVideoServiceGetVideoVisitCountInRedisArgs()
}

func newVideoServiceGetVideoVisitCountInRedisResult() interface{} {
	return videos.NewVideoServiceGetVideoVisitCountInRedisResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FeedService(ctx context.Context, req *videos.FeedServiceRequest) (r *videos.FeedServiceResponse, err error) {
	var _args videos.VideoServiceFeedServiceArgs
	_args.Req = req
	var _result videos.VideoServiceFeedServiceResult
	if err = p.c.Call(ctx, "FeedService", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoPublishStart(ctx context.Context, req *videos.VideoPublishStartRequest) (r *videos.VideoPublishStartResponse, err error) {
	var _args videos.VideoServiceVideoPublishStartArgs
	_args.Req = req
	var _result videos.VideoServiceVideoPublishStartResult
	if err = p.c.Call(ctx, "VideoPublishStart", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoPublishUploading(ctx context.Context, req *videos.VideoPublishUploadingRequest) (r *videos.VideoPublishUploadingResponse, err error) {
	var _args videos.VideoServiceVideoPublishUploadingArgs
	_args.Req = req
	var _result videos.VideoServiceVideoPublishUploadingResult
	if err = p.c.Call(ctx, "VideoPublishUploading", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoPublishComplete(ctx context.Context, req *videos.VideoPublishCompleteRequest) (r *videos.VideoPublishCompleteResponse, err error) {
	var _args videos.VideoServiceVideoPublishCompleteArgs
	_args.Req = req
	var _result videos.VideoServiceVideoPublishCompleteResult
	if err = p.c.Call(ctx, "VideoPublishComplete", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoPublishCancle(ctx context.Context, req *videos.VideoPublishCancleRequest) (r *videos.VideoPublishCancleResponse, err error) {
	var _args videos.VideoServiceVideoPublishCancleArgs
	_args.Req = req
	var _result videos.VideoServiceVideoPublishCancleResult
	if err = p.c.Call(ctx, "VideoPublishCancle", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoDelete(ctx context.Context, req *videos.VideoDeleteRequest) (r *videos.VideoDeleteResponse, err error) {
	var _args videos.VideoServiceVideoDeleteArgs
	_args.Req = req
	var _result videos.VideoServiceVideoDeleteResult
	if err = p.c.Call(ctx, "VideoDelete", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoIdList(ctx context.Context, req *videos.VideoIdListRequest) (r *videos.VideoIdListResponse, err error) {
	var _args videos.VideoServiceVideoIdListArgs
	_args.Req = req
	var _result videos.VideoServiceVideoIdListResult
	if err = p.c.Call(ctx, "VideoIdList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoFeedList(ctx context.Context, req *videos.VideoFeedListRequest) (r *videos.VideoFeedListResponse, err error) {
	var _args videos.VideoServiceVideoFeedListArgs
	_args.Req = req
	var _result videos.VideoServiceVideoFeedListResult
	if err = p.c.Call(ctx, "VideoFeedList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoSearch(ctx context.Context, req *videos.VideoSearchRequest) (r *videos.VideoSearchResponse, err error) {
	var _args videos.VideoServiceVideoSearchArgs
	_args.Req = req
	var _result videos.VideoServiceVideoSearchResult
	if err = p.c.Call(ctx, "VideoSearch", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoPopular(ctx context.Context, req *videos.VideoPopularRequest) (r *videos.VideoPopularResponse, err error) {
	var _args videos.VideoServiceVideoPopularArgs
	_args.Req = req
	var _result videos.VideoServiceVideoPopularResult
	if err = p.c.Call(ctx, "VideoPopular", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoInfo(ctx context.Context, req *videos.VideoInfoRequest) (r *videos.VideoInfoResponse, err error) {
	var _args videos.VideoServiceVideoInfoArgs
	_args.Req = req
	var _result videos.VideoServiceVideoInfoResult
	if err = p.c.Call(ctx, "VideoInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VideoVisit(ctx context.Context, req *videos.VideoVisitRequest) (r *videos.VideoVisitResponse, err error) {
	var _args videos.VideoServiceVideoVisitArgs
	_args.Req = req
	var _result videos.VideoServiceVideoVisitResult
	if err = p.c.Call(ctx, "VideoVisit", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateVisitCount(ctx context.Context, req *videos.UpdateVisitCountRequest) (r *videos.UpdateVisitCountResponse, err error) {
	var _args videos.VideoServiceUpdateVisitCountArgs
	_args.Req = req
	var _result videos.VideoServiceUpdateVisitCountResult
	if err = p.c.Call(ctx, "UpdateVisitCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetVideoVisitCount(ctx context.Context, req *videos.GetVideoVisitCountRequest) (r *videos.GetVideoVisitCountResponse, err error) {
	var _args videos.VideoServiceGetVideoVisitCountArgs
	_args.Req = req
	var _result videos.VideoServiceGetVideoVisitCountResult
	if err = p.c.Call(ctx, "GetVideoVisitCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetVideoVisitCountInRedis(ctx context.Context, req *videos.GetVideoVisitCountInRedisRequest) (r *videos.GetVideoVisitCountInRedisResponse, err error) {
	var _args videos.VideoServiceGetVideoVisitCountInRedisArgs
	_args.Req = req
	var _result videos.VideoServiceGetVideoVisitCountInRedisResult
	if err = p.c.Call(ctx, "GetVideoVisitCountInRedis", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
