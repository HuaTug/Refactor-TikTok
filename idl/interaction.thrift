namespace go interactions

include "base.thrift"

struct LikeActionRequest {
    1: i64 user_id
    2: i64 video_id
    3: i64 comment_id
    4: string action_type
}
struct LikeActionResponse {
    1: base.Status base
}

struct LikeListRequest {
    1: i64 user_id
    2: i64 page_num
    3: i64 page_size
}
struct LikeListResponse {
    1: base.Status base
    2: list<base.Video> items
}

struct CreateCommentRequest {
    1: i64 user_id
    2: i64 video_id 
    3: i64 comment_id
    4: i64 mode
    5: string content 
}

struct CreateCommentResponse {
    1: base.Status base
}
struct ListCommentRequest {
    1: i64 video_id
    2: i64 comment_id
    3: i64 page_num
    4: i64 page_size
}
struct ListCommentResponse {
    1: base.Status base
    2: list<base.Comment> items
}

struct CommentDeleteRequest {
    1: i64 video_id    
    2: i64 comment_id
    3: i64 from_user_id
}    
struct CommentDeleteResponse {
    1: base.Status base
}

struct VideoPopularListRequest {
    1: i64 page_num
    2: i64 page_size
}
struct VideoPopularListResponse {
    1: base.Status base
    2: list<string> data
}

// 删除与视频相关的信息
struct DeleteVideoInfoRequest {
    1: i64 video_id
}
struct DeleteVideoInfoResponse {
    1: base.Status base
}

service InteractionService {
    LikeActionResponse LikeAction(1: LikeActionRequest req)(api.post="/v1/action/like")
    LikeListResponse LikeList(1: LikeListRequest req)(api.get="/v1/action/list")
    CreateCommentResponse CreateComment(1:CreateCommentRequest req)(api.post="/v1/comment/publish")
    ListCommentResponse ListComment(1:ListCommentRequest req)(api.get="/v1/comment/list")
    CommentDeleteResponse DeleteComment(1:CommentDeleteRequest req)(api.delete="/v1/comment/delete")
    VideoPopularListResponse VideoPopularList(1: VideoPopularListRequest req)
    DeleteVideoInfoResponse DeleteVideoInfo(1: DeleteVideoInfoRequest req)
}    