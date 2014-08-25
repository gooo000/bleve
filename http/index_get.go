//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.
package http

import (
	"fmt"
	"net/http"

	"github.com/couchbaselabs/bleve"

	"github.com/gorilla/mux"
)

type GetIndexHandler struct{}

func NewGetIndexHandler() *GetIndexHandler {
	return &GetIndexHandler{}
}

func (h *GetIndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// find the name of the index to create
	indexName := mux.Vars(req)["indexName"]
	if indexName == "" {
		showError(w, req, fmt.Sprintf("index name is required", indexName), 400)
		return
	}

	index := IndexByName(indexName)
	if index == nil {
		showError(w, req, fmt.Sprintf("no such index '%s'", indexName), 404)
		return
	}

	rv := struct {
		Status  string              `json:"status"`
		Name    string              `json:"name"`
		Mapping *bleve.IndexMapping `json:"mapping"`
	}{
		Status:  "ok",
		Name:    indexName,
		Mapping: index.Mapping(),
	}
	mustEncode(w, rv)
}