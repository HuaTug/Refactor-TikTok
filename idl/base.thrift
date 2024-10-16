namespace go base

struct Status {
    1: i64 code
    2: string msg
}

struct User {
    1: i64 user_id
    2: string user_name
    3: string password
    4: string avatar_url
    5: string created_at
    6: string updated_at
    7: string deleted_at
}

struct UserLite {
    1: i64 uid
    2: string user_name
    3: string avatar_url
}

struct Video{
    1: i64 video_id
    2: i64 user_id
    3: string video_url
    4: string cover_url
    5: string title
    6: string description
    7: i64 visit_count
    8: i64 like_count
    9: i64 comment_count
    10: string created_at
    11: string updated_at
    12: string deleted_at
}

struct Comment{
    1: i64 comment_id
    2: i64 user_id
    3: i64 video_id
    4: i64 parent_id
    5: i64 like_count
    6: i64 child_count
    7: string content
    8: string created_at
    9: string updated_at
    10: string deleted_at
}
