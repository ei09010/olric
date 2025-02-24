// Copyright 2018-2021 Burak Sezer
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

package storage

type Entry interface {
	SetKey(string)
	Key() string
	SetValue([]byte)
	Value() []byte
	SetTTL(int64)
	TTL() int64
	SetTimestamp(int642 int64)
	Timestamp() int64
	Encode() []byte
	Decode([]byte)
	SetMemFlags(int32)
	GetMemFlags() int32
	SetExpire(int64)
	GetExpire() int64
	SetCasUnique(int64)
	GetCasUnique() int64
}
