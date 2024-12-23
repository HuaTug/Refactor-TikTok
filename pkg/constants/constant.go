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

package constants

const (
	NoteTableName           = "note"
	UserTableName           = "user"
	SecretKey               = "secret key"
	IdentityKey             = "id"
	Total                   = "total"
	Notes                   = "notes"
	NoteID                  = "note_id"
	ApiServiceName          = "demoapi"
	NoteServiceName         = "demonote"
	UserServiceName         = "demouser"
	ESNoKeywordsFlag       = ``
	ESNoTimeFilterFlag     = -1
	ESNoUsernameFilterFlag = ``
	ESNoPageParamFlag      = -1
	CPURateLimit    float64 = 80.0

)

var (
	MySQLDefaultDSN = "root:root@tcp(localhost:3306)/note?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress     = "localhost:2379"
)

const DataFormate = "2006-01-02"
const DefaultPageSize = 10
const DefaultLimit = 10
