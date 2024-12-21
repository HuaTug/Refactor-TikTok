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
    4: string lab_name
    5: string category
    6: i64 open
    7: i64 chunk_total_number (vt.gt="0")
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

//
struct UpdateVideoCommentCountRequest{
    1: i64 video_id
    2: i64 comment_count
}
struct UpdateVideoCommentCountResponse{
    1: base.Status base
}

// 更新视频点赞数 
struct UpdateLikeCountRequest{
    1: i64 video_id
    2: i64 like_count
}
struct UpdateLikeCountResponse{
    1: base.Status base
}

struct UpdateVideoHisLikeCountRequest{
    1: i64 video_id
    2: i64 his_like_count
}
struct UpdateVideoHisLikeCountResponse{
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

// 视频流播放
struct StreamVideoRequest{
    1: string index
}
struct StreamVideoResponse{
    1: base.Status base
    2: byte data
}

// 创建收藏夹
struct CreateFavoriteRequest{
    1: i64 user_id
    2: string name
    3: string description
    4: string cover_url
}
struct CreateFavoriteResponse{
    1: base.Status base
}

// 获取收藏夹列表 (获取用户收藏夹列表，全是收藏夹，没有视频信息)
struct GetFavoriteListRequest{
    1: i64 user_id
    2: i64 page_num
    3: i64 page_size
}
struct GetFavoriteListResponse{
    1: base.Status base
    2: list<base.Favorite> favorite_list
}

// 添加收藏夹视频
struct AddFavoriteVideoRequest{
    1: i64 favorite_id
    2: i64 user_id
    3: i64 video_id

}
struct AddFavoriteVideoResponse{
    1: base.Status base
}

// 获取收藏夹视频列表
struct GetFavoriteVideoListRequest{
    1: i64 user_id
    2: i64 favorite_id
    3: i64 page_num
    4: i64 page_size
}
struct GetFavoriteVideoListResponse{
    1: base.Status base
    2: list<base.Video> video_list
}

// 获取收藏夹中的某一个视频
struct GetVideoFromFavoriteRequest{
    1: i64 favorite_id
    2: i64 user_id
    3: i64 video_id
    4: i64 page_num
    5: i64 page_size
}
struct GetVideoFromFavoriteResponse{
    1: base.Status base
    2: base.Video video
}

struct DeleteFavoriteRequest{
    1: i64 user_id
    2: i64 favorite_id
}
struct DeleteFavoriteResponse{
    1: base.Status base
}

struct DeleteVideoFromFavoriteRequest{
    1: i64 favorite_id
    2: i64 user_id
    3: i64 video_id
}
struct DeleteVideoFromFavoriteResponse{
    1: base.Status base
}

// 分享视频
struct SharedVideoRequest{
    1: i64 user_id
    2: i64 to_user_id
    3: i64 video_id                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 
}
struct SharedVideoResponse{
    1:base.Status base
}

// 获取热门视频
struct GetPopularVideoRequest{
    1: i64 page_num
    2: i64 page_size
}
struct GetPopularVideoResponse{
    1: base.Status base
    2: list<base.Video> video_list
}

struct RecommendVideoRequest{
    1: i64 user_id
}
struct RecommendVideoResponse{
    1: base.Status base
    2: list<base.Video> video_list
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
    VideoVisitResponse VideoVisit(1: VideoVisitRequest req)(api.post="/v1/visit/:id")
    UpdateVisitCountResponse UpdateVisitCount(1: UpdateVisitCountRequest req)
    UpdateVideoCommentCountResponse UpdateVideoCommentCount(1: UpdateVideoCommentCountRequest req)
    UpdateLikeCountResponse UpdateVideoLikeCount(1: UpdateLikeCountRequest req)
    UpdateVideoHisLikeCountResponse UpdateVideoHisLikeCount(1: UpdateVideoHisLikeCountRequest req)
    GetVideoVisitCountResponse GetVideoVisitCount(1: GetVideoVisitCountRequest req)
    GetVideoVisitCountInRedisResponse GetVideoVisitCountInRedis(1: GetVideoVisitCountInRedisRequest req)
    StreamVideoResponse StreamVideo(1: StreamVideoRequest req) (api.post="/v1/stream")
    CreateFavoriteResponse CreateFavorite(1: CreateFavoriteRequest req) (api.post="/v1/favorite/create")        
    GetFavoriteVideoListResponse GetFavoriteVideoList(1: GetFavoriteVideoListRequest req)(api.get="/v1/favorite/video/list")
    GetFavoriteListResponse GetFavoriteList(1: GetFavoriteListRequest req)(api.get="/v1/favorite/list")
    GetVideoFromFavoriteResponse GetVideoFromFavorite(1: GetVideoFromFavoriteRequest req)(api.get="/v1/favorite/video")
    AddFavoriteVideoResponse AddFavoriteVideo(1: AddFavoriteVideoRequest req)(api.post="/v1/favorite/video/add") 
    DeleteFavoriteResponse DeleteFavorite(1: DeleteFavoriteRequest req)(api.delete="/v1/favorite/delete")
    DeleteVideoFromFavoriteResponse DeleteVideoFromFavorite(1: DeleteVideoFromFavoriteRequest req)(api.delete="/v1/favorite/video/delete")
    SharedVideoResponse SharedVideo(1: SharedVideoRequest req)(api.post="/v1/share/video")
    RecommendVideoResponse RecommendVideo(1: RecommendVideoRequest req)(api.get="/v1/recommend/video")
}
