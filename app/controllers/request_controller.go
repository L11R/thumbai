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

package controllers

import (
	"net/http"
	"thumbai/app/access"
	"thumbai/app/gomod"
	"thumbai/app/models"
	"thumbai/app/proxy"
	"thumbai/app/settings"
	"thumbai/app/vanity"

	"aahframe.work"
	"aahframe.work/ahttp"
)

// RequestController handles the classic `go get` handling, gonna become legacy.
type RequestController struct {
	*aah.Context
}

// Handle method handles Go vanity package request. If not found then it passes
// control over to proxy pass.
func (c *RequestController) Handle() {
	var pkg *models.VanityPackage
	if c.Req.Method == ahttp.MethodGet {
		pkg = vanity.Lookup(c.Req.Host, c.Req.Path)
	}

	if pkg == nil {
		proxy.Do(c.Context)
		return
	}

	c.Reply().HTMLl("goget.html", aah.Data{
		"Vanity":    pkg,
		"GoDocHost": settings.GoDocHost,
	})
}

// Health method returns the health of the proxies and go mod respository.
func (c *RequestController) Health() {
	result := aah.Data{
		"proxies": proxy.Health(c.Context),
	}
	if !gomod.Settings.Enabled || access.GoModDisabled {
		result["gomod_repository"] = aah.Data{
			"status":      "service unavailable",
			"status_code": http.StatusServiceUnavailable,
		}
	} else {
		result["gomod_repository"] = aah.Data{
			"status":      "service available",
			"status_code": http.StatusOK,
		}
	}
	c.Reply().JSON(result)
}
