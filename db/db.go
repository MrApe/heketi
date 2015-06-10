//
// Copyright (c) 2014 The heketi Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

type HeketiDB interface {
	Close()
	Put(key, val []byte, index uint64) error
	Get(key, val []byte, index uint64) error
	Delete(key []byte, index uint64) error
	String() string
}

func NewHeketiDB(config interface{}, name string) HeketiDB {
	switch name {
	case "boltdb":
		return NewBoltDB((config.(string)))
	default:
		return nil
	}
}
