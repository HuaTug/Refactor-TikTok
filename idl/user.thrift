namespace go users

include "base.thrift"


struct CreateUserRequest{
    1: string user_name      (api.body="user_name", api.form="user_name", api.vd="(len($) > 0 && len($) < 100)")
    2: string password (api.body="password", api.form="password", api.vd="len($)>5 &&len($)<12")
    3: string email (api.body="email", api.form="email", api.vd="(len($)>3 && len($) < 12)")
    4: i64 sex (api.body="sex", api.form="sex", api.vd="$ == 0 || $ == 1")
}

struct CreateUserResponse{ 
    1:base.Status base 
}

struct QueryUserRequest{
    1: optional string Keyword (api.body="keyword", api.form="keyword", api.query="keyword")
    2: i64 page (api.body="page", api.form="page", api.query="page", api.vd="$ > 0")
    3: i64 page_size (api.body="page_size", api.form="page_size", api.query="page_size", api.vd="($ > 0 || $ <= 100)")
}

struct QueryUserResponse{
    1: base.Status base
    3: list<base.User> users
    4: i64 totoal
}   

struct DeleteUserRequest{
    1: i64 userId
}

struct DeleteUserResponse{
    1: base.Status base
}

struct UpdateUserRequest{
    1: string user_name (api.body="user_name", api.form="user_name", api.vd="(len($) > 0 && len($) < 100)")
    2: i64 userId
    3: string password (api.body="password", api.form="password", api.vd="(len($)>5 &&len($)<12)")
    4: binary data
    5: i64 filesize
}

struct UpdateUserResponse{
    1: base.Status base
    2: base.User data
}

struct LoginUserResquest{
    1: string user_name   (api.body="user_name", api.form="user_name", api.vd="(len($)>3&&len($)<12)")
    2: string Password   (api.body="password", api.form="password", api.vd="(len($)>5&&len($)<10)")
    3: string Email      (api.body="email", api.form="email", api.vd="(len($)>3&&len($)<12)")
}

struct LoginUserResponse{
    1: base.Status base
    2: string token
    3: string RefreshToken
    4: base.User user
}

struct CheckUserExistsByIdRequst{
    1: i64 userId
}
struct CheckUserExistsByIdResponse{
    1: base.Status base
    2: bool exists
}

struct GetUserInfoRequest{
    1: i64 userId
}
struct GetUserInfoResponse{
    1: base.Status base
    2: base.User User
}

// 对验证码进行验证
struct VerifyCodeRequest{
    1: string code
    2: string email
}
struct VerifyCodeResponse{
    1: base.Status base
}

// 用户在登录时，需要发送验证码（以完成双重认证）
struct SendCodeRequest{
    1: string email
}
struct SendCodeResponse{
    1: base.Status base
}
service UserService {
   UpdateUserResponse UpdateUser(1: UpdateUserRequest req)(api.post="/v1/user/update")
   DeleteUserResponse DeleteUser(1: DeleteUserRequest req)(api.delete="/v1/user/delete")
   QueryUserResponse  QueryUser(1: QueryUserRequest req)(api.post="/v1/user/query")
   CreateUserResponse CreateUser(1: CreateUserRequest req)(api.post="/v1/user/create")
   LoginUserResponse  LoginUser(1: LoginUserResquest req)(api.post="/v1/user/login")
   GetUserInfoResponse GetUserInfo(1: GetUserInfoRequest req)(api.get="/v1/user/get")
   CheckUserExistsByIdResponse CheckUserExistsById(1: CheckUserExistsByIdRequst req)(api.post="/v1/user/check")
   VerifyCodeResponse VerifyCode(1: VerifyCodeRequest req)(api.post="/v1/user/verifycode")
   SendCodeResponse SendCode(1: SendCodeRequest req)(api.post="/v1/user/sendcode")
}