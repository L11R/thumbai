// Copyright Jeevanandam M. (https://github.com/jeevatkm, jeeva@myjeeva.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package settings

import (
	"strings"

	"aahframe.work"
)

// THUMBAI settings
var (
	ServerHeader string
	GoDocHost    string
)

// Load method loads required thumbai config values on app startup.
func Load(_ *aah.Event) {
	cfg := aah.App().Config()
	ServerHeader = cfg.StringDefault("thumbai.server.header", "")
	GoDocHost = strings.TrimSuffix(cfg.StringDefault("thumbai.admin.godoc_host", "https://godoc.org"), "/")
}
