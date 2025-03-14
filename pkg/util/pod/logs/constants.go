// Copyright 2019 Yunion
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

package logs

const (
	// fixed width version of time.RFC3339Nano
	RFC3339NanoFixed = "2006-01-02T15:04:05.000000000Z07:00"
	// variable width RFC3339 time format for lenient parsing of strings into timestamps
	RFC3339NanoLenient = "2006-01-02T15:04:05.999999999Z07:00"
)
