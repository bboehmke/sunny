// Copyright 2021 Benjamin Böhmke <benjamin@boehmke.net>.
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

package sunny

// Log is used to log some internal trace messages
var Log Logger = new(NopeLogger)

// Logger interface for trace logger
type Logger interface {
	Printf(format string, v ...interface{})
}

// NopeLogger implements Logger without any action
type NopeLogger struct{}

func (n NopeLogger) Printf(format string, v ...interface{}) {}