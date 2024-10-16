namespace go videos

include "base.thrift"

struct FeedServiceRequest{
    1: string last_time
}
struct FeedServiceResponse{
    1: base.Status base
    2: list<base.Video> video_list
}

struct VideoPublishStartRequest{
    1: i64 user_id
    2: string title (vt.min_size="1")
    3: string description
    4: i64 chunk_total_number (vt.gt="0")
}
struct VideoPublishStartResponse{
    1: base.Status base
    2: string uuid
}

struct VideoPublishUploadingRequest{
    1: i64 user_id
    2: string uuid //唯一标识符
    3: binary data //视频数据块 以二进制形式存储
    4: string md5   //上传数据的MD5哈希值
    5: bool is_m3u8
    6: string filename
    7: i64 chunk_number //分片序号
}
struct VideoPublishUploadingResponse{
    1: base.Status base
}

struct VideoPublishCompleteRequest{
    1: i64 user_id
    2: string uuid
}
struct VideoPublishCompleteResponse{
    1: base.Status base
}

struct VideoPublishCancleRequest{
    1: i64 user_id
    2: string uuid
}
struct VideoPublishCancleResponse{
    1: base.Status base
}

struct VideoFeedListRequest{
    1: i64 user_id
    2: i64 page_num
    3: i64 page_size
}
struct VideoFeedListResponse{
    1: base.Status base
    2: list<base.Video> video_list
    3: i64 total
}

struct VideoSearchRequest{
    1: string keyword
    2: i64 page_num
    3: i64 page_size
    4: string from_date
    5: string to_date
}
struct VideoSearchResponse{
    1: base.Status base
    2: list<base.Video> video_search
    3: i64 count
}

struct VideoPopularRequest{
    1: i64 page_num
    2: i64 page_size
}
struct VideoPopularResponse{
    1: base.Status base
    2: list<base.Video> Popular
}

struct VideoInfoRequest{
    1: i64 video_id
}
struct VideoInfoResponse{
    1: base.Status base
    2: base.Video items
}

struct VideoDeleteRequest{
    1: i64 user_id
    2: i64 video_id
}
struct VideoDeleteResponse{
    1: base.Status base
}

struct VideoVisitRequest{
    1: i64 from_id
    2: i64 video_id
}
struct VideoVisitResponse{
    1: base.Status base
    2: base.Video item
}

struct VideoIdListRequest{
    1: i64 page_num
    2: i64 page_size
}
struct VideoIdListResponse{
    1: base.Status base
    2: bool is_end
    3: list<string> list
}

struct UpdateVisitCountRequest{
    1: i64 video_id
    2: i64 visit_count
}
struct UpdateVisitCountResponse{
    1: base.Status base
}

struct GetVideoVisitCountRequest{
    1: i64 video_id
}
struct GetVideoVisitCountResponse{
    1: base.Status base
    2: i64 visit_count
}

struct GetVideoVisitCountInRedisRequest{
    1: i64 video_id
}
struct GetVideoVisitCountInRedisResponse{
    1: i64 visit_count
    2: base.Status base
}

service VideoService {
    FeedServiceResponse FeedService(1: FeedServiceRequest req)(api.get="/v1/video/feed")
    VideoPublishStartResponse VideoPublishStart(1: VideoPublishStartRequest req)(api.post="/v1/publish/start")
    VideoPublishUploadingResponse VideoPublishUploading(1 :VideoPublishUploadingRequest req)(api.post="/v1/publish/uploading")
    VideoPublishCompleteResponse VideoPublishComplete(1: VideoPublishCompleteRequest req)(api.post="/v1/publish/complete")
    VideoPublishCancleResponse VideoPublishCancle(1: VideoPublishCancleRequest req)(api.post="/v1/publish/cancle")
    VideoDeleteResponse VideoDelete(1: VideoDeleteRequest req)(api.delete="/v1/video/delete")
    VideoIdListResponse VideoIdList(1: VideoIdListRequest req)
    VideoFeedListResponse VideoFeedList(1: VideoFeedListRequest req)(api.get="/v1/video/list")
    VideoSearchResponse  VideoSearch(1: VideoSearchRequest req)(api.post="/v1/video/search")
    VideoPopularResponse VideoPopular(1: VideoPopularRequest req)(api.get="/v1/video/popular")
    VideoInfoResponse VideoInfo(1: VideoInfoRequest req)
    VideoVisitResponse VideoVisit(1: VideoVisitRequest req)(api.get="/v1/visit/:id")
    UpdateVisitCountResponse UpdateVisitCount(1: UpdateVisitCountRequest req)
    GetVideoVisitCountResponse GetVideoVisitCount(1: GetVideoVisitCountRequest req)
    GetVideoVisitCountInRedisResponse GetVideoVisitCountInRedis(1: GetVideoVisitCountInRedisRequest req)
}
