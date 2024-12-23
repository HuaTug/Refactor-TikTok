// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode    = 10000
	ServiceErrCode = 10001
	ParamErrCode   = 10002

	UserAlreadyExistErrCode    = 10003
	AuthorizationFailedErrCode = 10004
	UserNotExistErrCode        = 10005

	MysqlErrCode         = 10006
	RedisErrCode         = 10007
	ElasticSearchErrCode = 10008
	OssErrCode           = 10009
	RabbitMQErrCode      = 10010

	RpcErrCode   = 10011
	RequestError = 10012

	TokenInvailedErrCode   = 10013
	TokenExpireTimeErrCode = 10014

	DataProcessFailed = 10015
	VerifyCodeErrCode = 10016
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(SuccessCode, "Success")
	ServiceErr             = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	UserNotExistErr        = NewErrNo(UserNotExistErrCode, "User not exisis")

	MysqlErr         = NewErrNo(MysqlErrCode, "Mysql connection failed")
	RedisErr         = NewErrNo(RedisErrCode, "Redis connection failed")
	ElasticSearchErr = NewErrNo(ElasticSearchErrCode, "ElasticSearch startup failed")
	OssErr           = NewErrNo(OssErrCode, "Oss startup failed")
	RabbitMQErr      = NewErrNo(RabbitMQErrCode, "RabbitMQ connnection failed")

	RpcErr     = NewErrNo(RpcErrCode, "Rpc startup failed")
	RequestErr = NewErrNo(RequestError, "Request failed")

	TokenInvailedErr   = NewErrNo(TokenInvailedErrCode, "Token is invailed")
	TokenExpireTimeErr = NewErrNo(TokenExpireTimeErrCode, "Token is expired")

	DataProcessErr = NewErrNo(DataProcessFailed, "DataProcess failed")
	VerifyCodeErr  = NewErrNo(VerifyCodeErrCode, "VerifyCode failed")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
