namespace go relations

include "base.thrift"

struct RelationServiceRequest {
    1: i64 action_type
    2: i64 from_user_id
    3: i64 to_user_id 
}
struct RelationServiceResponse {
    1: base.Status base
}

struct FollowingListRequest {
    1: i64 user_id 
    2: i64 page_num (vt.ge="0")
    3: i64 page_size (vt.gt="0")
}
struct FollowingListResponse {
    1: base.Status base
    2: list<base.UserLite> items
    3: i64 total
}

struct FollowerListRequest {
    1: i64 user_id 
    2: i64 page_num (vt.ge="0")
    3: i64 page_size (vt.gt="0")    
}
struct FollowerListResponse {
    1: base.Status base
    2: list<base.UserLite> items
    3: i64 total
}

struct FriendListRequest {
    1: i64 user_id 
    2: i64 page_num (vt.ge="0")
    3: i64 page_size (vt.gt="0")  
}
struct FriendListResponse {
    1: base.Status base
    2: list<base.UserLite> items
    3: i64 total    
}

service FollowService {
    RelationServiceResponse RelationService (1: RelationServiceRequest req)(api.post="/v1/relation/action")
    FollowingListResponse FollowingList (1: FollowingListRequest req)(api.get="/v1/following/list")
    FollowerListResponse FollowerList (1: FollowerListRequest req)(api.get="/v1/follower/list")
    FriendListResponse FriendList (1: FriendListRequest req)(api.get="/v1/friend/list")
}